import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { createPersistedState } from 'pinia-plugin-persistedstate'

import App from './App.vue'
import router from './router'
import { useLanguageStore } from './stores/language'
import { useUserStore } from './stores/user'
import apiClient from './api/client'

const app = createApp(App)

const pinia = createPinia()
pinia.use(createPersistedState())
app.use(pinia)

// Initialize language store and set up language
const languageStore = useLanguageStore()
apiClient.setLanguage(languageStore.currentLanguage)

// Initialize user store and set token if exists
const userStore = useUserStore()
if (userStore.token) {
  apiClient.setToken(userStore.token)
}

app.use(router)

app.mount('#app')
