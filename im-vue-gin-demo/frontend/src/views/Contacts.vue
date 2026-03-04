<template>
  <div class="contacts-container">
    <div class="contacts-list">
      <div class="contact-menu">
        <el-button type="primary" style="width: 100%;" @click="addFriendDialog = true">Add New Friend 喵~</el-button>
        <el-menu default-active="1" style="border: none; margin-top: 10px;">
          <el-menu-item index="1">New Friends</el-menu-item>
          <el-menu-item index="2">Group Chats</el-menu-item>
        </el-menu>
      </div>
      <el-divider style="margin: 0;" />
      <el-scrollbar>
        <div v-for="letter in letters" :key="letter">
          <div class="letter-group">{{ letter }}</div>
          <div v-for="i in 3" :key="i" class="contact-item">
            <el-avatar :size="36" :src="`https://api.dicebear.com/7.x/avataaars/svg?seed=${letter}${i}`" />
            <span>Friend {{ letter }}{{ i }} 喵</span>
          </div>
        </div>
      </el-scrollbar>
    </div>
    <div class="contact-detail">
      <div class="empty-state">
        <el-empty description="Select a contact to view profile 喵~" />
      </div>
    </div>

    <el-dialog v-model="addFriendDialog" title="Add Friend 喵" width="400px">
      <el-input placeholder="Enter email address" />
      <template #footer>
        <el-button @click="addFriendDialog = false">Cancel</el-button>
        <el-button type="primary" @click="addFriendDialog = false">Search 喵!</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref } from 'vue'
const letters = ['A', 'B', 'C', 'D', 'E']
const addFriendDialog = ref(false)
</script>

<style scoped>
.contacts-container {
  display: flex;
  height: 100%;
}
.contacts-list {
  width: 300px;
  border-right: 1px solid #eee;
  display: flex;
  flex-direction: column;
}
.contact-menu {
  padding: 15px;
}
.letter-group {
  padding: 5px 15px;
  background: #f8f9fa;
  font-size: 12px;
  color: #999;
}
.contact-item {
  display: flex;
  align-items: center;
  padding: 10px 15px;
  gap: 12px;
  cursor: pointer;
}
.contact-item:hover {
  background: #f5f7fa;
}
.contact-detail {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
