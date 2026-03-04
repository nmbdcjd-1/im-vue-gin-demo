<template>
  <div class="chat-container">
    <div class="chat-list">
      <div class="search-bar">
        <el-input placeholder="Search 喵~" :prefix-icon="Search" />
      </div>
      <el-scrollbar>
        <div v-for="i in 10" :key="i" class="chat-item" :class="{ active: i === 1 }">
          <el-avatar :size="40" :src="`https://api.dicebear.com/7.x/avataaars/svg?seed=${i}`" />
          <div class="chat-info">
            <div class="chat-header">
              <span class="name">Goshujin-sama {{ i }}</span>
              <span class="time">14:00</span>
            </div>
            <div class="last-msg">Hello AA! 喵~</div>
          </div>
        </div>
      </el-scrollbar>
    </div>
    <div class="chat-window">
      <div class="window-header">
        <span>Goshujin-sama 1</span>
      </div>
      <el-scrollbar ref="msgScroll" class="message-list">
        <div v-for="(msg, idx) in messages" :key="idx" :class="['message', msg.self ? 'self' : '']">
          <el-avatar :size="32" :src="msg.avatar" />
          <div class="message-bubble">{{ msg.content }}</div>
        </div>
      </el-scrollbar>
      <div class="message-input">
        <div class="toolbar">
          <el-icon><Smile /></el-icon>
          <el-icon><Image /></el-icon>
          <el-icon><Mic /></el-icon>
        </div>
        <el-input
          v-model="inputMsg"
          type="textarea"
          :rows="3"
          placeholder="Type a message... 喵~"
          @keyup.enter="sendMsg"
        />
        <div class="send-btn">
          <el-button type="primary" @click="sendMsg">Send 喵!</el-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { Search, Smile, Image, Mic } from 'lucide-vue-next'

const inputMsg = ref('')
const messages = reactive([
  { content: 'Welcome back! 喵~', self: false, avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=AA' }
])

// WebSocket 连接
const socket = new WebSocket('ws://' + window.location.hostname + ':8080/ws')

socket.onmessage = (event) => {
  const msg = JSON.parse(event.data)
  messages.push({
    content: msg.content,
    self: false,
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=AA'
  })
}

const sendMsg = () => {
  if (!inputMsg.value.trim()) return
  const msgPayload = {
    content: inputMsg.value,
    from_user_id: 1, // 演示用
  }
  socket.send(JSON.stringify(msgPayload))
  
  messages.push({
    content: inputMsg.value,
    self: true,
    avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=1'
  })
  inputMsg.value = ''
}
</script>

<style scoped>
.chat-container {
  display: flex;
  height: 100%;
}
.chat-list {
  width: 300px;
  border-right: 1px solid #eee;
  display: flex;
  flex-direction: column;
}
.search-bar {
  padding: 15px;
}
.chat-item {
  display: flex;
  padding: 15px;
  gap: 12px;
  cursor: pointer;
  transition: background 0.2s;
}
.chat-item:hover {
  background: #f5f7fa;
}
.chat-item.active {
  background: #ecf5ff;
}
.chat-info {
  flex: 1;
  overflow: hidden;
}
.chat-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 4px;
}
.name {
  font-weight: 500;
  font-size: 14px;
}
.time {
  font-size: 12px;
  color: #999;
}
.last-msg {
  font-size: 12px;
  color: #666;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
.chat-window {
  flex: 1;
  display: flex;
  flex-direction: column;
}
.window-header {
  padding: 15px 20px;
  border-bottom: 1px solid #eee;
  font-weight: 500;
}
.message-list {
  flex: 1;
  padding: 20px;
}
.message {
  display: flex;
  gap: 10px;
  margin-bottom: 20px;
}
.message.self {
  flex-direction: row-reverse;
}
.message-bubble {
  max-width: 60%;
  padding: 10px 15px;
  border-radius: 8px;
  background: #f4f4f5;
  font-size: 14px;
  line-height: 1.5;
}
.message.self .message-bubble {
  background: #409eff;
  color: #fff;
}
.message-input {
  padding: 20px;
  border-top: 1px solid #eee;
}
.toolbar {
  display: flex;
  gap: 15px;
  margin-bottom: 10px;
  color: #666;
}
.toolbar i {
  cursor: pointer;
}
.send-btn {
  display: flex;
  justify-content: flex-end;
  margin-top: 10px;
}
</style>
