<template>
  <!-- Header -->
  <!-- 头部 -->
  <div class="header">
    <!-- navigation -->
    <!-- 导航 -->
    <div class="nav">
      <router-link to="/">{{ $t("nav.home") }}</router-link>
      <router-link to="/about">{{ $t("nav.about") }}</router-link>
    </div>
    <!-- Menu -->
    <!-- 菜单 -->
    <div class="menu">
      <div class="language">
        <div
          v-for="(item) in languages"
          :key="item"
          :class="{ active: item === locale }"
          @click="onclickLanguageHandle(item)"
          class="lang-item"
        >{{ $t("languages." + item) }}</div>
      </div>
      <div class="bar">
        <div class="bar-btn" @click="onclickMinimise">{{ $t("topbar.minimise") }}</div>
        <div class="bar-btn" @click="onclickExit">{{ $t("topbar.exit") }}</div>
      </div>
    </div>
  </div>
  <!-- Page -->
  <!-- 页面 -->
  <div class="view">
    <router-view />
  </div>
</template>

<script>
import { ref, watch } from 'vue'
import i18n from "@/i18n";

export default {
  setup() {
    // List of supported languages
    // 支持的语言列表
    const languages = i18n.global.availableLocales
    // Current language
    // 当前语言
    const locale = ref("zh-Hans")
    locale.value = i18n.global.locale

    // Click to switch language
    // 点击切换语言
    const onclickLanguageHandle = (item) => {
      (item !== locale.value) ? (locale.value = item) : false

    }
    // Monitor current language changes
    // 监听当前语言变动
    watch(locale, (newValue, oldValue) => {
      console.log('The new language is: ' + locale.value)
      i18n.global.locale = newValue
    })


    // Since the current js runtime has not been developed yet, so first call Go to complete. Later, 
    // it will be updated to be called directly when js is running.
    // 由于当前js运行时还没有开发完成，所以先调用Go完成。后续会更新为js运行时直接调用。
    const onclickMinimise = () => {
      alert(i18n.global.t("global.not-supported"))

    }
    const onclickExit = () => {
      alert(i18n.global.t("global.not-supported"))

    }

    return {
      languages,
      locale,
      onclickLanguageHandle,
      onclickMinimise,
      onclickExit
    }
  },
}
</script>

<style lang="stylus">
@import url("./assets/css/reset.css");
html,
body {
  width: 100%;
  height: 100%;
  margin 0
  padding 0
}
#app {
  position: relative;
  // width: 900px;
  // height: 520px;
  width 100%
  height 100%
  background-color: #dbbcef;
}
.header {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  align-items: center;
  justify-content: space-between;
  height: 50px;
  padding: 0 10px;
  background-color: #ab7edc;
  .nav {
    a {
      display: inline-block;
      width: 50px;
      height: 30px;
      line-height: 30px;
      padding 0 5px
      margin-right: 8px;
      background-color: #ab7edc;
      border-radius: 2px;
      text-align: center;
      text-decoration: none;
      color: #000000;
      font-size: 14px;
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
        padding 0 5px
        background-color: transparent;
        text-align: center;
        text-decoration: none;
        color: #000000;
        font-size: 14px;
        &:hover{
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
      justify-content flex-end
      min-width 150px
      .bar-btn {
        display: inline-block;
        min-width: 50px;
        height: 30px;
        line-height: 30px;
        padding 0 5px
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
  overflow hidden
}
</style>
