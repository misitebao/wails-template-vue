<template>
  <div class="about">
    <!-- Title -->
    <div class="title">{{ $t('aboutpage.title') }}</div>
    <!-- Information -->
    <!-- 信息 -->
    <ul class="info">
      <li class="info-item">
        <span class="name">{{ $t('aboutpage.project-repository') + ': ' }}</span>
        <OpenLink
          class="link"
          href="https://github.com/misitebao/wails-template-vue"
        >https://github.com/misitebao/wails-template-vue</OpenLink>
      </li>
      <li class="info-item">
        <span class="name">{{ $t('aboutpage.author') + ': ' }}</span>
        <OpenLink class="link" href="https://github.com/misitebao">{{ $t('aboutpage.misitebao') }}</OpenLink>
      </li>
      <li class="info-item">
        <span class="name">{{ $t('aboutpage.wails-repository') + ': ' }}</span>
        <OpenLink
          class="link"
          href="https://github.com/wailsapp/wails"
        >https://github.com/wailsapp/wails</OpenLink>
      </li>
    </ul>
    <!-- Thanks -->
    <!-- 谢语 -->
    <div class="thank">{{ $t('aboutpage.thanks') }}</div>
  </div>
</template>
<style lang="scss">
.about {
  .title {
    margin: 30px auto;
    font-size: 36px;
    color: #a150b5;
    text-align: center;
  }
  .info {
    margin: 14px 60px;
    // background-color red
    font-size: 24px;
    .info-item {
      height: 60px;
      line-height: 60px;
      white-space: nowrap;
      .name {
        color: #6d6363;
      }
      .link {
        color: #5f6c86;
      }
    }
  }
  .thank {
    height: 68px;
    line-height: 68px;
    margin: 47px auto;
    text-align: center;
    font-size: 50px;
  }
}
</style>
