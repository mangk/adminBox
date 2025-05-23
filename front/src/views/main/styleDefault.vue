<template>
  <el-container style="width: 100%; height: 100%">
    <el-header class="box-header" :style="{
      'background-color': darkSidebar ? darkSidebarColor : '',
      color: darkSidebar ? '#fff' : ''
    }">
      <div class="header-logo-box" :style="{
        boxShadow: headerMenu || mobileDevice ? '2px 0 4px rgba(0, 0, 0, 0.16)' : 'unset'
      }">
        <img class="header-logo" :src="logo" />
        <div class="header-name" v-if="!mobileDevice">{{ name }}</div>
      </div>
      <el-scrollbar style="height: var(--box-header-height); margin: 0 8px">
        <el-menu v-if="headerMenu || mobileDevice" class="header-menu" :default-active="$route.name" @open="handleOpen"
          @close="handleClose" :collapse="isCollapse" unique-opened router mode="horizontal" :ellipsis="false"
          :background-color="darkSidebar ? darkSidebarColor : ''" :text-color="darkSidebar ? '#fff' : ''"
          style="width: auto">
          <MenuTree :menus="menuList" />
        </el-menu>
      </el-scrollbar>

      <el-popover>
        <template #reference>
          <div class="left header-user" :style="{
            boxShadow: headerMenu || mobileDevice ? '-2px 0 4px rgba(0, 0, 0, 0.16)' : 'unset'
          }">
            <span v-if="!mobileDevice">{{ user.nick_name }}</span>
            <el-avatar :src="user.avatar" icon="UserFilled" :size="30" style="margin-left: 5px" />
          </div>
        </template>
        <template #default>
          <div class="demo-rich-conent" style="display: flex; gap: 6px; flex-direction: column">
            <div v-if="mobileDevice">
              {{ user.nick_name }}
              <el-divider style="margin: 0;" />
            </div>
            <div>
              <el-tag v-for="color in colors" :key="color" :color="color.value"
                @click="setThemeElColorPrimary(color.value)" />
            </div>
            <el-divider style="margin: 0;" />
            <div style="display: flex; justify-content: space-between;align-items: center;">
              顶部菜单<el-switch v-model="headerMenu" @change="setHeaderMenu" :disabled="mobileDevice" />
            </div>
            <el-divider style="margin: 0;" />
            <div style="display: flex; justify-content: space-between;align-items: center;">
              深色边栏<el-switch v-model="darkSidebar" @change="setDarkSidebar" />
            </div>
            <el-divider style="margin: 0;" />
            <el-button link @click="changePasswordDialog = true" style="width: 100%;">修改密码</el-button>
            <el-divider style="margin: 0;" />
            <el-button link @click="logout">退出</el-button>
          </div>
        </template>
      </el-popover>
    </el-header>
    <el-container class="box-asside-and-main">
      <el-aside class="box-aside" v-if="!headerMenu && !mobileDevice" :style="{
        'background-color': darkSidebar ? darkSidebarColor : '',
        color: darkSidebar ? '#fff' : ''
      }">
        <el-scrollbar>
          <transition :duration="{ enter: 800, leave: 100 }" mode="out-in" name="el-fade-in-linear">
            <el-menu :default-active="$route.name" @open="handleOpen" @close="handleClose" :collapse="isCollapse"
              unique-opened router :background-color="darkSidebar ? darkSidebarColor : ''"
              :text-color="darkSidebar ? '#fff' : ''">
              <MenuTree :menus="menuList" />
            </el-menu>
          </transition>
        </el-scrollbar>
      </el-aside>
      <el-main class="box-main">
        <el-scrollbar>
          <router-view v-slot="{ Component }">
            <transition :duration="{ enter: 800, leave: 100 }" mode="out-in" name="el-fade-in-linear">
              <component :is="Component" :key="$route.path" />
            </transition>
          </router-view>
        </el-scrollbar>
      </el-main>
    </el-container>
  </el-container>

  <el-dialog v-model="changePasswordDialog" title="修改密码" width="500">
    <el-form label-position="right" label-width="auto">
      <el-form-item label="当前密码" label-position="right">
        <el-input v-model="passwordForm.old" type="password" autocomplete="off" />
      </el-form-item>
      <el-form-item label="新密码" label-position="right">
        <el-input v-model="passwordForm.new" type="password" autocomplete="off" />
      </el-form-item>
      <el-form-item label="重复新密码" label-position="right">
        <el-input v-model="passwordForm.confirmNew" type="password" autocomplete="off" />
      </el-form-item>
    </el-form>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="closeChangePassword">取消</el-button>
        <el-button type="primary" @click="changePassword">
          确认修改
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import { ref, onBeforeMount, reactive, getCurrentInstance } from 'vue'
import { useUserStore } from '@/pinia/useUserStore'
import { useRouterStore } from '@/pinia/useRouterStore.js'
import { useCssVar } from '@vueuse/core'

// theme
const cfg = reactive(
  localStorage.getItem('x-theme-config')
    ? JSON.parse(localStorage.getItem('x-theme-config'))
    : { darkSidebar: true }
)

const mobileDevice = ref(false)
const headerMenu = ref(cfg.headerMenu)
const setHeaderMenu = (v) => {
  cfg.headerMenu = v
  localStorage.setItem('x-theme-config', JSON.stringify(cfg))
}

