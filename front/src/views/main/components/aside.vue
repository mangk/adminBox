<template>
  <div>
    <div class="logo-box">
      <img class="logo" :src="logo" alt="" />
      <div class="logo-title" v-if="!isCollapse">{{ name }}</div>
    </div>
    <div class="desc" v-if="!isCollapse">{{ desc }}</div>
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
import MenuTree from '@/views/main/menuTree.vue'
import { onBeforeMount, ref } from 'vue'
import { useRouterStore } from '@/pinia/useRouterStore.js'

const logo = ref(window.adminBox.Logo ? window.adminBox.Logo : './images/logo.png')
const name = ref(window.adminBox.Name)
const desc = ref(window.adminBox.Desc ? window.adminBox.Desc : '')

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
  padding-bottom: 0;
  z-index: 2000;
  display: flex;
  flex-flow: row nowrap;
}

.logo {
  padding: 0;
}

.logo-title {
  font-family: emoji;
  margin-left: calc(var(--global-padding) / 2);
  /* display: inline-block; */
  color: #fff;
  font-weight: 600;
  font-size: 20px;
  font-style: italic;
}

.desc {
  color: #eee;
  font-size: 10px;
  font-weight: 200;
  margin-top: 5px;
  text-align: justify;
  text-align-last: justify;
  text-align: justify;
  text-justify: distribute-all-lines;
  padding: 0 var(--global-padding);
}

.el-menu {
  border-right: 0;
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
