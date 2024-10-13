<template>
  
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
