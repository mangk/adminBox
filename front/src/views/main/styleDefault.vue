<template>
  <el-container class="app-box">
    <el-container class="app-aside-and-main">
      <el-aside class="app-aside" v-if="!$route.meta.default_menu">
        <Aside />
      </el-aside>
      <el-main :class="['app-main', $route.meta.default_menu ? 'default-page' : '']">
        <el-header class="app-header" v-if="!$route.meta.default_menu">
          <Header />
        </el-header>
        <Main />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import Aside from '@/views/main/components/aside.vue'
import Header from '@/views/main/components/header.vue'
import Main from '@/views/main/components/main.vue'
</script>

<style lang="scss" scoped>
@import '@/assets/globalSet.scss';

.el-container,
.el-header,
.el-aside,
.el-main {
  margin: 0;
  padding: 0;
  border: 0;
}

.app-box {
  display: flex;
  flex-direction: column;
  width: 100vw;
  height: 100vh;
}

.app-header {
  width: 100%;
  height: unset;
  z-index: 10;
}

.app-aside-and-main {
  flex-grow: 1;
  overflow: hidden;
}

.app-aside {
  position: relative;
  scrollbar-width: none;
  width: unset;
  margin: 0 var(--global-interval) var(--global-interval) var(--global-interval);
  border-radius: var(--global-border-radius) 0 0 var(--global-border-radius);
  background-color: #191a23;
}

.app-main {
  scrollbar-width: none;
  margin: 0 var(--global-interval) var(--global-interval) 0;
  border-radius: 0 var(--global-border-radius) var(--global-border-radius) 0;
}

.default-page {
  margin: 0 !important;
  padding: 0 !important;
  border-radius: 0 !important;
}

@media screen and (max-width: 1000px) {
  .app-aside {
    position: fixed;
    z-index: 10;
    top: 0;
    height: 100%;
    left: -100px;
  }
  .app-main {
    margin-left: var(--global-interval);
    border-radius: var(--global-border-radius);
  }
}
</style>
