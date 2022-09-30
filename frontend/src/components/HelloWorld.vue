<script setup lang="ts">
defineProps<{
  msg: string;
}>();
</script>

<template>
  <div class="hello-world" v-text="msg"></div>
</template>

<style scoped>
.hello-world {
  height: 76px;
  line-height: 38px;
  margin: 16px 150px;
  text-align: center;
  font-size: 26px;
}
</style>
