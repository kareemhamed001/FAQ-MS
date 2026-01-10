<template>
    <div class="login-container">
        <h2>Login</h2>
        <form @submit.prevent="handleLogin">
        <div class="form-group">
            <label for="email">Email:</label>
            <input type="email" v-model="email" id="email" required />
        </div>  
        <div class="form-group">
            <label for="password">Password:</label>
            <input type="password" v-model="password" id="password" required /> 
        </div>
        <button type="submit">Login</button>
        <p class="register-link">
          Don't have an account? <router-link to="/register">Register here</router-link>
        </p>
        </form> 
    </div>  
</template> 
<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
//import user store
import { useUserStore } from '../stores/user'

const email = ref('');
const password = ref('');   
const router = useRouter(); 
const userStore = useUserStore();
const handleLogin = async () => {
    try {
        // Call the login action from the user store with positional args (email, password)
        await userStore.login(email.value, password.value);
        router.push('/dashboard');  
    } catch (error) {
        alert('Login failed: ' + error.message);
    }
};  


</script>

<style scoped>
.login-container {
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

input {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
  box-sizing: border-box;
}

input:focus {
  outline: none;
  border-color: #4CAF50;
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

.register-link {
  text-align: center;
  margin-top: 20px;
  color: #666;
}

.register-link a {
  color: #4CAF50;
  text-decoration: none;
}

.register-link a:hover {
  text-decoration: underline;
}
</style>