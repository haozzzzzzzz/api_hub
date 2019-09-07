<template>
    <el-tabs v-model="editableTabsValue" type="card" editable @edit="handleTabsEdit" >
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

    let documentTabId = "0";
    let defaultTabId = documentTabId;

    export default {
        name: "document",
        data(){
            return {
                editableTabsValue: defaultTabId,
                editableTabs: [
                    {
                        title: "document list",
                        name: defaultTabId,
                        route: "/document/list",
                        tabType: 'list',
                        tabData: {},
                        tabView: document_list,
                    }
                ],
                tabCounter: 0,
            }
        },
        methods: {
            handleTabsEdit(targetName, action) {
                if (action === 'add') {
                    this.tabOpen({
                        tabType: 'detail',
                        tabData: {
                            doc_id: 'new_tab' // required for identifying
                        }
                    })

                } else if (action === 'remove') {
                    let tabs = this.editableTabs;
                    let activeName = this.editableTabsValue;

                    if ( activeName === defaultTabId ) {
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
                    this.editableTabs = tabs.filter(tab=>tab.name !== targetName)
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
                        this.editableTabs.push({
                            title: (tabData.title ? "(" + tabData.doc_id + ") " + tabData.title : false) || 'New Tab',
                            name: newTabName,
                            route: "/document/detail",
                            tabType: 'detail',
                            tabData: tabData,
                            tabView: document_detail,
                        });

                        this.editableTabsValue = newTabName;
                    }
                }
            }
        }
    }
</script>

<style scoped>
</style>