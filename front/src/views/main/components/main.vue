<template>
  <router-view v-slot="{ Component }">
    <template v-if="$route.meta.keep_alive">
      <keep-alive>
        <component :is="Component" :key="$route.path" />
      </keep-alive>
    </template>
    <template v-else>
      <component :is="Component" :key="$route.path" />
    </template>
  </router-view>
</template>

<script setup></script>

<style lang="scss">
.app-content {
  background-color: #fff;
  margin: 10px;
}
</style>
