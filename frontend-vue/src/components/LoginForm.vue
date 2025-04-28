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
import { login } from '../api/user';

const emit = defineEmits(['login-success']);
const username = ref('');
const password = ref('');
const msg = ref('');

function loginHandler() {
  if (!username.value || !password.value) {
    msg.value = '请输入用户名和密码';
    return;
  }
  login(username.value, password.value).then(data => {
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
  gap: 10px;
  align-items: stretch;
}
.auth-msg {
  color: #d4380d;
  min-height: 20px;
}
</style>
