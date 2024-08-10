<template>
    <div>
      <h1>TODOリスト</h1>
      <ul>
        <li v-for="todo in todos" :key="todo.id">{{ todo.text }}</li>
      </ul>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  import { useFetch } from '#app'
  
  const todos = ref([])
  const { data, error } = await useFetch('http://backend:8080/api/todos')
  
  if (error.value) {
    console.error('Error fetching todos:', error.value)
  } else {
    todos.value = data.value
  }
  </script>
  
  <style scoped>
  /* ページ固有のスタイルをここに追加できます */
  </style>
  