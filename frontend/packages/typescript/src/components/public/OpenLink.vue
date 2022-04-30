<template>
  <span class="openlink" @click="onClickhandle">
    <slot></slot>
  </span>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { BrowserOpenURL } from "../../../../../wailsjs/runtime";
export default defineComponent({
  name: "OpenLink",
  props: {
    href: String,
  },
  setup(props) {
    const onClickhandle = () => {
      // You cannot use the a tag directly, you need to call the Go method here to open the link using the default browser.
      // 不能直接使用a标签，需要在这里调用Go方法以使用默认浏览器打开链接。

      BrowserOpenURL(props.href);
    };

    return {
      onClickhandle,
    };
  },
});
</script>

<style lang="scss">
.openlink {
  cursor: pointer;
  text-decoration: underline;
}
</style>
