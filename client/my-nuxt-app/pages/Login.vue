<template>
    <div class="login-container">
      <h2>ログイン</h2>
      <form @submit.prevent="handleLogin">
        <div class="form-group">
          <label for="email">Email</label>
          <input type="email" v-model="email" required />
        </div>
        <div class="form-group">
          <label for="password">パスワード</label>
          <input type="password" v-model="password" required />
        </div>
        <button type="submit">ログイン</button>
      </form>
      <!-- エラーメッセージの表示 -->
      <div v-if="statusMessage" class="status-message">{{ statusMessage }}</div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        email: '',
        password: '',
        statusMessage: ''
      };
    },
    methods: {
      async handleLogin() {
        try {
          const apiUrl = import.meta.env.VITE_API_URL;
          const response = await fetch(`${apiUrl}/todos/login`, {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json'
            },
            body: JSON.stringify({
              email: this.email,
              password: this.password
            })
          });

          if (!response.ok) {
            throw new Error('ログインに失敗しました');
          }

          const data = await response.json();
          console.log('ログイン成功:', data.token);
        localStorage.setItem('jwtToken', data.token)

          this.$router.push('/home');
        } catch (error) {
          console.error('エラー:', error);
          this.statusMessage = 'ログインに失敗しました。メールアドレスまたはパスワードが間違っています。';
        }
      }
    }
  };
  </script>
  
  <style scoped>
  .login-container {
    max-width: 400px;
    margin: 0 auto;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 5px;
  }
  
  .form-group {
    margin-bottom: 15px;
  }
  
  label {
    display: block;
    margin-bottom: 5px;
  }
  
  input {
    width: 100%;
    padding: 8px;
    box-sizing: border-box;
  }
  
  button {
    width: 100%;
    padding: 10px;
    background-color: #000;
    color: #fff;
    border: none;
    border-radius: 5px;
    cursor: pointer;
  }
  
  button:hover {
    background-color: #333;
  }

  /* エラーメッセージ用のスタイル */
.status-message {
  margin-top: 15px;
  color: red; /* エラーメッセージを赤色で表示 */
}
  </style>