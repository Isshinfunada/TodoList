<template>
  <div class="container">
    <h1>TODOリスト</h1>
    <nuxt-link to="/register" class="register-button">会員登録</nuxt-link>
  </div>
</template>

<script setup>
// No script needed for this simple component
</script>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  background-color: #f9f9f9;
}

h1 {
  font-size: 2rem;
  margin-bottom: 20px;
}

.register-button {
  padding: 10px 20px;
  background-color: #000;
  color: #fff;
  text-decoration: none;
  border-radius: 5px;
  transition: background-color 0.3s;
}

.register-button:hover {
  background-color: #333;
}
</style>