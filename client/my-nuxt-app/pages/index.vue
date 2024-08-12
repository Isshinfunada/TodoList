<<template>
  <div class="container">
    <h1>TODOリスト</h1>
    <div class="button-group">
      <nuxt-link to="/register" class="register-button">会員登録</nuxt-link>
      <nuxt-link to="/login" class="login-button">ログイン</nuxt-link>
    </div>
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

.button-group {
  display: flex;
  gap: 10px;
}

.register-button, .login-button {
  padding: 10px 20px;
  background-color: #000;
  color: #fff;
  text-decoration: none;
  border-radius: 5px;
  transition: background-color 0.3s;
}

.register-button:hover, .login-button:hover {
  background-color: #333;
}
</style>