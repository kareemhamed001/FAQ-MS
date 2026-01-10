<template>
  <div class="dashboard">
    <h1>{{ languageStore.t('dashboard_welcome') }}, {{ userStore.user?.name }}!</h1>
    <div class="user-info">
      <p><strong>{{ languageStore.t('dashboard_email') }}:</strong> {{ userStore.user?.email }}</p>
      <p><strong>{{ languageStore.t('dashboard_role') }}:</strong> {{ userStore.user?.role }}</p>
    </div>

    <div class="quick-links">
      <router-link v-if="isAdmin" to="/categories" class="card">
        <h3>📁 {{ languageStore.t('dashboard_categories') }}</h3>
        <p>{{ languageStore.t('dashboard_categories_desc') }}</p>
      </router-link>

      <router-link v-if="isAdmin || isMerchant" to="/faqs" class="card">
        <h3>❓ {{ languageStore.t('dashboard_faqs') }}</h3>
        <p>{{ languageStore.t('dashboard_faqs_desc') }}</p>
      </router-link>

      <router-link to="/stores" class="card">
        <h3>🏪 {{ languageStore.t('dashboard_stores') }}</h3>
        <p>{{ languageStore.t('dashboard_stores_desc') }}</p>
      </router-link>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useUserStore } from '../stores/user'
import { useLanguageStore } from '../stores/language'

const userStore = useUserStore()
const languageStore = useLanguageStore()

const isAdmin = computed(() => userStore.user?.role === 'admin')
const isMerchant = computed(() => userStore.user?.role === 'merchant')
</script>

<style scoped>
.dashboard {
  max-width: 1200px;
  margin: 0 auto;
  padding: 40px 20px;
}

h1 {
  color: #333;
  margin-bottom: 20px;
}

.user-info {
  background: #f5f5f5;
  padding: 20px;
  border-radius: 8px;
  margin-bottom: 40px;
}

.user-info p {
  margin: 10px 0;
  color: #555;
}

.quick-links {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 20px;
}

.card {
  background: white;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  text-decoration: none;
  color: inherit;
  transition: transform 0.3s, box-shadow 0.3s;
}

.card:hover {
  transform: translateY(-5px);
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.15);
}

.card h3 {
  margin: 0 0 10px 0;
  color: #4CAF50;
}

.card p {
  margin: 0;
  color: #666;
}
</style>
