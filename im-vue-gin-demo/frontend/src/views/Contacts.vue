<template>
  <div class="contacts-container">
    <div class="contacts-list">
      <div class="contact-menu">
        <el-button type="primary" style="width: 100%;" @click="addFriendDialog = true">Add New Friend 喵~</el-button>
        <el-menu default-active="1" style="border: none; margin-top: 10px;" @select="handleMenuSelect">
          <el-menu-item index="requests">New Friends 
            <el-badge v-if="requests.length" :value="requests.length" class="req-badge" />
          </el-menu-item>
          <el-menu-item index="groups">Group Chats</el-menu-item>
        </el-menu>
      </div>
      <el-divider style="margin: 0;" />
      <el-scrollbar>
        <div v-for="letter in letters" :key="letter">
          <div class="letter-group">{{ letter }}</div>
          <div v-for="friend in friendList[letter]" :key="friend.id" class="contact-item" @click="selectFriend(friend)">
            <el-avatar :size="36" :src="friend.avatar || `https://api.dicebear.com/7.x/avataaars/svg?seed=${friend.email}`" />
            <span>{{ friend.nickname || friend.email }}</span>
          </div>
        </div>
      </el-scrollbar>
    </div>
    <div class="contact-detail">
      <div v-if="activeTab === 'requests'" class="request-panel">
        <h3>New Friend Requests 喵</h3>
        <el-empty v-if="!requests.length" description="No new requests 喵~" />
        <div v-for="req in requests" :key="req.id" class="req-item">
          <div class="req-info">
            <el-avatar :size="40" :src="`https://api.dicebear.com/7.x/avataaars/svg?seed=${req.from_email}`" />
            <div>
              <div class="req-nick">{{ req.nickname }}</div>
              <div class="req-email">{{ req.from_email }}</div>
            </div>
          </div>
          <div class="req-btns">
            <el-button type="success" size="small" @click="handleRequest(req.id, 1)">Accept</el-button>
            <el-button type="danger" size="small" @click="handleRequest(req.id, 2)">Reject</el-button>
          </div>
        </div>
      </div>
      <div v-else class="empty-state">
        <el-empty description="Select a contact or tab 喵~" />
      </div>
    </div>

    <el-dialog v-model="addFriendDialog" title="Add Friend 喵" width="400px">
      <el-input v-model="searchEmail" placeholder="Enter email address" />
      <template #footer>
        <el-button @click="addFriendDialog = false">Cancel</el-button>
        <el-button type="primary" @click="submitAddFriend">Search & Add 喵!</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue'
import axios from 'axios'
import { ElMessage } from 'element-plus'

const API_BASE = 'http://localhost:8080/api'
const letters = ref([])
const friendList = reactive({})
const requests = ref([])
const activeTab = ref('')
const addFriendDialog = ref(false)
const searchEmail = ref('')

const fetchRequests = async () => {
  try {
    const res = await axios.get(`${API_BASE}/friend/requests`)
    requests.value = res.data || []
  } catch (e) { console.error(e) }
}

const fetchFriends = async () => {
  try {
    const res = await axios.get(`${API_BASE}/friends`)
    const list = res.data || []
    // 简单首字母排序逻辑
    letters.value = ['A', 'F', 'S'] // 演示
    friendList['A'] = list.filter(u => u.email.startsWith('a'))
    friendList['F'] = list.filter(u => u.email.startsWith('f'))
    friendList['S'] = list.filter(u => u.email.startsWith('s'))
  } catch (e) { console.error(e) }
}

const handleMenuSelect = (index) => {
  activeTab.value = index
}

const handleRequest = async (id, action) => {
  try {
    await axios.post(`${API_BASE}/friend/handle`, { id, action })
    ElMessage.success(action === 1 ? 'Accepted! 喵~' : 'Rejected 喵')
    fetchRequests()
    fetchFriends()
  } catch (e) { ElMessage.error('Action failed 喵') }
}

const submitAddFriend = async () => {
  try {
    await axios.post(`${API_BASE}/friend/add`, { email: searchEmail.value })
    ElMessage.success('Request sent 喵!')
    addFriendDialog.value = false
  } catch (e) { ElMessage.error(e.response?.data?.error || 'Failed 喵') }
}

onMounted(() => {
  fetchRequests()
  fetchFriends()
})
</script>

<style scoped>
.contacts-container { display: flex; height: 100%; }
.contacts-list { width: 300px; border-right: 1px solid #eee; display: flex; flex-direction: column; }
.contact-menu { padding: 15px; }
.req-badge { margin-left: 10px; }
.letter-group { padding: 5px 15px; background: #f8f9fa; font-size: 12px; color: #999; }
.contact-item { display: flex; align-items: center; padding: 10px 15px; gap: 12px; cursor: pointer; transition: 0.2s; }
.contact-item:hover { background: #f5f7fa; }
.contact-detail { flex: 1; padding: 40px; }
.request-panel h3 { margin-bottom: 30px; border-bottom: 2px solid #409eff; padding-bottom: 10px; display: inline-block; }
.req-item { display: flex; justify-content: space-between; align-items: center; padding: 15px; border: 1px solid #eee; border-radius: 8px; margin-bottom: 15px; }
.req-info { display: flex; gap: 15px; align-items: center; }
.req-nick { font-weight: 500; font-size: 14px; }
.req-email { font-size: 12px; color: #999; }
.req-btns { display: flex; gap: 10px; }
.empty-state { height: 100%; display: flex; align-items: center; justify-content: center; }
</style>
