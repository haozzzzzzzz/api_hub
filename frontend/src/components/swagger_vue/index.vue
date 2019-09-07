<template>
    <div class="swagger-vue"></div>
</template>

<script>
    import SwaggerUI from 'swagger-ui'
    import 'swagger-ui/dist/swagger-ui.css'
    import config from '@/config/config'
    import RequestPlugins from "../../plugins/request";

    // https://github.com/swagger-api/swagger-ui/blob/HEAD/docs/usage/configuration.md
    export default {
        name: "swagger_vue",
        props: ["swagger_data"],
        mounted() {
            SwaggerUI({
                domNode: this.$el,
                url: this.swagger_data.url || config.AppConfig.default_swagger_json_url,
                defaultModelExpandDepth: 4,
                filter: true,
                displayRequestDuration: true,
                requestInterceptor: function (req) {
                    req.headers['Product-Id'] = '39';
                    return RequestPlugins.SwaggerRequestPlugins.requestInterceptor(req);
                },
                responseInterceptor: function (resp) {
                    return RequestPlugins.SwaggerRequestPlugins.responseInterceptor(resp);
                },
            });
        }
    }
</script>

<style scoped>

</style>