<template>
  <div class="profile-container">
    <el-card class="profile-card">
      <div class="profile-header">
        <el-upload
          class="avatar-uploader"
          action="#"
          :show-file-list="false"
        >
          <el-avatar :size="100" src="https://api.dicebear.com/7.x/avataaars/svg?seed=AA" />
          <div class="upload-hover">Edit 喵</div>
        </el-upload>
        <h3>{{ user.nickname || 'Goshujin-sama' }}</h3>
        <p>{{ user.email }}</p>
      </div>
      <el-form label-position="top" style="margin-top: 30px;">
        <el-form-item label="Nickname">
          <el-input v-model="user.nickname" placeholder="Change nickname 喵~" />
        </el-form-item>
        <el-form-item label="Status Message">
          <el-input v-model="status" placeholder="How are you feeling? 喵~" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" style="width: 100%;">Save Changes 喵!</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const user = ref({
  email: 'user@example.com',
  nickname: 'Goshujin-sama'
})
const status = ref('Coding with AA 喵~')

onMounted(() => {
  const savedUser = localStorage.getItem('user')
  if (savedUser) {
    user.value = JSON.parse(savedUser)
  }
})
</script>

<style scoped>
.profile-container {
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #f5f7fa;
}
.profile-card {
  width: 500px;
  padding: 20px;
}
.profile-header {
  text-align: center;
}
.avatar-uploader {
  position: relative;
  display: inline-block;
  cursor: pointer;
}
.upload-hover {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  border-radius: 50%;
  background: rgba(0,0,0,0.5);
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: opacity 0.3s;
  font-size: 14px;
}
.avatar-uploader:hover .upload-hover {
  opacity: 1;
}
.profile-header h3 {
  margin: 15px 0 5px;
}
.profile-header p {
  color: #999;
  font-size: 14px;
}
</style>
