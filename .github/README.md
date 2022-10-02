<h1 align="center">Wails Template Vue</h1>

<p align="center">
  <img src="./logo.png" height="280" />
</p>

<p align="center">
  Wails template based on Vue ecology
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

## Table of Contents

<details>
  <summary>Click me to Open/Close the directory listing</summary>

- [Table of Contents](#table-of-contents)
- [Introductions](#introductions)
  - [Official Website](#official-website)
  - [Background](#background)
- [Graphic Demo](#graphic-demo)
- [Features](#features)
- [Architecture](#architecture)
- [Getting Started](#getting-started)
- [Maintainer](#maintainer)
- [Contributors](#contributors)
- [Community Exchange](#community-exchange)
- [Part Of Users](#part-of-users)
- [Changelog](#changelog)
- [Donators](#donators)
- [Sponsors](#sponsors)
- [Special Thanks](#special-thanks)
- [License](#license)

</details>

## Introductions

This project is a [Wails](https://github.com/wailsapp/wails) template based on the [Vue](https://vuejs.org/) ecosystem. You can use this project to quickly develop your application.

### Background

The official built-in template only provides the minimum content required to develop an application. If you want to develop an application quickly, you still need to do a lot of work yourself. Based on this, this template project provides you with out-of-the-box functions, you can Get started developing your app quickly.

## Graphic Demo

[![Demo Screenshot](https://cdn.jsdelivr.net/gh/misitebao/wails-template-vue@main/.github/preview.en.png "click to view gif demo")](https://cdn.jsdelivr.net/gh/misitebao/wails-template-vue@main/.github/preview.gif)

## Features

- Single page routing support
- Built-in internationalization
- Built-in bright and dark themes
- Built-in FontAwesome icon library
- Integrated TailwindCSS
- Integrate TypeScript
- Perfect for Windows, MacOS, Linux platforms

## Getting Started

**New Project**:

```
wails init -n <Your Appname> -t https://github.com/misitebao/wails-template-vue[@version]
```

flag description:

- n - The name of the application to be created
- t - Template name, supports built-in template names and third-party templates in the form of hyperlinks
- @version - Specify the version of a specific Git Tag, if no version is specified, the code of the master branch will be used by default

**Reference document**:

The front-end part uses Vue Vue-Router and Vue-I18N:

- Vue - Use vue3.x version, please refer to the official [Vue3.x Documents](https://vuejs.org/) for specific usage.
- Vue-Router - Use Vue-Router 4.x version, please refer to the official [Vue-Router Documents](https://router.vuejs.org/) for specific usage.
- Vue-I18N - Use Vue-I18N 9.x version, please refer to official [Vue-I18N Documents](https://vue-i18n.intlify.dev/) for specific usage.
- TypeScript - Please refer to official [TypeScript Documents](https://www.typescriptlang.org/) for specific usage.
- TailwindCSS - Please refer to official [TailwindCSS Documents](https://tailwindcss.com/) for specific usage.

Then you can refer to the official [Wails document](https://wails.io) to start developing your application🤞.

## Maintainer

Thanks to the maintainers of these projects:

<a href="https://github.com/misitebao">
  <img src="https://github.com/misitebao.png" width="40" height="40" alt="misitebao" title="misitebao"/>
</a>

## Contributors

Thanks to all the contributors involved in the development. [Contributors](https://github.com/misitebao/wails-template-vue/graphs/contributors)

<a href="https://github.com/eighteenzheng">
  <img src="https://github.com/eighteenzheng.png" width="40" height="40" alt="eighteenzheng" title="eighteenzheng"/>
</a>
<a href="https://github.com/daoif">
  <img src="https://github.com/daoif.png" width="40" height="40" alt="daoif" title="daoif"/>
</a>

## Community Exchange

- [Github Discussions](https://github.com/wailsapp/wails/discussions) - The official Github communication community of the Wails project
- [Wails Slack](https://invite.slack.golangbridge.org/) - Wails official communication channel
- [Twitter](https://twitter.com/wailsapp) - Wails official Twitter account

Chinese Community:

- <a target="_blank" href="https://qm.qq.com/cgi-bin/qm/qr?k=utlUvDwtcNG5knHBLwVdMvG39WeHh7oj&jump_from=webapi">QQ Group: 1067173054</a> - QQ Chinese Community Exchange Group

## Part Of Users

- [Wails API Demos](https://github.com/misitebao/wails-api-demos) - A sample program for exploring Wails API, inspired by [Electron API Demos](https://github.com/electron/electron-api-demos)
- [Hayate](https://github.com/misitebao/hayate) - Windows application installation boot program implemented using Wails.

## Changelog

Check out the changelog [here](./CHANGELOG.md).

## License

[License MIT](./LICENSE)
