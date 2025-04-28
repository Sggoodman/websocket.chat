<template>
  <div>
    <LoginForm v-if="!user" @login-success="onLoginSuccess" />
    <RegisterForm v-if="showRegister" @register-success="onRegisterSuccess" @cancel-register="showRegister=false" />
    <ChatRoom v-if="user" :user="user" @logout="onLogout" />
    <div v-if="!user && !showRegister" style="margin-top:16px;text-align:center;">
      <button @click="showRegister=true">还没有账号？注册</button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import LoginForm from '../components/LoginForm.vue';
import RegisterForm from '../components/RegisterForm.vue';
import ChatRoom from '../components/ChatRoom.vue';

const user = ref(JSON.parse(localStorage.getItem('user')) || null);
const showRegister = ref(false);

function onLoginSuccess(u) {
  user.value = u;
  localStorage.setItem('user', JSON.stringify(u));
}
function onRegisterSuccess() {
  showRegister.value = false;
}
function onLogout() {
  user.value = null;
  localStorage.removeItem('user');
}
</script>
