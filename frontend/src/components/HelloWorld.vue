<template>
  <div class="hello-world" v-text="msg"></div>
</template>

<script>
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
  height: 30px;
  line-height: 30px;
  margin: 16px auto;
  text-align: center;
  font-size: 26px;
}
</style>
