<template>
  <div class="login-container">
    <el-card class="login-card">
      <template #header>
        <h2 style="text-align: center;">IM Chat 喵~</h2>
      </template>
      <el-form :model="form" @submit.prevent="handleLogin">
        <el-form-item label="Email">
          <el-input v-model="form.email" placeholder="Enter email" />
        </el-form-item>
        <el-form-item label="Password">
          <el-input v-model="form.password" type="password" placeholder="Enter password" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" native-type="submit" style="width: 100%;">Login</el-button>
        </el-form-item>
        <div style="text-align: center; font-size: 12px; color: #999;">
          Don't have an account? <el-link type="primary" @click="isRegister = true">Register 喵!</el-link>
        </div>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import axios from 'axios'

const router = useRouter()
const isRegister = ref(false)
const form = reactive({
  email: '',
  password: ''
})

const handleLogin = async () => {
  try {
    const res = await axios.post('http://localhost:8080/api/login', form)
    localStorage.setItem('token', res.data.token)
    localStorage.setItem('user', JSON.stringify(res.data.user))
    ElMessage.success('Welcome back, Goshujin-sama! 喵~')
    router.push('/')
  } catch (err) {
    ElMessage.error(err.response?.data?.error || 'Login failed')
  }
}
</script>

<style scoped>
.login-container {
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #f5f7fa;
}
.login-card {
  width: 400px;
}
</style>
