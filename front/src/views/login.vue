<template>
  <div class="login-box">
    <el-image src="/images/login_bg.png" />
    <div class="form-box">
      <el-form
        ref="formRef"
        :model="form"
        :rules="rules"
        status-icon
        label-position="top"
        :show-all-levels="false"
        style="width: 100%"
        hide-required-asterisk
      >
        <h2 style="color: #666; margin-bottom: 30px">登录</h2>
        <el-form-item label="用户名" prop="username">
          <el-input size="large" v-model="form.username" />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input size="large" type="password" v-model="form.password" />
        </el-form-item>
        <el-form-item label="验证码" class="v-box" prop="verification_code">
          <el-input size="large" v-model="form.verification_code" style="width: 55%" />
          <el-image
            style="width: 40%; height: 38px"
            :src="verificationPath"
            @click="changeVerification"
          />
        </el-form-item>
        <el-form-item style="margin-top: 30px">
          <el-button size="large" type="primary" style="width: 100%" @click="onSubmit(formRef)"
            >登录
          </el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>
<script setup>
import { IsRewriteIndex, verificationCode } from '@/api/auth'
import { useUserStore } from '@/pinia/useUserStore'
import { getCurrentInstance, reactive, ref, onMounted, onBeforeMount } from 'vue'
import { useRouter } from 'vue-router'

const { proxy } = getCurrentInstance()
const formRef = ref()
const form = reactive({
  username: '',
  password: '',
  verification_code: '',
  captcha_id: ''
})

const verificationPath = ref('')

const changeVerification = () => {
  verificationCode().then((res) => {
    verificationPath.value = res.data.pic_path
    form.captcha_id = res.data.captcha_id
  })
}
changeVerification()

const rules = reactive({
  username: [{ required: true, message: '填写用户名', trigger: 'blur' }],
  password: [{ required: true, message: '填写密码', trigger: 'blur' }],
  verification_code: [{ required: true, message: '填写验证码', trigger: 'blur' }]
})

const userStroe = useUserStore()
const router = useRouter()

const onSubmit = async (formEl) => {
  if (!formEl) return
  await formEl.validate(async (valid, fields) => {
    if (!valid) {
      proxy.$message.error('请补全信息')
    } else {
      userStroe.logIn(form).then((toRouteName) => {
        if (!toRouteName) {
          changeVerification()
        } else {
          router.push(toRouteName)
        }
      })
    }
  })
}

onBeforeMount(() => {
  IsRewriteIndex().then((res) => {
    if (res.data) {
      window.location.href = window.location.origin
    }
  })
})

onMounted(() => {
  if (userStroe.isLogIn()) {
    router.push({ path: '/' })
  }
})
</script>
<style lang="scss" scoped>
.el-form-item {
  margin-bottom: 15px;
}

.login-box {
  width: 100%;
  height: 100%;
  display: flex;
  flex-flow: row nowrap;
}

.form-box {
  width: 450px;
  min-width: 450px;
  box-sizing: border-box;
  padding: calc(var(--global-padding) * 6);
  display: flex;
  align-items: center;
  /* 垂直居中 */
  justify-content: center;
  /* 水平居中（可选） */
}

:deep(.v-box > .el-form-item__content) {
  display: flex;
  flex-flow: row nowrap;
  justify-content: space-between;
}

:deep(.el-form-item__label) {
  content: '';
}
</style>
