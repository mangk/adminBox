<template>
  <div class="in-header">
    <el-breadcrumb separator-icon="ArrowRight">
      <el-breadcrumb-item v-for="item in matched.slice(1, matched.length)" :key="item.path"
        >{{ fmtTitle(item.meta.title, route) }}
      </el-breadcrumb-item>
    </el-breadcrumb>
    <div style="display: flex; justify-content: space-between">
      <div class="username">{{ user.username }}</div>
      <el-popover :width="80">
        <template #reference>
          <el-avatar :src="user.avatar" @error="errorHandler">
            <el-icon style="font-size: 20px">
              <UserFilled />
            </el-icon>
          </el-avatar>
        </template>
        <template #default>
          <div class="demo-rich-conent" style="display: flex; gap: 16px; flex-direction: column">
            <el-button link @click="logout">退出</el-button>
          </div>
        </template>
      </el-popover>
    </div>
  </div>
  <Tab v-if="!$route.meta.default_menu" />
</template>

<script setup>
import Tab from '@/views/main/components/tab.vue'
import { useRoute } from 'vue-router'
import { fmtTitle } from '@/utils/fmtRouterTitle'
import { computed, ref } from 'vue'
import { useUserStore } from '@/pinia/useUserStore'
const route = useRoute()
const matched = computed(() => route.matched)

const errorHandler = () => true

const userStroe = useUserStore()
const user = ref({})
user.value = await userStroe.userInfo()

const logout = () => {
  userStroe.logOut()
}
</script>

<style lang="scss" scoped>
@import '@/assets/globalSet.scss';

.in-header {
  box-sizing: border-box;
  margin-bottom: var(--global-interval);
  height: var(--global-header-height);
  background-color: var(--global-content-bg);
  padding: var(--global-padding);
  display: flex;
  justify-content: space-between;
}

.el-breadcrumb {
  line-height: calc(var(--global-header-height) - var(--global-padding) * 2);
}

.username {
  line-height: calc(var(--global-header-height) - var(--global-padding) * 2);
  margin-right: var(--global-padding);
  font-size: 16px;
  color: #333;
}

.el-avatar {
  --el-avatar-size: calc(var(--global-header-height) - var(--global-padding) * 2);
}
</style>