const darkSidebar = ref(cfg.darkSidebar)
const setDarkSidebar = (v) => {
  cfg.darkSidebar = v
  localStorage.setItem('x-theme-config', JSON.stringify(cfg))
}

const themeElColorPrimary = useCssVar('--el-color-primary')
if (cfg.colorPrimary) {
  themeElColorPrimary.value = cfg.colorPrimary
}
const setThemeElColorPrimary = (v) => {
  themeElColorPrimary.value = v
  cfg.colorPrimary = v
  localStorage.setItem('x-theme-config', JSON.stringify(cfg))
}

const darkSidebarColor = '#2d2d32'
const colors = [
  {
    value: '#4167F0',
    label: 'blue'
  },
  {
    value: '#E63415',
    label: 'red'
  },
  {
    value: '#e0620e',
    label: 'orange'
  },
  {
    value: '#1EC79D',
    label: 'green'
  },
  {
    value: '#6222C9',
    label: 'purple'
  },
  {
    value: '#000',
    label: 'black'
  }
]

// header
const logo = ref(window.adminBox.Logo ? window.adminBox.Logo : './images/logo.svg')
const name = ref(window.adminBox.Name)
const userStroe = useUserStore()
const user = ref({})
user.value = await userStroe.userInfo()
const logout = () => {
  userStroe.logOut()
}

// aside
const isCollapse = ref(false)

const menuList = await useRouterStore().loadServerRouter()

const handleOpen = (key, keyPath) => { }
const handleClose = (key, keyPath) => { }
const setCollapse = () => {
  if (document.body.clientWidth >= 1100) {
    isCollapse.value = false
  } else {
    isCollapse.value = true
  }
  mobileDevice.value = document.body.clientWidth <= 550
}

const { proxy } = getCurrentInstance()

const changePasswordDialog = ref(false)
const closeChangePassword = () => {
  changePasswordDialog.value = false
}
const passwordForm = ref({ old: '', new: '', confirmNew: '' })
const changePassword = () => {

  if (!passwordForm.value.old) {
    proxy.$message.error('当前密码 不能为空')
    return
  }
  if (!passwordForm.value.new) {
    proxy.$message.error('新密码 不能为空')
    return
  }
  if (!passwordForm.value.confirmNew) {
    proxy.$message.error('重复新密码 不能为空')
    return
  }
  if (passwordForm.value.old == passwordForm.value.new) {
    proxy.$message.error('新密码不能与旧密码相同')
    return
  }
  if (passwordForm.value.new !== passwordForm.value.confirmNew) {
    proxy.$message.error('新密码两次输入不一致')
    return
  }

  proxy.$http({
    url: "/sys/setting/user/changePassword",
    method: "PUT",
    data: passwordForm.value
  }).then(res => {
    if (res.code == 0) {
      changePasswordDialog.value = false
      proxy.$message.success('修改完成，下次登录请使用新密码')
    }
  })
}



onBeforeMount(setCollapse)
window.onresize = () => {
  return setCollapse()
}
</script>

<style>
.el-container,
.el-header,
.el-aside,
.el-main {
  margin: 0;
  padding: 0;
  border: 0;
  position: relative;
}

.el-menu {
  border-right: 0;
}

.el-menu--collapse {
  width: auto;
}

.box-header {
  position: sticky;
  z-index: 3;
  height: var(--box-header-height);
  font-size: 12px;
  display: flex;
  flex-flow: row nowrap;
  align-items: center;
  box-sizing: border-box;
  box-shadow: 0 2px 4px 0 rgba(0, 0, 0, 0.16);
}

.header-logo-box {
  height: 100%;
  display: flex;
  flex-flow: row nowrap;
  justify-content: start;
  align-items: center;
  padding: var(--global-padding);
  box-shadow: 2px 0 4px rgba(0, 0, 0, 0.16);
  box-sizing: border-box;
}

.header-logo {
  flex-grow: 0;
  box-sizing: border-box;
  height: 100%;
}

.header-name {
  flex-grow: 0;
  margin-left: calc(var(--global-padding) / 2);
  font-weight: 550;
  font-size: 20px;
}

.header-menu {
  --el-menu-item-height: var(--global-padding);
  flex-grow: 1;
  height: calc(var(--box-header-height));
  font-size: 12px;
  border-bottom: 0px;

  .el-menu--horizontal.el-menu {
    border-bottom: 0px;
  }

  .el-menu--horizontal>.el-menu-item {
    border-bottom: 0px;
  }

  .el-menu--horizontal .el-menu {
    border-bottom: 0px;
  }
}

.header-user {
  display: flex;
  flex-flow: row nowrap;
  align-items: center;
  padding: var(--global-padding);
}

.left {
  margin-left: auto;
}

.box-asside-and-main {
  height: calc(100vh - var(--box-header-height));
  overflow: hidden;
}

.box-aside {
  position: relative;
  z-index: 2;
  width: auto;
  box-shadow: 1px 2px 4px 0 rgba(0, 0, 0, 0.16);
}

.box-main {
  position: relative;
  z-index: 1;
  overflow-x: hidden;
  background-color: #f0f2f5;
}
</style>
