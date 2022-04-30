<template>
  <!-- Header -->
  <!-- 头部 -->
  <div class="header" data-wails-drag>
    <!-- navigation -->
    <!-- 导航 -->
    <div class="nav" data-wails-no-drag>
      <router-link to="/">{{ t("nav.home") }}</router-link>
      <router-link to="/about">{{ t("nav.about") }}</router-link>
    </div>
    <!-- Menu -->
    <!-- 菜单 -->
    <div class="menu" data-wails-no-drag>
      <div class="language">
        <div v-for="item in languages" :key="item" :class="{ active: item === locale }"
          @click="onclickLanguageHandle(item)" class="lang-item">
          {{ t("languages." + item) }}
        </div>
      </div>
      <div class="bar">
        <div class="bar-btn" @click="onclickMinimise">
          {{ t("topbar.minimise") }}
        </div>
        <div class="bar-btn" @click="onclickQuit">{{ t("topbar.quit") }}</div>
      </div>
    </div>
  </div>
  <!-- Page -->
  <!-- 页面 -->
  <div class="view">
    <router-view />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useI18n } from "vue-i18n";
import { WindowMinimise, Quit } from "../../../wailsjs/runtime";

export default defineComponent({
  setup() {
    const { t, availableLocales, locale } = useI18n({ useScope: "global" });
    // List of supported languages
    // 支持的语言列表
    const languages = availableLocales;

    // Click to switch language
    // 点击切换语言
    const onclickLanguageHandle = (item: string) => {
      item !== locale.value ? (locale.value = item) : false;
    };

    const onclickMinimise = () => {
      WindowMinimise();
    };
    const onclickQuit = () => {
      Quit();
    };

    return {
      t,
      languages,
      locale,
      onclickLanguageHandle,
      onclickMinimise,
      onclickQuit,
    };
  },
});
</script>

<style lang="scss">
@import url("./assets/css/reset.css");
@import url("./assets/css/font.css");

html {
  width: 100%;
  height: 100%;
}

body {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
  font-family: "JetBrainsMono";
  background-color: transparent;
}

#app {
  position: relative;
  // width: 900px;
  // height: 520px;
  height: 100%;
  background-color: rgba(219, 188, 239, 0.9);
  overflow: hidden;
}

.header {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  align-items: center;
  justify-content: space-between;
  height: 50px;
  padding: 0 10px;
  background-color: rgba(171, 126, 220, 0.9);

  .nav {
    a {
      display: inline-block;
      min-width: 50px;
      height: 30px;
      line-height: 30px;
      padding: 0 5px;
      margin-right: 8px;
      background-color: #ab7edc;
      border-radius: 2px;
      text-align: center;
      text-decoration: none;
      color: #000000;
      font-size: 14px;
      white-space: nowrap;

      &:hover,
      &.router-link-exact-active {
        background-color: #d7a8d8;
        color: #ffffff;
      }
    }
  }

  .menu {
    display: flex;
    flex-direction: row;
    flex-wrap: nowrap;
    align-items: center;
    justify-content: space-between;

    .language {
      margin-right: 20px;
      border-radius: 2px;
      background-color: #c3c3c3;
      overflow: hidden;

      .lang-item {
        display: inline-block;
        min-width: 50px;
        height: 30px;
        line-height: 30px;
        padding: 0 5px;
        background-color: transparent;
        text-align: center;
        text-decoration: none;
        color: #000000;
        font-size: 14px;

        &:hover {
          background-color: #ff050542;
          cursor: pointer;
        }

        &.active {
          background-color: #ff050542;
          color: #ffffff;
          cursor: not-allowed;
        }
      }
    }

    .bar {
      display: flex;
      flex-direction: row;
      flex-wrap: nowrap;
      align-items: center;
      justify-content: flex-end;
      min-width: 150px;

      .bar-btn {
        display: inline-block;
        min-width: 80px;
        height: 30px;
        line-height: 30px;
        padding: 0 5px;
        margin-left: 8px;
        background-color: #ab7edc;
        border-radius: 2px;
        text-align: center;
        text-decoration: none;
        color: #000000;
        font-size: 14px;

        &:hover {
          background-color: #d7a8d8;
          color: #ffffff;
          cursor: pointer;
        }
      }
    }
  }
}

.view {
  position: absolute;
  top: 50px;
  left: 0;
  right: 0;
  bottom: 0;
  overflow: hidden;
}
</style>
