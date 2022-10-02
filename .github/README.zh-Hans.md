<h1 align="center">Wails Template Vue</h1>

<p align="center">
  <img src="./logo.png" height="280" />
</p>

<p align="center">
  基于 Vue 生态的 Wails 模板
</p>

<p align="center">
  <a href="https://github.com/misitebao/wails-template-vue/blob/main/LICENSE">
    <img alt="GitHub" src="https://img.shields.io/github/license/misitebao/wails-template-vue"/>
  </a>
  <a href="https://github.com/misitebao/standard-repository">
    <img alt="GitHub" src="https://cdn.jsdelivr.net/gh/misitebao/standard-repository@main/assets/badge_flat.svg"/>
  </a>
  <a href="https://github.com/misitebao/wails-template-vue/releases">
    <img alt="GitHub release (latest by date including pre-releases)" src="https://img.shields.io/github/v/release/misitebao/wails-template-vue?include_prereleases&sort=semver">
  </a>
  <a href="https://github.com/wailsapp/awesome-wails">
    <img alt="Awesome-Wails" src="https://cdn.jsdelivr.net/gh/sindresorhus/awesome@main/media/badge.svg"/>
  </a>
  <br/>
  <img src="https://img.shields.io/badge/platform-windows%20%7C%20macos%20%7C%20linux-brightgreen?"/>
</p>

<div align="center">
<strong>
<samp>

[English](README.md) · [简体中文](README.zh-Hans.md)

</samp>
</strong>
</div>

## 内容目录

<details>
  <summary>点我 打开/关闭 目录列表</summary>

- [国际化](#国际化)
- [内容目录](#内容目录)
- [项目介绍](#项目介绍)
  - [官方网站](#官方网站)
  - [背景](#背景)
- [图形演示](#图形演示)
- [功能](#功能)
- [架构](#架构)
- [快速入门](#快速入门)
- [维护者](#维护者)
- [贡献者](#贡献者)
- [社区交流](#社区交流)
- [部分用户](#部分用户)
- [更新日志](#更新日志)
- [捐赠者](#捐赠者)
- [赞助商](#赞助商)
- [特别感谢](#特别感谢)
- [许可证](#许可证)

</details>

## 项目介绍

本项目是一个基于 [Vue](https://vuejs.org/) 生态的 [Wails](https://github.com/wailsapp/wails) 模板，使用此项目可以快速开发您的应用。

### 背景

官方内置的模板仅仅提供了开发应用所需的最少的内容，如果你想快速开发应用，则仍然需要自己做很多工作，基于此，本模板项目为您提供了开箱即用的功能，您可以快速开始开发您的应用。

## 图形演示

[![演示截图](https://cdn.jsdelivr.net/gh/misitebao/wails-template-vue@main/.github/preview.zh-Hans.png "点击查看gif演示")](https://cdn.jsdelivr.net/gh/misitebao/wails-template-vue@main/.github/preview.gif)

## 功能

- 单页路由支持
- 内置国际化
- 内置明亮、暗黑主题
- 内置 FontAwesome 图标库
- 集成 TailwindCSS
- 集成 TypeScript
- 完美适配 Windows、MacOS、Linux 平台

## 快速入门

**新建项目**:

```
wails init -n <你的应用名称> -t https://github.com/misitebao/wails-template-vue[@version]
```

参数说明：

- n - 将要创建的应用名称
- t - 模板名称，支持内置模板名称以及超链接形式的第三方模板
- @version - 指定特定 Git Tag 的版本，如果不指定版本，则默认使用主分支的代码

**参考文档**：

前端部分使用了 Vue、Vue-Router、Vue-I18N、TypeScript 和 TailwindCSS。

- Vue - 使用 Vue 3.x 版本，具体使用方式请参考[Vue3.x 官方文档](https://vuejs.org/)。
- Vue-Router - 使用 Vue-Router 4.x 版本，具体使用方式请参考[Vue-Router 官方文档](https://router.vuejs.org/)。
- Vue-I18N - 使用 Vue-I18N 9.x 版本，具体使用方式请参考[Vue-I18N 官方文档](https://vue-i18n.intlify.dev/)。
- TypeScript - 具体使用方式请参考[TypeScript 官方文档](https://www.typescriptlang.org/)。
- TailwindCSS - 具体使用方式请参考[TailwindCSS 官方文档](https://tailwindcss.com/)

然后您就可以参考 [Wails 官方文档](https://wails.io)开始开发您的应用啦 🤞!

## 维护者

感谢这些项目的维护者：

<a href="https://github.com/misitebao">
  <img src="https://github.com/misitebao.png" width="40" height="40" alt="misitebao" title="misitebao"/>
</a>

## 贡献者

感谢所有参与开发的贡献者。[贡献者列表](https://github.com/misitebao/wails-template-vue/graphs/contributors)

<a href="https://github.com/eighteenzheng">
  <img src="https://github.com/eighteenzheng.png" width="40" height="40" alt="eighteenzheng" title="eighteenzheng"/>
</a>
<a href="https://github.com/daoif">
  <img src="https://github.com/daoif.png" width="40" height="40" alt="daoif" title="daoif"/>
</a>

## 社区交流

- [Github Discussions](https://github.com/wailsapp/wails/discussions) - Github 官方推荐的 Wails 项目讨论区
- [Wails Slack](https://invite.slack.golangbridge.org/) - Wails 官方交流频道
- [Twitter](https://twitter.com/wailsapp) - Wails 官方推特账号

中文社区：

- <a target="_blank" href="https://qm.qq.com/cgi-bin/qm/qr?k=utlUvDwtcNG5knHBLwVdMvG39WeHh7oj&jump_from=webapi">QQ 群：1067173054</a> - QQ 中文社区交流群

## 部分用户

- [Wails API Demos](https://github.com/misitebao/wails-api-demos) - 一个用于探索 Wails API 的示例程序，灵感来源自 [Electron API Demos](https://github.com/electron/electron-api-demos)
- [Hayate](https://github.com/misitebao/hayate) - 使用 Wails 实现的 Windows 应用安装引导程序。

## 更新日志

更新日志请查看[这里](./CHANGELOG.md)。

## 许可证

[License MIT](./LICENSE)
