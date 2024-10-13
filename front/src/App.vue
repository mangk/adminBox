<script setup>
import zhCn from 'element-plus/es/locale/lang/zh-cn.mjs'
import en from 'element-plus/es/locale/lang/en.mjs'
let locale = ""
switch (window.adminBox.Locale) {
  case "en":
    locale = en
    break;

  default:
    locale = zhCn
    break;
}
</script>
<template>
  <el-config-provider :locale="locale">
    <Suspense>
      <router-view />
    </Suspense>
  </el-config-provider>
</template>

<style scoped></style>
