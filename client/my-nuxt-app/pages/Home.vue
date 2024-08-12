<template>
    <div class="todo-main">
      <h1>TODOリスト</h1>
      <div class="input-group">
        <input v-model="newTask" placeholder="新しいタスクを入力" @keyup.enter="addTodo" />
        <button @click="addTodo">追加</button>
      </div>
      <div v-if="todos.length > 0">
        <div v-for="todo in todos" :key="todo.ID" class="todo-item">
          <input v-if="todo.isEditing" v-model="todo.Text" @blur="editTodo(todo)" @keyup.enter="editTodo(todo)" class="edit-input" />
          <span v-else @click="enableEdit(todo)" :class="{'completed-text': todo.Status === 'completed'}">{{ todo.Text }}</span>
          <div class="buttons">
            <button 
              @click="updateStatus(todo)"
              :class="{'completed': todo.Status === 'completed'}">
              ✔️
            </button>
            <button @click="deleteTodo(todo.ID)" class="delete-button">🗑️</button>
          </div>
        </div>
      </div>
      <div v-else>
        <p>タスクがありません。</p>
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
  mounted() {
    this.fetchTodos();
  },
  methods: {
    async fetchTodos() {
      const token = localStorage.getItem('jwtToken');
      if (!token) {
        console.error("No token found");
        return;
      }

      try {
        const response = await fetch(`http://localhost:8080/todos/list`, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        if (!response.ok) throw new Error('タスクの取得に失敗しました');
        this.todos = await response.json();
        this.todos.forEach(todo => todo.isEditing = false);
      } catch (error) {
        console.error(error);
        this.statusMessage = 'タスクの取得に失敗しました';
      }
    },
    async addTodo() {
      if (this.newTask.trim() === '') return;

      const token = localStorage.getItem('jwtToken');
      if (!token) {
        console.error("No token found");
        return;
      }

      try {
        const response = await fetch('http://localhost:8080/todos/create', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${token}`
          },
          body: JSON.stringify({
            text: this.newTask,
            status: 'pending'
          })
        });

        if (!response.ok) {
          throw new Error('Failed to create todo');
        }

        const newTodo = await response.json();
        newTodo.isEditing = false;
        this.todos.push(newTodo);
        this.newTask = '';
        this.statusMessage = 'タスクが追加されました';
      } catch (error) {
        console.error('Error creating todo:', error);
        this.statusMessage = 'タスクの作成に失敗しました';
      }
    },
    enableEdit(todo) {
      todo.isEditing = true;
    },
    async editTodo(todo) {
      todo.isEditing = false;
      const token = localStorage.getItem('jwtToken');
      if (!token) {
        console.error("No token found");
        return;
      }

      try {
        const response = await fetch('http://localhost:8080/todos/edit', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${token}`
          },
          body: JSON.stringify({
            id: todo.ID,
            text: todo.Text
          })
        });

        if (!response.ok) {
          throw new Error('Failed to edit todo');
        }

        this.statusMessage = 'タスクが編集されました';
      } catch (error) {
        console.error('Error editing todo:', error);
        this.statusMessage = 'タスクの編集に失敗しました';
      }
    },
    async updateStatus(todo) {
      const token = localStorage.getItem('jwtToken');
      if (!token) {
        console.error("No token found");
        return;
      }

      try {
        const response = await fetch(`http://localhost:8080/todos/${todo.ID}/status`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${token}`
          },
          body: JSON.stringify({
            status: todo.Status === 'completed' ? 'pending' : 'completed'
          })
        });

        if (!response.ok) {
          throw new Error('Failed to update todo status');
        }

        todo.Status = todo.Status === 'completed' ? 'pending' : 'completed';
        this.statusMessage = 'タスクのステータスが変更されました';
      } catch (error) {
        console.error('Error updating todo status:', error);
        this.statusMessage = 'ステータスの更新に失敗しました';
      }
    },
    async deleteTodo(id) {
      const token = localStorage.getItem('jwtToken');
      if (!token) {
        console.error("No token found");
        return;
      }

      try {
        const response = await fetch('http://localhost:8080/todos/delete', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${token}`
          },
          body: JSON.stringify({
            id: id
          })
        });

        if (!response.ok) {
          throw new Error('Failed to delete todo');
        }

        this.todos = this.todos.filter(todo => todo.ID !== id);
        this.statusMessage = 'タスクが削除されました';
      } catch (error) {
        console.error('Error deleting todo:', error);
        this.statusMessage = 'タスクの削除に失敗しました';
      }
    }
  }
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
  background-color: #f0f0f0;
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

.edit-input {
  flex: 1;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
}
</style>