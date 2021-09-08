<template>
  <div class="home">
    <!-- Logo -->
    <img class="logo" alt="Vue logo" src="../assets/logo.png" />
    <HelloWorld :msg="$t('homepage.welcome')" />
    <!-- Bottom button -->
    <!-- 底部按钮 -->
    <div class="link">
      <OpenLink href="https://beta.wails.io" class="btn start">{{ $t('homepage.getting-started') }}</OpenLink>
      <OpenLink
        href="https://github.com/misitebao/wails-template-vue"
        class="btn star"
      >{{ $t('homepage.star-me') }}</OpenLink>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from "@/components/HelloWorld.vue";

export default {
  name: "Home",
  components: {
    HelloWorld
  },
  setup() {
    return {};
  },
};
</script>
<style lang="stylus">
.home{
  .logo{
    display block
    width 620px
    height 280px
    margin 30px auto 10px
  }
  .link{
    display flex
    flex-direction row
    flex-wrap nowrap
    align-items center
    justify-content center
    margin 24px auto
    .btn{
      display block
      width 136px
      height 46px
      line-height 46px
      margin 0 10px
      border-radius 8px
      text-align center
      font-weight bold
      font-size: 14px;
      white-space nowrap
      text-decoration: none;
      cursor pointer
      &.start{
        background-color #fd0404
        color #ffffff
        &:hover{
          background-color #ec2e2e
        }
      }
      &.star{
        background-color #ffffff
        color #fd0404
        &:hover{
          background-color #f3f3f3
        }
      }
    }
  }
}
</style>
