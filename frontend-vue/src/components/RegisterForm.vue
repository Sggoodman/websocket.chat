<template>
  <div class="auth-form">
    <h2>注册</h2>
    <input v-model="username" placeholder="用户名" />
    <input v-model="password" type="password" placeholder="密码" />
    <div class="btn-group">
      <button @click="register">注册</button>
      <button class="cancel-btn" @click="$emit('cancel-register')">返回登录</button>
    </div>
    <div class="auth-msg">{{ msg }}</div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { register } from '../api/user';

const emit = defineEmits(['register-success', 'cancel-register']);
const username = ref('');
const password = ref('');
const msg = ref('');

async function registerHandler() {
  if (!username.value || !password.value) {
    msg.value = '请输入用户名和密码';
    return;
  }
  const data = await register(username.value, password.value);
  if (data.id) {
    msg.value = '注册成功，请登录';
    emit('register-success');
  } else {
    msg.value = data.error || '注册失败';
  }
}
</script>

<style scoped>
.auth-form {
  display: flex;
  flex-direction: column;
  gap: 10px;
  align-items: stretch;
}
.auth-msg {
  color: #d4380d;
  min-height: 20px;
  text-align: center;
}
.btn-group {
  display: flex;
  gap: 10px;
}
.cancel-btn {
  background: #f0f0f0;
  color: #666;
}
button {
  flex: 1;
  padding: 8px;
  border: none;
  border-radius: 4px;
  background: #1677ff;
  color: white;
  cursor: pointer;
  transition: opacity 0.2s;
}
button:hover {
  opacity: 0.8;
}
input {
  padding: 8px;
  border: 1px solid #d9d9d9;
  border-radius: 4px;
  font-size: 14px;
}
input:focus {
  border-color: #1677ff;
  outline: none;
}
</style>
