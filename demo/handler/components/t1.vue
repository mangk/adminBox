<template>
  <div style="padding: var(--global-padding)" class="t">
    测试
    <el-button @click="count++">Count is: {{ count }}</el-button>
    <el-button @click="open">Count is: {{ count }}</el-button>
    <div style="width: 100px;height: 100px;background-color: #529b2e"></div>
    <div style="" class="text-3xl underline font-black border-amber-700 bg-red-500">
      Hello world!
    </div>

    <BBB></BBB>
  </div>
</template>

<script setup>
import { ref, getCurrentInstance } from "vue";
const count = ref(0)

const { proxy } = getCurrentInstance()

proxy.$loadJS("https://cdn.tailwindcss.com")

const BBB = proxy.$loadTMPL('http://127.0.0.1:8910/b', "fdasdf")

const open = () => {
  console.log(proxy);
  proxy.$message("聪明2");
  console.log(proxy.$axios)
  proxy.$http.get("http://127.0.0.1:8910").then(res => {
    console.log(111, res);
  })
}

</script>

<style scoped>
.t {
  background-color: #ccc;
}
</style>