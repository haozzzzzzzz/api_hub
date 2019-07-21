<template>
    <div :id="swagger_data.id" class="swagger-vue">

    </div>
</template>

<script>
    import SwaggerUI from 'swagger-ui'
    import 'swagger-ui/dist/swagger-ui.css'

    export default {
        name: "swagger_vue",
        props: ["swagger_data"],
        mounted() {
            SwaggerUI({
                dom_id: "#"+this.swagger_data.id,
                url: this.swagger_data.url || "https://petstore.swagger.io/v2/swagger.json",
            });
        }
    }
</script>

<style scoped>

</style>