<template>
    <el-tabs v-model="editableTablsValue" type="card" editable @edit="handleTabsEdit">
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
                    <router-view :key="item.route"></router-view>
                </keep-alive>
            </transition>
        </el-tab-pane>
    </el-tabs>
</template>

<script>
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
                tabIndex: 0
            }
        },
        methods: {
            handleTabsEdit(targetName, action) {
                if (action === 'add') {
                    let newTabName = ++ this.tabIndex + '';
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