<template>
  <div class="in-header">Header</div>
  <Tab v-if="!$route.meta.default_menu" />
</template>

<script setup>
import Tab from '@/views/main/components/tab.vue'
</script>

<style lang="scss" scoped>
@import '@/assets/globalSet.scss';

.in-header {
  box-sizing: border-box;
  margin-bottom: var(--global-interval);
  height: var(--global-header-height);
  background-color: var(--global-content-bg);
}
</style>
