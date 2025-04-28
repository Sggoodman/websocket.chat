<template>
  <div class="chat-room">
    <div class="header">
      <h2>聊天室</h2>
      <div class="user-info">
        <span>{{ user.username }}</span>
        <button class="logout-btn" @click="$emit('logout')">退出</button>
      </div>
    </div>
    
    <div class="message-list" ref="messageList">
      <div v-for="msg in messages" :key="msg.ID" class="message-item">
        <div class="message-header">
          <span class="sender">{{ msg.Sender?.username || '未知用户' }}</span>
          <div class="actions" v-if="msg.SenderID === user.id">
            <button v-if="editingId !== msg.ID" @click="startEdit(msg)">编辑</button>
            <template v-else>
              <button @click="saveEdit(msg)">保存</button>
              <button @click="cancelEdit()">取消</button>
            </template>
          </div>
        </div>
        <div class="message-content" :class="{ 'self': msg.SenderID === user.id }">
          <template v-if="editingId === msg.ID">
            <input v-model="editContent" @keyup.enter="saveEdit(msg)" @keyup.esc="cancelEdit" />
          </template>
          <template v-else>
            {{ msg.Content }}
          </template>
        </div>
      </div>
    </div>

    <div class="input-area">
      <input 
        v-model="newMessage" 
        @keyup.enter="sendMessage"
        placeholder="输入消息..."
      />
      <button @click="sendMessage">发送</button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, nextTick } from 'vue';
import { WS_BASE_URL } from '../api/config';

const props = defineProps({
  user: {
    type: Object,
    required: true
  }
});

const emit = defineEmits(['logout']);
const ws = ref(null);
const messages = ref([]);
const newMessage = ref('');
const editingId = ref(null);
const editContent = ref('');
const messageList = ref(null);

// WebSocket 连接管理
function connectWebSocket() {
  ws.value = new WebSocket(WS_BASE_URL);
  
  ws.value.onmessage = (event) => {
    const msg = JSON.parse(event.data);
    const idx = messages.value.findIndex(m => m.ID === msg.ID);
    if (idx !== -1) {
      messages.value[idx] = msg;
    } else {
      messages.value.push(msg);
    }
    scrollToBottom();
  };

  ws.value.onclose = () => {
    setTimeout(connectWebSocket, 1000);
  };
}

// 发送消息
function sendMessage() {
  if (!newMessage.value.trim() || !ws.value) return;
  
  ws.value.send(JSON.stringify({
    Content: newMessage.value,
    SenderID: props.user.id
  }));
  
  newMessage.value = '';
}

// 编辑消息
function startEdit(msg) {
  editingId.value = msg.ID;
  editContent.value = msg.Content;
}

function cancelEdit() {
  editingId.value = null;
  editContent.value = '';
}

function saveEdit(msg) {
  if (!editContent.value.trim()) return;
  
  fetch('/message', {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
      ID: msg.ID,
      Content: editContent.value,
      SenderID: props.user.id
    })
  }).then(res => {
    if (!res.ok) alert('消息更新失败');
    cancelEdit();
  });
}

// 滚动到底部
function scrollToBottom() {
  nextTick(() => {
    if (messageList.value) {
      messageList.value.scrollTop = messageList.value.scrollHeight;
    }
  });
}

onMounted(() => {
  connectWebSocket();
});

onUnmounted(() => {
  if (ws.value) {
    ws.value.close();
  }
});
</script>

<style scoped>
.chat-room {
  display: flex;
  flex-direction: column;
  height: 80vh;
}
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-bottom: 16px;
  border-bottom: 1px solid #f0f0f0;
}
.user-info {
  display: flex;
  align-items: center;
  gap: 12px;
}
.message-list {
  flex: 1;
  overflow-y: auto;
  padding: 16px 0;
}
.message-item {
  margin-bottom: 16px;
}
.message-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 4px;
}
.sender {
  font-size: 14px;
  color: #666;
}
.message-content {
  display: inline-block;
  padding: 8px 12px;
  background: #f0f0f0;
  border-radius: 4px;
  max-width: 80%;
}
.message-content.self {
  background: #e6f4ff;
}
.input-area {
  display: flex;
  gap: 8px;
  padding-top: 16px;
  border-top: 1px solid #f0f0f0;
}
.input-area input {
  flex: 1;
  padding: 8px;
  border: 1px solid #d9d9d9;
  border-radius: 4px;
}
button {
  padding: 8px 16px;
  border: none;
  border-radius: 4px;
  background: #1677ff;
  color: white;
  cursor: pointer;
}
.actions button {
  padding: 2px 8px;
  font-size: 12px;
  background: transparent;
  color: #1677ff;
}
.logout-btn {
  background: #ff4d4f;
}
</style>
