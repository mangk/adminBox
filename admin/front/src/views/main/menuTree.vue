<template>
  <template v-for="menu in menus">
    <!-- 有可见子项时，显示为可展开的子菜单 -->
    <el-sub-menu v-if="menu.children && menu.children.length && !menu.hidden && hasVisibleChild(menu.children)"
      :index="menu.name" :key="menu.id">
      <template #title>
        <el-icon v-if="menu.meta.icon">
          <component :is="menu.meta.icon" />
        </el-icon>
        <span>{{ menu.meta.title }}</span>
      </template>
      <MenuTree :menus="menu.children"></MenuTree>
    </el-sub-menu>
    <!-- 无可见子项或自己被隐藏时，显示为菜单项 -->
    <el-menu-item v-else-if="!menu.hidden" :index="menu.name" :route="menu" :class="{ active: shouldHighlight(menu) }">
      <el-icon v-if="menu.meta.icon">
        <component :is="menu.meta.icon" />
      </el-icon>
      <span>{{ menu.meta.title }}</span>
    </el-menu-item>
  </template>
</template>

<script setup>
import { useRoute } from 'vue-router'

defineProps(['menus'])

const route = useRoute()

// 检查是否有可见的子项
const hasVisibleChild = (children) => {
  if (!children || children.length === 0) return false
  return children.some(child => {
    if (!child.hidden) return true
    if (child.children && child.children.length) {
      return hasVisibleChild(child.children)
    }
    return false
  })
}

// 检查是否应该高亮该菜单项（包括其隐藏的子项被激活的情况）
const shouldHighlight = (menu) => {
  // 如果菜单本身就匹配当前路由，直接返回 true
  if (menu.name === route.name) {
    return true
  }

  // 如果菜单有隐藏的子项，检查当前路由是否在其子树中
  if (menu.children && menu.children.length && !hasVisibleChild(menu.children)) {
    return isRouteInSubtree(menu.children, route.name)
  }

  return false
}

// 递归检查路由是否在子树中
const isRouteInSubtree = (children, routeName) => {
  if (!children || children.length === 0) return false
  return children.some(child => {
    if (child.name === routeName) return true
    if (child.children && child.children.length) {
      return isRouteInSubtree(child.children, routeName)
    }
    return false
  })
}
</script>

<style lang="scss" scoped>
:deep(.active) {
  color: var(--el-menu-active-color);
  background-color: var(--el-menu-active-bg-color);
}
</style>