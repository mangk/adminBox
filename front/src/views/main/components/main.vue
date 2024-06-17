<template>
  <div class="app-content">
    <router-view v-slot="{ Component }">
      <keep-alive>
        <component :is="Component" :key="$route.path" />
      </keep-alive>
    </router-view>
  </div>
</template>

<script setup></script>

<style lang="scss">
@import '@/assets/globalSet.scss';

.app-content {
  background: var(--global-box-bg);
  height: calc(
    100% - var(--global-header-height) - var(--global-interval) * 2 - var(--global-tab-height)
  );
  overflow-y: scroll;
}
</style>
