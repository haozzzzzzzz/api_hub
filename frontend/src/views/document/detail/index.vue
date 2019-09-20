<template>
    <div class="document-detail">
        <Swagger_vue :swagger_data="swaggerData"></Swagger_vue>
    </div>
</template>

<script>
    import apis from '@/api/apis'
    import Swagger_vue from "@/components/swagger_vue/index";
    export default {
        name: "document_detail",
        components: {
            Swagger_vue,
        },
        props: ['data'],
        computed: {
            swaggerData() {
                let url = apis.docDetailSpecUrl(this.data.doc_id);
                return {
                    url: url
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