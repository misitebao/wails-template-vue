import OpenLink from "@/components/public/OpenLink.vue";

// Encapsulate global components as plug-ins
// 将全局组件封装为插件

export default {
  install(app) {
    app.component(OpenLink.name, OpenLink);
  },
};
