<template>
  <div class="register-container">
    <h1>会員登録</h1>
    <client-only>
      <div id="firebaseui-auth-container"></div>
    </client-only>
  </div>
</template>


<script setup>
import { onMounted } from 'vue'
import { useNuxtApp } from '#app'
import { EmailAuthProvider, GoogleAuthProvider } from 'firebase/auth'

const { $firebaseAuth } = useNuxtApp()

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
      // その他のオプション（必要に応じて追加）
    }

    // 既にAuthUIインスタンスが存在する場合は再利用する
    let ui = firebaseui.auth.AuthUI.getInstance() || new firebaseui.auth.AuthUI($firebaseAuth)
    ui.start('#firebaseui-auth-container', uiConfig)
  } catch (error) {
    console.error('FirebaseUIの初期化に失敗しました:', error)
  }
})
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
</style>
