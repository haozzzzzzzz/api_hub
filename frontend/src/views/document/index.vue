<template>
    <el-tabs v-model="tabActiveName" type="card" closable @edit="handleTabsEdit" @tab-click="tabClick">
        <el-tab-pane
                v-for="item in tabs"
                :key="item.name"
                :label="item.title"
                :name="item.name"
                :disabled="item.disable"
                :closable="item.closable"
        >
            <transition name="fade-transform" mode="out-in">
                <keep-alive>
                    <component v-bind:is="item.tabView" @custom-tab-open="customTabOpen" @search-change="searchChange" :tab-data="item.tabData"></component>
                </keep-alive>
            </transition>
        </el-tab-pane>
    </el-tabs>
</template>

<script>
    import swagger_detail from './detail/swagger'
    import markdown_detail from './detail/markdown'
    import document_list from './list'

    let documentTabId = '0'; // document tab id
    let defaultTabId = documentTabId; // default tab id

    // default tab
    let defaultDocumentTab = {
        title: "document list",
        name: defaultTabId,
        tabType: 'list',
        tabData: {
            search: "",
        },
        tabView: document_list,
    };

    function tabOpen(data) {
        let tabType = data.tabType;
        let tabData = data.tabData;

        let tabActive = null;
        switch (tabType) {
            case 'detail':
            {
                let existTabName = "";
                for( let tab of this.tabs) {
                    if (tab.tabType === tabType && tab.tabData.doc_id == tabData.doc_id) { // 存在
                        existTabName = tab.name;
                        tabActive = tab;
                        break;
                    }
                }

                if (existTabName !== '') {
                    this.tabActiveName = existTabName;
                    return tabActive
                }

                let newTabName = ++ this.tabCounter + '';
                let newTab = {
                    title: (tabData.title ? "(" + tabData.doc_id + ") " + tabData.title : false) || 'New Tab',
                    name: newTabName,
                    tabType: 'detail',
                    tabData: tabData,
                    // tabView: swagger_detail,
                };

                const DocTypeSwagger = "swagger";
                const DocTypeMarkdown = "markdown";

                switch (tabData.doc_type) {
                    case DocTypeSwagger:
                      newTab.tabView = swagger_detail;
                      break;
                    case DocTypeMarkdown:
                      newTab.tabView = markdown_detail;
                      break;
                    default:
                      newTab.tabView = undefined;
                }

                tabActive = newTab;
                this.tabs.push(newTab);
                this.tabActiveName = newTabName;

            }
        }

        this.tabMap = {};
        this.tabs.forEach((tab, index)=>{
            this.tabMap[tab.name] = tab;
        });

        return tabActive;
    }

    export default {
        name: "document",
        data(){
            let retData = {
                tabActiveName: defaultTabId,
                tabs: [
                    defaultDocumentTab
                ],
                tabMap: {
                    defaultTabId: defaultDocumentTab
                },
                tabCounter: 0,
            };

            let query = this.$route.query;

            switch (query.tab_type) {
                case 'detail':
                {
                    let doc_id = parseInt(query.doc_id, 10);
                    let title = query.title;
                    let doc_type = query.doc_type || "swagger";

                    tabOpen.call(retData, {
                        tabType: "detail",
                        tabData: {
                            doc_id: doc_id,
                            title: title,
                            doc_type: doc_type,
                        }
                    });

                }
                    break;

                case 'list':
                {
                    let search = query.search;
                    if (search != undefined && search != "") {
                        defaultDocumentTab.tabData.search = search
                    }
                }
                    break;
            }

            return retData;
        },

        watch:{
            tabs(newTabs, oldTabs) {
                this.tabMap = {};
                newTabs.forEach((tab, index)=>{
                    this.tabMap[tab.name] = tab;
                });
            }
        },

        methods: {
            handleTabsEdit(targetName, action) {
                switch (action) {
                    case "add":
                    {
                        this.tabOpen({
                            tabType: 'detail',
                            tabData: {
                                doc_id: 'new_tab' // required for identifying
                            }
                        });
                        break;
                    }
                    case "remove":
                    {
                        let tabs = this.tabs;
                        let activeName = this.tabActiveName;
                        if ( targetName === defaultTabId ) {
                            return
                        }

                        tabs.forEach((tab, index)=>{
                            if ( tab.name === targetName ) {
                                let nextTab = tabs[index+1] || tabs[index-1];
                                if (nextTab) {
                                    activeName = nextTab.name
                                }
                            }
                        });

                        this.tabActiveName = activeName;
                        this.tabs = tabs.filter(tab=>tab.name !== targetName);
                        break;
                    }
                    default:
                        break;
                }

            },
            customTabOpen(eventData) {
                this.tabOpen(eventData)
            },

            searchChange(eventData){
                defaultDocumentTab.tabData.search = eventData.tabData.search;
                this.changeUrl(defaultDocumentTab);
            },

            tabOpen(data){
                let activeTab = tabOpen.call(this, data);
                this.changeUrl(activeTab);
            },

            tabClick(tab){
                let tabName = tab.name;
                let editableTab = this.tabMap[tabName];
                this.changeUrl(editableTab);
            },

            changeUrl(tab){
                if (tab == undefined) {
                    console.error("change url(undefined)");
                    return
                }

                let query = {};
                switch (tab.tabType) {
                    case "detail":
                        query = {
                            tab_type: tab.tabType,
                            doc_id: tab.tabData.doc_id,
                            title: tab.tabData.title,
                            doc_type: tab.tabData.doc_type
                        };
                        break;

                    case "list":
                        query = {
                            tab_type: tab.tabType,
                            search: tab.tabData.search,
                        };
                        break;
                }

                this.$router.push({
                    query: query
                })
            }
        }
    }
</script>

<style scoped>
</style>