<template>
  <div class="markdown-vue markdown-body">
    <vue-markdown :source="content" toc toc-id="markdown-toc"></vue-markdown>
  </div>
</template>

<script>
  import apis from "@/api/apis"

  // https://github.com/miaolz123/vue-markdown
  import VueMarkdown from "vue-markdown"
  import "github-markdown-css"

  export default {
    name: "markdown_vue",
    data() {
        return {
          content: ""
        }
    },

    props: ["markdown_data"],
    components: {VueMarkdown},
    mounted() {
      let url = this.markdown_data.url;
      if (!url) {
        return
      }

      apis.docDetailText(url, (text) => {
          this.content = text;
      });

    }
  }
</script>

<style>
.markdown-body {
  box-sizing: border-box;
  min-width: 200px;
  /*max-width: 980px;*/
  margin: 0 auto;
  padding: 20px;
}
</style>