import { createI18n } from "vue-i18n";

import en from "./locales/en.json";
import zhHans from "./locales/zh-Hans.json";

const i18n = createI18n({
  locale: "ja",
  fallbackLocale: "en",
  legacy: false,
  messages: {
    en,
    "zh-Hans": zhHans,
  },
});

export default i18n;
