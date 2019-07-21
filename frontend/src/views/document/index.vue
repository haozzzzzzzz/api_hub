<template>
    <el-tabs v-model="editableTabsValue" type="card" editable @edit="handleTabsEdit">
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
                    <component v-bind:is="currentTabComponent"></component>
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
                        route: "/document/list"
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
                    let newTabName = ++ this.tabCounter + '';
                    this.editableTabs.push({
                        title: 'New Tab ' + newTabName,
                        name: newTabName,
                        route: "/document/detail"
                    });

                    this.editableTabsValue = newTabName

                } else if (action === 'remove') {
                    let tabs = this.editableTabs;
                    let activeName = this.editableTabsValue;
                    if (activeName === targetName) {
                        tabs.forEach((tab, index)=>{
                            if ( tab.Name === targetName ) {
                                let nextTab = tabs[index+1] || tabs[index-1];
                                if (nextTab) {
                                    activeName = nextTab.name
                                }
                            }
                        })
                    }

                    this.editableTabsValue = activeName;
                    this.editableTabs = tabs.filter(tab=>tab.name !== targetName)
                }

            }
        }
    }
</script>

<style scoped>

</style>