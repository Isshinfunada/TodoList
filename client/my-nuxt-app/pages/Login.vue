<template>
  <div class="login-container">
    <h2>ログイン</h2>
    <client-only>
      <div id="firebaseui-auth-container"></div>
    </client-only>
    <!-- エラーメッセージの表示 -->
    <div v-if="statusMessage" class="status-message">{{ statusMessage }}</div>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { useNuxtApp } from '#app'
import { EmailAuthProvider, GoogleAuthProvider } from 'firebase/auth'

const { $firebaseAuth } = useNuxtApp()
const statusMessage = ref('')

onMounted(async () => {
  try {
    // 動的インポートで firebaseui を読み込む
    const firebaseuiModule = await import('firebaseui')
    const firebaseui = firebaseuiModule.default || firebaseuiModule

    // firebaseui の CSS をインポート
    await import('firebaseui/dist/firebaseui.css')

    const uiConfig = {
      signInSuccessUrl: '/home',
      signInOptions: [
        EmailAuthProvider.PROVIDER_ID,
        GoogleAuthProvider.PROVIDER_ID,
      ],
      callbacks: {
        signInSuccessWithAuthResult: function(authResult, redirectUrl) {
          // ユーザーが正常にサインインした場合の処理
          return true // 自動的にリダイレクトを続行
        },
        uiShown: function() {
          // FirebaseUI ウィジェットが表示された際の処理
          const loader = document.getElementById('loader')
          if (loader) {
            loader.style.display = 'none'
          }
        }
      },
      signInFlow: 'popup'
    }

    // 既にAuthUIインスタンスが存在する場合は再利用する
    let ui = firebaseui.auth.AuthUI.getInstance() || new firebaseui.auth.AuthUI($firebaseAuth)
    ui.start('#firebaseui-auth-container', uiConfig)
  } catch (error) {
    console.error('FirebaseUIの初期化に失敗しました:', error)
    statusMessage.value = 'ログインに失敗しました。再度お試しください。'
  }
})
</script>

<style scoped>
.login-container {
  max-width: 400px;
  margin: 0 auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

h2 {
  text-align: center;
  margin-bottom: 20px;
}

/* エラーメッセージ用のスタイル */
.status-message {
  margin-top: 15px;
  color: red;
}
</style>
