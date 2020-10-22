<template>
  <div class="markdown-vue markdown-body">
    <vue-markdown :source="content" toc toc-id="markdown-toc" @toc-rendered="tocRendered"></vue-markdown>
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
    },
    methods: {
      tocRendered: function (tocHtml) {
        // 如果存在哈希，则尝试跳转至哈希(不一定能100%成功)
        if(this.$route.hash) {
          setTimeout(()=>{
            location.href = this.$route.hash;
          }, 1)
        }
      }
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