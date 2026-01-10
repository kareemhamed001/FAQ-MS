import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

const API_BASE = 'http://localhost:8080'

export const useUserStore = defineStore(
  'user',
  () => {
    const user = ref(null)
    const token = ref(null)
    const isAuthenticated = computed(() => user.value !== null)

    async function login(email, password) {
      const res = await fetch(`${API_BASE}/auth/login`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      })

      if (!res.ok) {
        throw new Error('Login failed')
      }

      const data = await res.json()
      user.value = data.data.user
      token.value = data.data.token
      return data.data
    }

    async function register(name, email, password, role) {
      const res = await fetch(`${API_BASE}/auth/register`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name, email, password, role }),
      })

      if (!res.ok) {
        throw new Error('Registration failed')
      }

      const data = await res.json()
      user.value = data.user
      token.value = data.token
      return data
    }

    function logout() {
      user.value = null
      token.value = null
    }

    return { user, token, isAuthenticated, login, register, logout }
  },
  {
    persist: {
      enabled: true,
      strategies: [
        {
          key: 'user',
          storage: localStorage,
          paths: ['user', 'token'],
        },
      ],
    },
  },
)
