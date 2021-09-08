import { createI18n } from 'vue-i18n'

const i18n = createI18n({
  locale: "en",
  fallbackLocale: 'en',
  messages: {
    "zh-Hans": require('./messages/zh-Hans.json'),
    "en": require('./messages/en.json')
  }
})

export default i18n