<template>
  <nav class="navbar" :class="{ 'rtl': languageStore.isRTL }">
    <div class="nav-container">
      <router-link to="/dashboard" class="logo">{{ languageStore.t('navbar_faq_system') }}</router-link>

      <div class="nav-links">
        <router-link to="/dashboard" v-if="userStore.isAuthenticated">{{ languageStore.t('nav_dashboard') }}</router-link>
        <router-link to="/categories" v-if="isAdmin">{{ languageStore.t('nav_categories') }}</router-link>
        <router-link to="/faqs" v-if="isAdmin || isMerchant">{{ languageStore.t('nav_faqs') }}</router-link>
        <router-link to="/stores">{{ languageStore.t('nav_stores') }}</router-link>

        <!-- Language Selector -->
        <div class="language-selector">
          <select @change="changeLanguage" :value="languageStore.currentLanguage" class="language-dropdown">
            <option v-for="lang in languageStore.supportedLanguages" :key="lang.code" :value="lang.code">
              {{ lang.flag }} {{ lang.name }}
            </option>
          </select>
        </div>

        <div class="user-menu" v-if="userStore.isAuthenticated">
          <span class="user-name">{{ userStore.user?.name }}</span>
          <button @click="handleLogout" class="btn-logout">{{ languageStore.t('nav_logout') }}</button>
        </div>

        <router-link to="/login" v-else class="btn-login">{{ languageStore.t('nav_login') }}</router-link>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '../stores/user'
import { useLanguageStore } from '../stores/language'
import apiClient from '../api/client'

const router = useRouter()
const userStore = useUserStore()
const languageStore = useLanguageStore()

const isAdmin = computed(() => userStore.user?.role === 'admin')
const isMerchant = computed(() => userStore.user?.role === 'merchant')

const changeLanguage = (e) => {
  const lang = e.target.value
  languageStore.setLanguage(lang)
  apiClient.setLanguage(lang)
}

const handleLogout = () => {
  userStore.logout()
  router.push('/login')
}
</script>

<style scoped>
.navbar {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.navbar.rtl {
  direction: rtl;
}

.nav-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 15px 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo {
  font-size: 24px;
  font-weight: bold;
  color: white;
  text-decoration: none;
}

.nav-links {
  display: flex;
  align-items: center;
  gap: 20px;
}

.nav-links a {
  color: white;
  text-decoration: none;
  padding: 8px 16px;
  border-radius: 4px;
  transition: background-color 0.3s;
}

.nav-links a:hover,
.nav-links a.router-link-active {
  background-color: rgba(255, 255, 255, 0.2);
}

.language-selector {
  position: relative;
}

.language-dropdown {
  background-color: rgba(255, 255, 255, 0.2);
  color: white;
  border: 1px solid white;
  padding: 8px 12px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
  transition: background-color 0.3s;
}

.language-dropdown:hover {
  background-color: rgba(255, 255, 255, 0.3);
}

.language-dropdown option {
  background-color: #667eea;
  color: white;
}

.user-menu {
  display: flex;
  align-items: center;
  gap: 15px;
  margin-left: 10px;
  padding-left: 20px;
  border-left: 1px solid rgba(255, 255, 255, 0.3);
}

.user-name {
  font-weight: 500;
}

.btn-logout,
.btn-login {
  background-color: rgba(255, 255, 255, 0.2);
  color: white;
  border: 1px solid white;
  padding: 8px 16px;
  border-radius: 4px;
  cursor: pointer;
  text-decoration: none;
  display: inline-block;
  transition: background-color 0.3s;
}

.btn-logout:hover,
.btn-login:hover {
  background-color: rgba(255, 255, 255, 0.3);
}
</style>
