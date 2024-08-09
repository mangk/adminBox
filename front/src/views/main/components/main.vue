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
</style>
