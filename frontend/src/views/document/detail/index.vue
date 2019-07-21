<template>
    <div class="document-detail">
        <Swagger_vue :swagger_data="swaggerData"></Swagger_vue>
    </div>
</template>

<script>
    import Swagger_vue from "@/components/swagger_vue/index";
    export default {
        name: "document_detail",
        components: {Swagger_vue},
        props: ['data'],
        computed: {
            swaggerData() {
                return {
                    id: 'swagger-' + this.data.doc_id,
                    url: this.data.spec_url
                }
            }
        }
    }
</script>

<style scoped>
    .document-detail {
        height: 100%;
    }
</style>