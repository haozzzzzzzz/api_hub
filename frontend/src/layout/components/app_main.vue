<template>
    <transition name="fade-transform" mode="out-in">
        <keep-alive>
            <router-view :key="key"></router-view>
        </keep-alive>
    </transition>
</template>

<script>
    export default {
        name: "app_main",
        computed: {
            key(){
                return this.$route.path
            }
        }
    }
</script>

<style scoped>

</style>