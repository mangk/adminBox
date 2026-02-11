<template>
  <div class="serverComponent">
    <ServerComponent />
  </div>
</template>

<script setup>
import { getCurrentInstance } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const { proxy } = getCurrentInstance()
const ServerComponent = proxy.$loadTMPL(route.meta.sc_path, route.path)
</script>
<style scoped>
.serverComponent {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}
</style>
