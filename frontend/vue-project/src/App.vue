<script setup>
import Navbar from './components/Navbar.vue'
import { useLanguageStore } from './stores/language'
import apiClient from './api/client'

const languageStore = useLanguageStore()

// Initialize language on app load
languageStore.setLanguage(languageStore.currentLanguage)
apiClient.setLanguage(languageStore.currentLanguage)
</script>

<template>
  <div id="app" :class="{ 'rtl': languageStore.isRTL }">
    <Navbar />
    <router-view></router-view>
  </div>
</template>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
  background-color: #f5f5f5;
}

#app {
  min-height: 100vh;
}

#app.rtl {
  direction: rtl;
  text-align: right;
}
</style>
