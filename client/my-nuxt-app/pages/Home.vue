<template>
    <div class="todo-main">
      <h1>TODOリスト</h1>
      <div class="input-group">
        <input v-model="newTask" placeholder="新しいタスクを入力" @keyup.enter="addTodo" />
        <button @click="addTodo">追加</button>
      </div>
      <div v-for="todo in todos" :key="todo.id" class="todo-item">
        <span :class="{'completed-text': todo.status === 'completed'}">{{ todo.text }}</span>
        <div class="buttons">
        <button 
          @click="updateStatus(todo.id)"
          :class="{'completed': todo.status === 'completed'}">
          ✔️
        </button>
        <button @click="deleteTodo(todo.id)" class="delete-button">🗑️</button>
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
          // ユーザーID 1 のタスクを取得
          const response = await fetch(`http://localhost:8080/todos/1`); // ユーザーIDを適切に設定
          if (!response.ok) throw new Error('タスクの取得に失敗しました');
          this.todos = await response.json();
          console.log('取得したタスク:', this.todos);  // デバッグ用のログ
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
        // todo.status = 'completed';
        todo.status = todo.status === 'completed' ? 'pending' : 'completed';
        this.statusMessage = 'タスクのステータスが変更されました';
      },
      deleteTodo(id) {
        this.todos = this.todos.filter(todo => todo.id !== id);
        this.statusMessage = 'タスクが削除されました';
      },
    },
  };
  </script>
    <style scoped>
    .todo-main {
      max-width: 400px;
      margin: 20px auto;
      padding: 20px;
      border-radius: 8px;
      box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
      background-color: #fff;
    }

    .input-group {
      display: flex;
      margin-bottom: 20px;
    }
    input {
      flex: 1;
      padding: 10px;
      border: 1px solid #ddd;
      border-radius: 4px;
      margin-right: 10px;
    }
    button {
      padding: 10px;
      border: none;
      background-color: #333;
      color: #fff;
      border-radius: 4px;
      cursor: pointer;
    }
    .todo-item {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 10px;
      padding: 10px;
      border: 1px solid #ddd;
      border-radius: 4px;
    }
    .buttons button {
      background-color: #f0f0f0; /* 明るめの背景色に設定 */
      color: #333;
      margin-left: 5px;
      border-radius: 4px;
      padding: 5px 10px;
      transition: background-color 0.3s, color 0.3s;
    }

    .buttons button.completed {
      background-color: #333; 
      color: #fff;
    }

    .buttons button.completed::before {
      color: #fff;
    }

    .buttons button.delete-button {
      color: white;
    }

    .status-message {
      margin-top: 20px;
      padding: 10px;
      background-color: #f0f0f0;
      border-radius: 4px;
    }

    .completed-text {
  text-decoration: line-through;
}
    </style>