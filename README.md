<p align="center">
  <img src="./.github/logo.gif" height="280" />
</p>
<p align="center">
  A Wails template based on Vue and Vue-Router
</p>
<p align="center">
  <a href="https://github.com/misitebao/wails-template-vue/blob/main/LICENSE">
    <img alt="GitHub" src="https://img.shields.io/github/license/misitebao/wails-template-vue?style=flat-square"/>
  </a>
  <a href="https://github.com/misitebao/standard-repository">
    <img alt="GitHub" src="https://cdn.jsdelivr.net/gh/misitebao/standard-repository@main/assets/badge_flat-square.svg"/>
  </a>
  <a href="https://github.com/misitebao/wails-template-vue">
    <img alt="GitHub Repo stars" src="https://img.shields.io/github/stars/misitebao/wails-template-vue?style=flat-square"/>
  </a>
  <a href="https://github.com/misitebao/wails-template-vue/releases">
    <img alt="GitHub release (latest by date including pre-releases)" src="https://img.shields.io/github/v/release/misitebao/wails-template-vue?include_prereleases&sort=semver&style=flat-square">
  </a>
  <a href="https://github.com/misitebao">
    <img alt="GitHub user" src="https://img.shields.io/badge/author-misitebao-brightgreen?style=flat-square"/>
  </a>
  <a href="https://github.com/wailsapp/wails">
    <img alt="Gitter" src="https://img.shields.io/badge/For-Wails-brightgreen?style=flat-square&color=ff3c3c"/>
  </a>
  <a href="https://github.com/wailsapp/awesome-wails">
    <img alt="Awesome-Wails" src="https://cdn.jsdelivr.net/gh/sindresorhus/awesome@main/media/badge-flat.svg"/>
  </a>
  <img src="https://img.shields.io/badge/platform-windows%20%7C%20macos%20%7C%20linux-brightgreen?style=flat-square"/>
</p>

<span id="nav-1"></span>

## Internationalization

[English](README.md) | [简体中文](./.github/README.zh-Hans.md)

<span id="nav-2"></span>

## Table of Contents

<details>
  <summary>Click me to Open/Close the directory listing</summary>

- [Internationalization](#nav-1)
- [Table of Contents](#nav-2)
- [Introductions](#nav-3)
  - [Official Website](#nav-3-1)
  - [Background](#nav-3-2)
- [Graphic Demo](#nav-4)
- [Features](#nav-5)
- [Architecture](#nav-6)
- [Getting Started](#nav-7)
- [Maintainer](#nav-8)
- [Contributors](#nav-9)
- [Community Exchange](#nav-10)
- [Part Of Users](#nav-11)
- [Release History](CHANGE.md)
- [Donators](#nav-12)
- [Sponsors](#nav-13)
- [Special Thanks](#nav-14)
- [License](#nav-15)

</details>

<span id="nav-3"></span>

## Introductions

The wails-template-vue template is a Vue template that supports [Wails](https://github.com/wailsapp/wails) programs and provides Router and i18n functions by default.

<span id="nav-3-1"></span>

<!-- ### Official Website -->

<span id="nav-3-2"></span>

### Background

I hope that Vue and its powerful community ecology can be used in Wails applications to develop single-page applications that support multiple routes.

<span id="nav-4"></span>

## Graphic Demo

[![Demo Screenshot](https://cdn.jsdelivr.net/gh/misitebao/wails-template-vue@main/.github/preview.png "click to view gif demo")](https://cdn.jsdelivr.net/gh/misitebao/wails-template-vue@main/.github/preview.gif)

<span id="nav-5"></span>

## Features

- Compatible with Windows, MacOS, Linux platforms
- Support single page routing and i18n.
- Built-in Sass preprocessor.
- A consistent UI experience across platforms(Comes with JetbrainsMono font package).
- Comes with a complete API example(Currently under development and testing...).
- Support JavaScript and TypeScript

<span id="nav-6"></span>

<!-- ## Architecture -->

<span id="nav-7"></span>

## Getting Started

### New Project

```
wails init -n [Your Appname] -t https://github.com/misitebao/wails-template-vue
```

flag description:

- n - The name of the application to be created
- t - Template name, supports built-in template names and third-party templates in the form of hyperlinks

After the project is created, the JavaScript template is used by default. You can change the value of the `"frontend:build"` field in the `wails.json` file to `"npm run build -w ts"` to use the TypeScript template.

Note: In order to support TypeScript templates under the existing functions, the front-end part uses the NPM workspace function, so the NPM version must be `>=7.0.0`, please run `npm -v` to check your NPM version.

### Reference document

The front-end part uses Vue Vue-Router and Vue-I18N:

- Vue - Use vue3.x version, please refer to the official [Vue3.x Documents](https://v3.vuejs.org/guide/introduction.html) for specific usage.
- Vue-Router - Use Vue-Router 4.x version, please refer to the official [Vue-Router Documents](https://next.router.vuejs.org/) for specific usage.
- Vue-I18N - Use Vue-I18N 9.x version, please refer to official [Vue-I18N Documents](https://vue-i18n.intlify.dev/) for specific usage.
- TypeScript - Please refer to official [TypeScript Documents](https://www.typescriptlang.org/) for specific usage.

Then you can refer to the official [Wails document](https://wails.io) to start developing your application🤞.

<span id="nav-8"></span>

## Maintainer

Thanks to the maintainers of these projects:

<a href="https://github.com/misitebao"><img src="https://github.com/misitebao.png" width="40" height="40" alt="misitebao" title="misitebao"/></a>

<details>
  <summary>Click me to Open/Close the maintainer listing</summary>

- [Misitebao](https://github.com/misitebao) - Project author, full stack engineer.

</details>

<span id="nav-9"></span>

## Contributors

Thank you to all the contributors who participated in the development of wails-template-vue. [Contributors](https://github.com/misitebao/wails-template-vue/graphs/contributors)

<a href="https://github.com/crushonyou18"><img src="https://github.com/crushonyou18.png" width="40" height="40" alt="crushonyou18" title="crushonyou18"/></a>

<span id="nav-10"></span>

## Community Exchange

- [Github Discussions](https://github.com/wailsapp/wails/discussions) - The official Github communication community of the Wails project
- [Wails Slack](https://invite.slack.golangbridge.org/) - Wails official communication channel
- [Twitter](https://twitter.com/wailsapp) - Wails official Twitter account

Chinese Community:

- <a target="_blank" href="https://qm.qq.com/cgi-bin/qm/qr?k=utlUvDwtcNG5knHBLwVdMvG39WeHh7oj&jump_from=webapi">QQ Group: 1067173054</a> - QQ Chinese Community Exchange Group

<span id="nav-11"></span>

## Part Of Users

- [Wails API Demos](https://github.com/misitebao/wails-api-demos) - A sample program for exploring Wails API, inspired by [Electron API Demos](https://github.com/electron/electron-api-demos)
- [Hayate](https://github.com/misitebao/hayate) - Windows application installation boot program implemented using Wails.

<span id="nav-12"></span>

<!-- ## Donators -->

<span id="nav-13"></span>

<!-- ## Sponsors -->

<span id="nav-14"></span>

<!-- ## Special Thanks -->

<span id="nav-15"></span>

## License

[License MIT](./LICENSE)
