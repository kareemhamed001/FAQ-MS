<template>
  <div class="register-container">
    <h2>Register</h2>
    <form @submit.prevent="handleRegister">
      <div class="form-group">
        <label for="name">Name:</label>
        <input type="text" v-model="name" id="name" required />
      </div>

      <div class="form-group">
        <label for="email">Email:</label>
        <input type="email" v-model="email" id="email" required />
      </div>

      <div class="form-group">
        <label for="password">Password:</label>
        <input type="password" v-model="password" id="password" required />
        <small>Must be at least 8 characters with a number and special character</small>
      </div>

      <div class="form-group">
        <label for="role">Account Type:</label>
        <select v-model="role" id="role" required>
          <option value="">Select type</option>
          <option value="merchant">Merchant</option>
          <option value="customer">Customer</option>
        </select>
      </div>

      <button type="submit">Register</button>
      <p class="login-link">
        Already have an account? <router-link to="/login">Login here</router-link>
      </p>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '../stores/user'

const name = ref('')
const email = ref('')
const password = ref('')
const role = ref('')
const router = useRouter()
const userStore = useUserStore()

const handleRegister = async () => {
  try {
    await userStore.register(name.value, email.value, password.value, role.value)
    router.push('/dashboard')
  } catch (error) {
    alert('Registration failed: ' + error.message)
  }
}
</script>

<style scoped>
.register-container {
  max-width: 400px;
  margin: 50px auto;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  background: white;
}

h2 {
  text-align: center;
  margin-bottom: 30px;
  color: #333;
}

.form-group {
  margin-bottom: 20px;
}

label {
  display: block;
  margin-bottom: 5px;
  color: #555;
  font-weight: 500;
}

input,
select {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
  box-sizing: border-box;
}

input:focus,
select:focus {
  outline: none;
  border-color: #4CAF50;
}

small {
  display: block;
  margin-top: 5px;
  color: #888;
  font-size: 12px;
}

button {
  width: 100%;
  padding: 12px;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.3s;
}

button:hover {
  background-color: #45a049;
}

.login-link {
  text-align: center;
  margin-top: 20px;
  color: #666;
}

.login-link a {
  color: #4CAF50;
  text-decoration: none;
}

.login-link a:hover {
  text-decoration: underline;
}
</style>
