<template>
  <div class="stores-page">
    <h1>Merchant Stores</h1>

    <div v-if="loading" class="loading">Loading...</div>

    <div v-else-if="error" class="error">{{ error }}</div>

    <div v-else class="stores-grid">
      <div v-for="store in stores" :key="store.id" class="store-card">
        <div class="store-icon">🏪</div>
        <h3>{{ store.name }}</h3>
        <p class="store-merchant">Merchant ID: {{ store.merchant_id }}</p>
        <router-link :to="`/stores/${store.id}`" class="btn-view">View Details</router-link>
      </div>
    </div>

    <div v-if="!loading && stores.length === 0" class="empty-state">
      <p>No stores found</p>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import apiClient from '../api/client'

const stores = ref([])
const loading = ref(false)
const error = ref(null)

const loadStores = async () => {
  loading.value = true
  error.value = null
  try {
    const response = await apiClient.getStores()

    stores.value = response.data.stores || []
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadStores()
})
</script>

<style scoped>
.stores-page {
  max-width: 1200px;
  margin: 0 auto;
  padding: 40px 20px;
}

h1 {
  margin-bottom: 30px;
  color: #333;
}

.stores-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 20px;
}

.store-card {
  background: white;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  text-align: center;
  transition: transform 0.3s;
}

.store-card:hover {
  transform: translateY(-5px);
}

.store-icon {
  font-size: 48px;
  margin-bottom: 15px;
}

.store-card h3 {
  margin: 0 0 10px 0;
  color: #333;
}

.store-merchant {
  color: #888;
  font-size: 14px;
  margin-bottom: 20px;
}

.btn-view {
  display: inline-block;
  background-color: #4CAF50;
  color: white;
  padding: 10px 20px;
  border-radius: 4px;
  text-decoration: none;
  transition: background-color 0.3s;
}

.btn-view:hover {
  background-color: #45a049;
}

.loading,
.error,
.empty-state {
  text-align: center;
  padding: 40px;
  font-size: 18px;
}

.error {
  color: #f44336;
}

.empty-state {
  color: #888;
}
</style>
