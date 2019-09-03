<template>
    <el-tabs v-model="editableTabsValue" type="card" editable @edit="handleTabsEdit" >
        <el-tab-pane
                v-for="(item, index) in editableTabs"
                :key="item.name"
                :label="item.title"
                :name="item.name"
                :disabled="item.disable"
                :closable="item.closable"
        >
            <transition name="fade-transform" mode="out-in">
                <keep-alive>
                    <component v-bind:is="currentTabComponent" @custom-tab-open="customTabOpen" :data="item.tabData"></component>
                </keep-alive>
            </transition>
        </el-tab-pane>
    </el-tabs>
</template>

<script>
    import document_detail from './detail'
    import document_list from './list'

    export default {
        name: "document",
        data(){
            return {
                editableTabsValue: '0',
                editableTabs: [
                    {
                        title: "document list",
                        name: "0",
                        route: "/document/list",
                        tabType: 'list',
                        tabData: {}
                    }
                ],
                tabCounter: 0,
            }
        },
        computed: {
            currentTabComponent: function () {
                if ( this.editableTabsValue !== '0' ) {
                    return document_detail
                }
                return document_list
            },
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
                            tabData: tabData
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