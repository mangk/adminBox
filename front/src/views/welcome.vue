<template>
  1234
  {{ user }}
  <el-button @click="logout">退出</el-button>
</template>

<script setup>
import { useUserStore } from '@/pinia/useUserStore'
import { ref } from 'vue'

const userStroe = useUserStore()
const user = ref({})
user.value = await userStroe.userInfo()

const logout = () => {
  userStroe.logOut()
}
</script>
