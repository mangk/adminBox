<template>
  <template v-for="menu in menus">
    <el-sub-menu
      v-if="menu.children && menu.children.length && !menu.hidden"
      :index="menu.name"
      :key="menu.id"
    >
      <template #title>
        <el-icon v-if="menu.meta.icon">
          <component :is="menu.meta.icon" />
        </el-icon>
        <span>{{ menu.meta.title }}</span>
      </template>
      <MenuTree :menus="menu.children"></MenuTree>
    </el-sub-menu>
    <el-menu-item v-else-if="!menu.hidden" :index="menu.name" :route="menu">
      <el-icon v-if="menu.meta.icon">
        <component :is="menu.meta.icon" />
      </el-icon>
      <span>{{ menu.meta.title }}</span>
    </el-menu-item>
  </template>
</template>

<script setup>
defineProps(['menus'])
</script>
<style lang="scss" scoped></style>
