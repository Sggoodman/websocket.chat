<template>
  <div class="auth-container">
    <LoginForm @login-success="onLoginSuccess" />
    <div class="auth-links">
      <router-link to="/register">还没有账号？注册</router-link>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router';
import LoginForm from '../components/LoginForm.vue';

const router = useRouter();

function onLoginSuccess(user) {
  localStorage.setItem('user', JSON.stringify(user));
  router.push('/chat');
}
</script>

<style scoped>
.auth-container {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 400px;
}

.auth-links {
  margin-top: 1rem;
  text-align: center;
}

.auth-links a {
  color: #1890ff;
  text-decoration: none;
}

.auth-links a:hover {
  text-decoration: underline;
}
</style>
