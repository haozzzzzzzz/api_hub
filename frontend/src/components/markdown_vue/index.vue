<template>
  <div class="markdown-vue">
    <vue-markdown :source="content"></vue-markdown>
  </div>
</template>

<script>
  import apis from "@/api/apis"

  // https://github.com/miaolz123/vue-markdown
  import VueMarkdown from "vue-markdown"

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

<style scoped>

</style>