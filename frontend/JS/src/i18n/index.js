import { createI18n } from "vue-i18n";

import zhHans from "./messages/zh-Hans.json";
import en from "./messages/en.json";
import fr from "./messages/fr.json";

const i18n = createI18n({
  locale: "en",
  fallbackLocale: "en",
  messages: {
    "zh-Hans": zhHans,
    en: en,
    fr: fr,
  },
});

export default i18n;
