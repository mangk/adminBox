<template>
    <div>
        <h1>Hello World</h1>
        <h2>{{ data }}</h2>
    </div>
</template>
<script setup>
import { ref, onMounted, getCurrentInstance } from 'vue'
const { proxy } = getCurrentInstance()
const data = ref("")

onMounted(() => {
    proxy.$http.post("api/example").then(res => {
        data.value = res.data.data
    })
})

</script>