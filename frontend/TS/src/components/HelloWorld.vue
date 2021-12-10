<template>
  <div class="hello-world" v-text="msg"></div>
</template>

<script lang="ts">
export default {
  name: "HelloWorld",
  props: {
    msg: String,
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.hello-world {
  height: 76px;
  line-height: 38px;
  margin: 16px 150px;
  text-align: center;
  font-size: 26px;
}
</style>
