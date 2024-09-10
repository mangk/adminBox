<template>
  <div class="app-content">
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
  </div>
</template>

<script setup></script>

<style lang="scss">
@import '@/assets/globalSet.scss';
</style>
