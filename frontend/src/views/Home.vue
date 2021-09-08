<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png" />
    <HelloWorld :msg="$t('welcome')" />
    <button @click="changeLanguage" v-text="'I18N'"></button>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from "@/components/HelloWorld.vue";
import i18n from "@/i18n";

export default {
  name: "Home",
  components: {
    HelloWorld,
  },
  setup() {
    const changeLanguage = function () {
      if (i18n.global.locale === "en") {
        i18n.global.locale = "zh-Hans";
      } else {
        i18n.global.locale = "en";
      }
    };
    return { changeLanguage };
  },
};
</script>
