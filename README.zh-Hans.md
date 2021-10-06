<p align="center">
  <img src="https://cdn.jsdelivr.net/gh/misitebao/CDN@main/logo/wails-template-vue.png" height="280" />
</p>
<p align="center">
  基于Vue和Vue-Router的Wails模板
</p>
<p align="center">
  <a href="https://github.com/misitebao/standard-repository/blob/main/LICENSE"><img alt="GitHub" src="https://img.shields.io/github/license/misitebao/standard-repository?style=flat-square"/></a>
  <a href="https://github.com/misitebao/standard-repository"><img alt="GitHub" src="https://img.shields.io/badge/Readme--Style-standard--repository-brightgreen?style=flat-square&color=f83500"/></a>
  <a href="https://github.com/misitebao/wails-template-vue"><img alt="GitHub Repo stars" src="https://img.shields.io/github/stars/misitebao/wails-template-vue?style=flat-square"/></a>
  <a href="https://github.com/misitebao"><img alt="GitHub user" src="https://img.shields.io/badge/author-misitebao-brightgreen?style=flat-square"/></a>
  <a href="https://github.com/wailsapp/wails"><img alt="Gitter" src="https://img.shields.io/badge/For-wails-brightgreen?style=flat-square&color=ff3c3c"/></a>
  <a href="https://github.com/misitebao/awesome-wails"><img alt="Awesome-Wails" src="https://cdn.jsdelivr.net/gh/sindresorhus/awesome@main/media/badge-flat.svg"/></a>
  <br/>
  <img alt="平台支持" src="https://img.shields.io/badge/support%20platform-windows%20%7C%20macos%20%7C%20linux-brightgreen?style=flat-square"/>
</p>

<span id="nav-1"></span>

## 国际化

[English](README.md) | [简体中文](README.zh-Hans.md)

<span id="nav-2"></span>

## 内容目录

<details>
  <summary>点我 打开/关闭 目录列表</summary>

- [国际化](#nav-1)
- [内容目录](#nav-2)
- [项目介绍](#nav-3)
  - [官方网站](#nav-3-1)
  - [背景](#nav-3-2)
- [图形演示](#nav-4)
- [功能特色](#nav-5)
- [架构](#nav-6)
- [新手入门](#nav-7)
- [维护者](#nav-8)
- [贡献者](#nav-9)
- [社区交流](#nav-10)
- [部分用户](#nav-11)
- [发布记录](CHANGE.md)
- [捐赠者](#nav-12)
- [赞助商](#nav-13)
- [特别感谢](#nav-14)
- [版权许可](#nav-15)

</details>

<span id="nav-3"></span>

## 项目介绍

wails-template-vue 模板是一个支持 Wails 应用的 Vue 模板，默认提供路由和国际化功能。

<span id="nav-3-1"></span>

<!-- ### 官方网站 -->

<span id="nav-3-2"></span>

### 背景

希望在 Wails 应用中使用 Vue 及其强大的社区生态开发支持多路由的单页应用。

<span id="nav-4"></span>

## 图形演示

[![演示截图](https://cdn.jsdelivr.net/gh/misitebao/CDN@main/preview/wails-template-vue-desktop.png "点击查看gif演示")](https://cdn.jsdelivr.net/gh/misitebao/CDN@main/preview/wails-template-vue.gif)

<span id="nav-5"></span>

## 功能特色

- 兼容 Windows、MacOS、Linux 平台
- 支持单页路由和国际化
- 内置 Sass 预处理器
- 跨平台一致的 UI 体验（内置 JetbrainsMono 字体包）
- 附带完整的 APi 示例（目前正在开发测试中。。。）

<span id="nav-6"></span>

<!-- ## 架构 -->

<span id="nav-7"></span>

## 新手入门

### 新建项目

```
wails init -n [你的应用名称] -t https://github.com/misitebao/wails-template-vue
```

参数说明：

- n - 将要创建的应用名称
- t - 模板名称，支持内置模板名称以及超链接形式的第三方模板

### 前端部分参考文档

前端部分使用了 Vue Vue-Router 和 Vue-I18N。

- Vue - 使用 Vue 3.x 版本，具体使用方式请参考[Vue3.x 官方文档](https://v3.vuejs.org/guide/introduction.html)。
- Vue-Router - 使用 Vue-Router 4.x 版本，具体使用方式请参考[Vue-Router 官方文档](https://next.router.vuejs.org/)。
- Vue-I18N - 使用 Vue-I18N 9.x 版本，具体使用方式请参考[Vue-I18N 官方文档](https://vue-i18n.intlify.dev/)。

然后您就可以参考 Wails 官方文档开始开发您的应用啦 🤞!

<span id="nav-8"></span>

## 维护者

感谢这些项目的维护者：

<a href="https://github.com/misitebao"><img src="https://github.com/misitebao.png" width="40" height="40" alt="misitebao" title="misitebao"/></a>

<details>
  <summary>点我 打开/关闭 贡献者列表</summary>

- [米司特包](https://github.com/misitebao) - 项目作者，全栈工程师。

</details>

<span id="nav-9"></span>

## 贡献者

感谢所有参与 wails-template-vue 开发的贡献者。[贡献者列表](https://github.com/misitebao/wails-template-vue/graphs/contributors)

<span id="nav-10"></span>

## 社区交流

- [Github Discussions](https://github.com/wailsapp/wails/discussions) - Github 官方推荐的 Wails 项目讨论区
- [Wails Slack](https://invite.slack.golangbridge.org/) - Wails 官方交流频道
- [Twitter](https://twitter.com/wailsapp) - Wails 官方推特账号

中文社区：

- <a target="_blank" href="https://qm.qq.com/cgi-bin/qm/qr?k=utlUvDwtcNG5knHBLwVdMvG39WeHh7oj&jump_from=webapi">QQ 群：1067173054</a> - QQ 中文社区交流群

<span id="nav-11"></span>

## 部分用户

- [Wails API Demos](https://github.com/misitebao/wails-api-demos) - 一个用于探索 Wails API 的示例程序，灵感来源自 [Electron API Demos](https://github.com/electron/electron-api-demos)

<span id="nav-12"></span>

<!-- ## 捐赠者 -->

<span id="nav-13"></span>

<!-- ## 赞助商 -->

<span id="nav-14"></span>

<!-- ## 特别感谢 -->

<span id="nav-15"></span>

## 版权许可

[License MIT](LICENSE)
