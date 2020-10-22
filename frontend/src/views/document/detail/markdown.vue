<template>
  <div class="document-detail">
    <Markdown_vue :markdown_data="markdownData"></Markdown_vue>
  </div>
</template>

<script>
import apis from "@/api/apis"
import Markdown_vue from "@/components/markdown_vue/index";
export default {
  name: "markdown",
  components: {Markdown_vue},
  props: ['tabData'],
  computed: {
    markdownData() {
      let url = apis.docDetailSpecUrl(this.tabData.doc_id);
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