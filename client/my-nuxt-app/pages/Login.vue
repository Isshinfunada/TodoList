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
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        email: '',
        password: ''
      };
    },
    methods: {
      async handleLogin() {
        try {
          const response = await fetch('http://localhost:8080/login', {
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
          // トークンを保存する処理を追加することができます

          // /home ページにリダイレクト
          this.$router.push('/home');
        } catch (error) {
          console.error('エラー:', error);
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
  </style>