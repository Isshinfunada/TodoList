<template>
    <div class="todo-main">
      <h1>TODOリスト</h1>
      <div class="input-group">
        <input v-model="newTask" placeholder="新しいタスクを入力" @keyup.enter="addTodo" />
        <button @click="addTodo">追加</button>
      </div>
      <div v-for="todo in todos" :key="todo.id" class="todo-item">
        <span>{{ todo.text }}</span>
        <div class="buttons">
          <button @click="updateStatus(todo.id)">✔️</button>
          <button @click="deleteTodo(todo.id)">🗑️</button>
        </div>
      </div>
      <div v-if="statusMessage" class="status-message">{{ statusMessage }}</div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        newTask: '',
        todos: [],
        statusMessage: '',
      };
    },
    async created() {
      await this.fetchTodos();
    },
    methods: {
      async fetchTodos() {
        try {
          const response = await fetch(`http://localhost:8080/todos/1`); // ユーザーIDを適切に設定
          if (!response.ok) throw new Error('タスクの取得に失敗しました');
          this.todos = await response.json();
        } catch (error) {
          console.error(error);
          this.statusMessage = 'タスクの取得に失敗しました';
        }
      },
      addTodo() {
        if (this.newTask.trim() === '') return;
        const newTodo = {
          id: Date.now(),
          text: this.newTask,
          status: 'pending',
        };
        this.todos.push(newTodo);
        this.newTask = '';
        this.statusMessage = 'タスクが追加されました';
      },
      updateStatus(id) {
        const todo = this.todos.find(todo => todo.id === id);
        todo.status = 'completed';
        this.statusMessage = 'タスクのステータスが変更されました';
      },
      deleteTodo(id) {
        this.todos = this.todos.filter(todo => todo.id !== id);
        this.statusMessage = 'タスクが削除されました';
      },
    },
  };
  </script>