<template>
    <el-tabs v-model="editableTabsValue" type="card" editable @edit="handleTabsEdit" @tab-click="tabClick">
        <el-tab-pane
                v-for="item in editableTabs"
                :key="item.name"
                :label="item.title"
                :name="item.name"
                :disabled="item.disable"
                :closable="item.closable"
        >
            <transition name="fade-transform" mode="out-in">
                <keep-alive>
                    <component v-bind:is="item.tabView" @custom-tab-open="customTabOpen" :data="item.tabData"></component>
                </keep-alive>
            </transition>
        </el-tab-pane>
    </el-tabs>
</template>

<script>
    import document_detail from './detail'
    import document_list from './list'

    let documentTabId = '0';
    let defaultTabId = documentTabId;

    let defaultDocumentTab = {
        title: "document list",
        name: defaultTabId,
        tabType: 'list',
        tabData: {},
        tabView: document_list,
    };

    export default {
        name: "document",
        data(){
            return {
                editableTabsValue: defaultTabId,
                editableTabs: [
                    defaultDocumentTab
                ],
                editableTabMap: {
                    defaultTabId: defaultDocumentTab
                },
                tabCounter: 0,
            }
        },
        watch:{
            editableTabs(newTabs, oldTabs) {
                this.editableTabMap = {};
                this.editableTabs.forEach((tab, index)=>{
                    this.editableTabMap[tab.name] = tab;
                });
            }
        },
        mounted(){
            let query = this.$route.query;
            switch (query.tab_type) {
                case 'detail':
                {
                    let doc_id = query.doc_id;
                    let title = query.title;
                    this.tabOpen({
                        tabType: "detail",
                        tabData: {
                            doc_id: doc_id,
                            title: title,
                        }
                    })
                }
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
                        let tabs = this.editableTabs;
                        let activeName = this.editableTabsValue;
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

                        this.editableTabsValue = activeName;
                        this.editableTabs = tabs.filter(tab=>tab.name !== targetName);
                        break;
                    }
                    default:
                        break;
                }

            },
            customTabOpen(eventData) {
                this.tabOpen(eventData)
            },
            tabOpen(data) {
                let tabType = data.tabType;
                let tabData = data.tabData;

                switch (tabType) {
                    case 'detail':
                    {
                        let existTabName = "";
                        for( let tab of this.editableTabs) {
                            if (tab.tabType === tabType && tab.tabData.doc_id === tabData.doc_id) { // 存在
                                existTabName = tab.name;
                                break;
                            }
                        }

                        if (existTabName !== '') {
                            this.editableTabsValue = existTabName;
                            return
                        }

                        let newTabName = ++ this.tabCounter + '';
                        let editableTable = {
                            title: (tabData.title ? "(" + tabData.doc_id + ") " + tabData.title : false) || 'New Tab',
                            name: newTabName,
                            tabType: 'detail',
                            tabData: tabData,
                            tabView: document_detail,
                        };

                        this.editableTabs.push(editableTable);
                        this.editableTabsValue = newTabName;
                        this.changeUrl(editableTable);
                    }
                }
            },

            tabClick(tab){
                let tabName = tab.name;
                let editableTab = this.editableTabMap[tabName];
                if (editableTab===undefined){
                    return
                }

                this.changeUrl(editableTab);
            },

            changeUrl(editableTab){
                this.$router.push({
                    query: {
                        tab_type: editableTab.tabType,
                        doc_id: editableTab.tabData.doc_id,
                        title: editableTab.tabData.title
                    }
                })
            }
        }
    }
</script>

<style scoped>
</style>