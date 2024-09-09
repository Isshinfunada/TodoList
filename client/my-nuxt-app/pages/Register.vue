<template>
    <div class="register-container">
      <h1>会員登録</h1>
      <form @submit.prevent="register">
        <div class="form-group">
          <label for="username">ユーザーネーム</label>
          <input type="text" id="username" v-model="username" required />
        </div>
        <div class="form-group">
          <label for="email">Email</label>
          <input type="email" id="email" v-model="email" required />
        </div>
        <div class="form-group">
          <label for="password">パスワード</label>
          <input type="password" id="password" v-model="password" required />
        </div>
        <button type="submit" class="register-button">登録</button>
      </form>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  import { useRouter } from 'vue-router'

  const router = useRouter()

  
  const username = ref('')
  const email = ref('')
  const password = ref('')
  
  const register = async () => {
    try {
      const response = await fetch('http://localhost:8080/users', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          Username: username.value,
          Email: email.value,
          Password: password.value
        })
      })

      if (!response.ok) {
        throw new Error('登録に失敗しました')
      }

      const user = await response.json()
      console.log('登録成功:', user)
        const apiUrl = import.meta.env.VITE_API_URL;
        const loginResponse = await fetch(`${apiUrl}/todos/login`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          email: email.value,
          password: password.value
        })
      })

      if (!loginResponse.ok) {
        throw new Error('自動ログインに失敗しました')
      }

      const loginData = await loginResponse.json()
      console.log('自動ログイン成功:', loginData.token)
      
      localStorage.setItem('jwtToken', loginData.token)
      router.push('/home'); 

    } catch (error) {
      console.error('エラー:', error)
    }
  }
</script>

<style scoped>
.register-container {
  max-width: 400px;
  margin: 0 auto;
  padding: 20px;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

h1 {
  text-align: center;
  margin-bottom: 20px;
}

.form-group {
  margin-bottom: 15px;
}

label {
  display: block;
  margin-bottom: 5px;
  font-weight: bold;
}

input {
  width: 100%;
  padding: 8px;
  box-sizing: border-box;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.register-button {
  width: 100%;
  padding: 10px;
  background-color: #000;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.register-button:hover {
  background-color: #333;
}
</style>