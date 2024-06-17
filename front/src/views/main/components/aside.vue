<template>
  <div class="logo-box">
    <img class="logo" :src="logo" alt="" />
    <div class="logo-title" v-if="!isCollapse">{{ name }}</div>
  </div>
  <el-menu
    :default-active="$route.name"
    @open="handleOpen"
    @close="handleClose"
    :collapse="isCollapse"
    unique-opened
    router
    background-color="#191a23"
    text-color="#fff"
  >
    <MenuTree :menus="menuList" />
  </el-menu>
  <div @click="isCollapse = !isCollapse" class="collapse-change-btn">
    <el-icon>
      <ArrowRightBold v-if="isCollapse" />
      <ArrowLeftBold v-if="!isCollapse" />
    </el-icon>
  </div>
</template>

<script setup>
import MenuTree from '@/views/main/components/menuTree.vue'
import { onBeforeMount, ref } from 'vue'
import { useRouterStore } from '@/pinia/useRouterStore.js'

const logo = ref(window.adminX.Logo ? window.adminX.Logo : './images/logo.png')
const name = ref(window.adminX.Name)

const isCollapse = ref(false)

const menuList = await useRouterStore().loadServerRouter()

const handleOpen = (key, keyPath) => {}
const handleClose = (key, keyPath) => {}

onBeforeMount(() => {
  isCollapse.value = document.body.clientWidth < 1200
})

window.onresize = () => {
  return (() => {
    isCollapse.value = document.body.clientWidth < 1200
  })()
}
</script>

<style scoped>
.logo-box {
  position: sticky;
  top: 0;
  width: 100%;
  height: calc(var(--global-header-height) - var(--global-padding) * 2);
  line-height: calc(var(--global-header-height) - var(--global-padding) * 2);
  padding: var(--global-padding);
  z-index: 2000;
  display: flex;
  flex-flow: row nowrap;
}

.logo-title {
  margin-left: var(--global-padding);
  display: inline-block;
  color: #fff;
  font-weight: 600;
  font-size: 20px;
}

.logo {
  width: calc(var(--global-header-height) - var(--global-padding) * 2.5);
  height: calc(var(--global-header-height) - var(--global-padding) * 2.5);
  margin-top: calc(var(--global-padding) * 0.5);
  border-radius: 50%;
  overflow: hidden;
  border: 3px solid #fff;
  box-sizing: border-box;
}

.el-menu {
  border-right: 0;
  max-width: 240px;
  overflow-x: hidden;
}

.collapse-change-btn {
  position: absolute;
  color: #fff;
  bottom: 0;
  width: 100%;
  /*width: 100%;*/
  height: 25px;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
