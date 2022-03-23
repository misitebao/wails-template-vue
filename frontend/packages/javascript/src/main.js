import { createApp } from "vue";
import App from "./App.vue";
import router from "@/router";
import i18n from "@/i18n";

// Register global common components
// 注册全局通用组件
import publicComponents from "@/components/public";

const app = createApp(App);

app.use(publicComponents);

app.use(router).use(i18n).mount("#app");
