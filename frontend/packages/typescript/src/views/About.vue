<template>
  <div class="about">
    <!-- Title -->
    <div class="title">{{ "Wails Template Vue" }}</div>
    <!-- Information -->
    <!-- 信息 -->
    <div class="content">
      <div class="comeon">
        <img :src="comeonGif" alt />
      </div>
      <ul class="info">
        <li class="info-item">
          <div class="name">{{ "Project Repository" }}</div>
          <OpenLink
            class="link"
            href="https://github.com/misitebao/wails-template-vue"
            >https://github.com/misitebao/wails-template-vue</OpenLink
          >
        </li>
        <li class="info-item">
          <div class="name">{{ "Wails Repository" }}</div>
          <OpenLink class="link" href="https://github.com/wailsapp/wails"
            >https://github.com/wailsapp/wails</OpenLink
          >
        </li>
        <li class="info-item">
          <div class="name">{{ "Author" }}</div>
          <OpenLink class="link" href="https://github.com/misitebao">{{
            "Misitebao"
          }}</OpenLink>
        </li>
      </ul>
    </div>

    <!-- Thanks -->
    <!-- 谢语 -->
    <div class="thank">{{ "Thank you all for your support🙏!" }}</div>
  </div>
</template>
<script lang="ts">
import comeonGif from "@/assets/images/comeon.gif";
export default {
  setup() {
    return {
      comeonGif,
    };
  },
};
</script>
<style lang="scss">
.about {
  .title {
    margin: 30px auto 10px;
    font-size: 38px;
    color: #a150b5;
    text-align: center;
  }
  .content {
    position: relative;
    margin: 36px 20px;
    .comeon {
      position: absolute;
      left: 26px;
      top: 38px;
      max-width: 66%;
      img {
        width: 220px;
        height: 180px;
      }
    }
    .info {
      margin: 0 0 0 33%;
      font-size: 24px;
      text-align: left;
      .info-item {
        margin-bottom: 10px;
        .name {
          line-height: 40px;
          font-size: 28px;
          color: #6d6363;
        }
        .link {
          line-height: 30px;
          font-size: 20px;
          color: #5f6c86;
        }
      }
    }
  }

  .thank {
    height: 68px;
    line-height: 68px;
    margin: 36px auto;
    text-align: center;
    font-size: 40px;
  }
}
</style>
