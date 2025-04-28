<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script setup>
import { RouterView } from 'vue-router';
</script>

<style>
#app {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #f5f5f5;
}
body {
  background: #f6f8fa;
  font-family: Arial, Helvetica, sans-serif;
}
.container {
  max-width: 500px;
  margin: 40px auto;
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 8px #0001;
  padding: 24px;
}
</style>
