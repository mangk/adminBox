<script setup>
import zhCn from 'element-plus/es/locale/lang/zh-cn.mjs'
</script>
<template>
  <el-config-provider :locale="zhCn">
    <Suspense>
      <router-view />
    </Suspense>
  </el-config-provider>
</template>

<style scoped></style>
