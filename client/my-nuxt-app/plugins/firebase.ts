import { initializeApp } from 'firebase/app'
import { getAuth } from 'firebase/auth'

const firebaseConfig = {
    apiKey: "AIzaSyB_WFztzGJtE815-rv7NCDlJrIVKXiowjs",
    authDomain: "todolist-e3d53.firebaseapp.com",
    projectId: "todolist-e3d53",
    storageBucket: "todolist-e3d53.appspot.com",
    messagingSenderId: "348511445961",
    appId: "1:348511445961:web:28d92d4debb0e17b8abe18"
  };

const app = initializeApp(firebaseConfig)
const auth = getAuth(app)

export { auth }