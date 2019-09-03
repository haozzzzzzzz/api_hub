<template>
    <div id="sidebar">
        <div class="sidebar-logo-container">
            <transition name="sidebarLogoFade">
                <router-link class="sidebar-logo-link" to="/">
                    <img v-if="logo" :src="logo" class="sidebar-logo">
                    <h1 class="sidebar-title">{{ title }} </h1>
                </router-link>
            </transition>
        </div>
        <el-scrollbar>
            <el-menu
                    mode="vertical"
                    :collapse="false"
                    :collapse-transition="false"
                    :unique-opened="true"

                    :background-color="variables.menuBg"
                    :text-color="variables.menuText"
                    :active-text-color="variables.menuActiveText"
                    router
            >
                <el-menu-item index="/document">
                    <i class="el-icon-document"></i>
                    <span slot="title">Document</span>
                </el-menu-item>
<!--                <el-menu-item index="/mock">-->
<!--                    <i class="el-icon-connection"></i>-->
<!--                    <span slot="title">Mock</span>-->
<!--                </el-menu-item>-->
<!--                <el-menu-item>-->
<!--                    <a class="menu-link" href="https://github.com/haozzzzzzzz/api_hub" target="_blank" ref="noopener">-->
<!--                        <i class="el-icon-service"></i>-->
<!--                        <span slot="title">Github</span>-->
<!--                    </a>-->
<!--                </el-menu-item>-->
            </el-menu>
        </el-scrollbar>
    </div>
</template>

<script>
    import variables from '@/styles/variables.scss'
    import logo from '@/assets/logo.png'

    export default {
        name: "sidebar",
        data: function() {
            return {
                title:"Api Hub",
                logo: logo
            }
        },
        computed: {
            variables() {
                return variables
            }
        }
    }
</script>

<style lang="scss" scoped>
    .sidebarLogoFade-enter-active {
        transition: opacity 1.5s;
    }

    .sidebarLogoFade-enter,
    .sidebarLogoFade-leave-to {
        opacity: 0;
    }

    .sidebar-logo-container {
        position: relative;
        width: 100%;
        height: 50px;
        line-height: 50px;
        background: #2b2f3a;
        text-align: center;
        overflow: hidden;

        & .sidebar-logo-link {
            height: 100%;
            width: 100%;

            & .sidebar-logo {
                width: 32px;
                height: 32px;
                vertical-align: middle;
                margin-right: 12px;
            }

            & .sidebar-title {
                display: inline-block;
                margin: 0;
                color: #fff;
                font-weight: 600;
                line-height: 50px;
                font-size: 14px;
                font-family: Avenir, Helvetica Neue, Arial, Helvetica, sans-serif;
                vertical-align: middle;
            }
        }

    }
    .menu-link {
        text-decoration: none;
        display: inline-block;
        width: 100%;
        overflow: hidden;
        color: inherit;
    }
</style>
<style lang="scss">
    #sidebar {
        height: 100%;

        .logo {
            height: 50px;
        }

        .el-menu {
            width: 100%;
            height: 100%;
        }

        .el-scrollbar {
            overflow: hidden;
            height: calc(100% - 50px);

        }
        .el-scrollbar__wrap{
            overflow: auto;
            overflow-y: hidden!important;
        }

        .el-scrollbar__view {
            height: 100%;
        }

        .el-menu-item {
            height: 56px;
        }
    }


</style>