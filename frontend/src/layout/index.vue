<template>
    <el-container>
        <el-aside>
            <sidebar></sidebar>
        </el-aside>
        <el-main>
            <app_main></app_main>
        </el-main>
    </el-container>
</template>

<script>
    import sidebar from './components/sidebar'
    import app_main from "./components/app_main";

    export default {
        name: "layout",
        components: {
            app_main,
            sidebar,
        }
    }
</script>

<style lang="scss">
    .el-container {
        height: 100%;
    }
    .el-aside {
        width: 210px !important;
    }
    .el-main {
        padding-top: 10px !important;
        padding-bottom: 10px !important;
    }
</style>