<template>
  <el-container>
    <el-header v-if="!$route.meta.default_menu">
      <img class="logo" :src="logo" alt="" style="height: 40px" />
      <div style="display: flex; flex-flow: row nowrap; align-items: center">
        <div
          style="
            margin-right: 40px;
            font-size: 14px;
            display: flex;
            flex-flow: row nowrap;
            align-items: center;
          "
        >
          <el-icon style="margin-right: 5px; font-size: 16px">
            <Message />
          </el-icon>
          <div>通知</div>
        </div>
        <el-avatar :src="user.avatar" style="margin-right: 10px; width: 29px; height: 29px">
          <el-icon style="font-size: 20px">
            <UserFilled />
          </el-icon>
        </el-avatar>
        <div style="margin-right: 10px; font-size: 14px">{{ user.nick_name }}</div>
        <el-button link @click="logout" style="color: #9b9faa; font-size: 12px">退出</el-button>
      </div>
    </el-header>
    <el-container class="main">
      <el-aside v-if="!$route.meta.default_menu">
        <el-menu
          :default-active="$route.name"
          @open="handleOpen"
          @close="handleClose"
          :collapse="isCollapse"
          unique-opened
          router
        >
          <MenuTree :menus="menuList" />
        </el-menu>
        <Aside />
      </el-aside>
      <el-main>
        <Main />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import { ref } from 'vue'
import { useUserStore } from '@/pinia/useUserStore'
import { useRouterStore } from '@/pinia/useRouterStore.js'
import Main from '@/views/main/components/main.vue'
import MenuTree from '@/views/main/components/menuTree.vue'

const userStroe = useUserStore()
const user = ref({})
user.value = await userStroe.userInfo()

const logout = () => {
  userStroe.logOut()
}

const logo = ref(window.adminX.Logo ? window.adminX.Logo : './images/logo.png')

const menuList = await useRouterStore().loadServerRouter()
</script>

<style lang="scss">
html,
body {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  background-color: #eff2f8;
  width: 100%;
  height: 100%;
}

.el-header,
.el-aside,
.el-main {
  margin: 0;
  padding: 5px;
  overflow: hidden;
  overflow-y: scroll;
}

.el-header {
  --el-header-height: none;
  display: flex;
  flex-flow: row nowrap;
  justify-content: space-between;
}

.el-aside {
  width: auto;
}

.el-main {
  height: calc(100vh - 50px - 5px);
}

.el-menu.el-menu--vertical {
  width: 220px;
  min-height: 100%;
  border-radius: 12px;
  padding: 10px 10px;
  border-right: 0;
}

.el-menu.el-menu--inline {
  padding: auto 0px;
}

.el-sub-menu__title .el-sub-menu .el-menu-item,
.el-menu-item {
  border-radius: 6px;
  height: auto;
  line-height: 50px;
}

.el-menu-item:hover {
  background: linear-gradient(to right, #eff2f8, #fff);
}

.el-menu-item.is-active {
  color: #fff;
  background: linear-gradient(to right, #2473ff, #fff);
}

.app-content {
  height: 100%;
  background-color: #fff;
  border-radius: 12px;
}

:root {
  --el-color-primary: #2473ff;

  .el-button {
    border-radius: 3px;
  }
}
</style>
