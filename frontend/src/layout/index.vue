<template>
    <el-container id="app-layout">
        <el-aside>
            <sidebar></sidebar>
        </el-aside>
        <el-main>main</el-main>
    </el-container>
</template>

<script>
    import sidebar from './components/sidebar'
    export default {
        name: "layout",
        el: "#layout",
        components: {
            sidebar,
        },
    }
</script>

<style lang="scss" scoped>
    #app-layout {
        height: 100%;
    }
</style>