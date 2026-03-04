<template>
  <el-container class="layout-container">
    <el-aside width="64px" class="side-nav">
      <div class="nav-icons">
        <el-avatar :size="40" src="https://api.dicebear.com/7.x/avataaars/svg?seed=AA" class="user-avatar" />
        <el-divider />
        <router-link to="/">
          <el-icon :size="24" :class="{ active: $route.path === '/' }"><MessageSquare /></el-icon>
        </router-link>
        <router-link to="/contacts">
          <el-icon :size="24" :class="{ active: $route.path === '/contacts' }"><Users /></el-icon>
        </router-link>
        <router-link to="/profile">
          <el-icon :size="24" :class="{ active: $route.path === '/profile' }"><User /></el-icon>
        </router-link>
      </div>
      <div class="bottom-icons">
        <el-icon :size="24" @click="logout"><LogOut /></el-icon>
      </div>
    </el-aside>
    <el-main class="main-content">
      <router-view />
    </el-main>
  </el-container>
</template>

<script setup>
import { MessageSquare, Users, User, LogOut } from 'lucide-vue-next'
import { useRouter } from 'vue-router'

const router = useRouter()

const logout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  router.push('/login')
}
</script>

<style scoped>
.layout-container {
  height: 100vh;
}
.side-nav {
  background: #2c3e50;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  padding: 20px 0;
}
.nav-icons {
  display: flex;
  flex-direction: column;
  gap: 25px;
  align-items: center;
}
.nav-icons a {
  color: #bdc3c7;
  transition: color 0.3s;
}
.nav-icons a:hover, .nav-icons .active {
  color: #409eff;
}
.bottom-icons {
  color: #bdc3c7;
  cursor: pointer;
}
.user-avatar {
  border: 2px solid #409eff;
}
.main-content {
  padding: 0;
  background: #fff;
}
</style>
