<template>
  <div class="chat-container">
    <ChatRoom :user="user" @logout="onLogout" />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import ChatRoom from '../components/ChatRoom.vue';

const router = useRouter();
const user = ref(null);

onMounted(() => {
  const userStr = localStorage.getItem('user');
  if (!userStr) {
    router.push('/login');
    return;
  }
  user.value = JSON.parse(userStr);
});

function onLogout() {
  localStorage.removeItem('user');
  router.push('/login');
}
</script>

<style scoped>
.chat-container {
  width: 100%;
  max-width: 800px;
  margin: 0 auto;
  padding: 1rem;
}
</style>
