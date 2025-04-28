<template>
  <div class="auth-form">
    <h2>登录</h2>
    <input v-model="username" placeholder="用户名" />
    <input v-model="password" type="password" placeholder="密码" />
    <button @click="login">登录</button>
    <div class="auth-msg">{{ msg }}</div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { login as loginApi } from '../api/user';

const emit = defineEmits(['login-success']);
const username = ref('');
const password = ref('');
const msg = ref('');

function loginHandler() {
  if (!username.value || !password.value) {
    msg.value = '请输入用户名和密码';
    return;
  }
  loginApi(username.value, password.value).then(data => {
    if (data.id) {
      emit('login-success', { id: data.id, username: username.value });
    } else {
      msg.value = data.error || '登录失败';
    }
  });
}

const login = loginHandler;
</script>

<style scoped>
.auth-form {
  display: flex;
  flex-direction: column;
  gap: 16px;
  align-items: stretch;
  background: white;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 400px;
}

h2 {
  margin: 0 0 1rem;
  text-align: center;
  color: #1890ff;
}

input {
  padding: 8px 12px;
  border: 1px solid #d9d9d9;
  border-radius: 4px;
  font-size: 14px;
}

input:focus {
  border-color: #1890ff;
  outline: none;
  box-shadow: 0 0 0 2px rgba(24, 144, 255, 0.2);
}

button {
  padding: 8px 16px;
  background: #1890ff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
  transition: background 0.3s;
}

button:hover {
  background: #40a9ff;
}

.auth-msg {
  color: #d4380d;
  min-height: 20px;
  text-align: center;
  font-size: 14px;
}
</style>
