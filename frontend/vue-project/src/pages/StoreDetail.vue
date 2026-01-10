<template>
  <div class="store-detail-page">
    <div v-if="loading" class="loading">Loading...</div>

    <div v-else-if="error" class="error">{{ error }}</div>

    <div v-else-if="store" class="store-detail">
      <div class="store-header">
        <div class="store-icon">🏪</div>
        <h1>{{ store.name }}</h1>
        <p class="store-merchant">Merchant ID: {{ store.merchant_id }}</p>
      </div>

      <div class="store-info">
        <h2>Store Information</h2>
        <div class="info-grid">
          <div class="info-item">
            <strong>Store ID:</strong>
            <span>{{ store.id }}</span>
          </div>
          <div class="info-item">
            <strong>Created:</strong>
            <span>{{ formatDate(store.created_at) }}</span>
          </div>
        </div>
      </div>

      <div v-if="store.faqs && store.faqs.length > 0" class="faqs-section">
        <h2>Frequently Asked Questions</h2>
        <div class="faqs-container">
          <div v-for="faq in store.faqs" :key="faq.id" class="faq-item">
            <div class="faq-header">
              <span class="category-badge">{{ faq.category.name }}</span>
              <span v-if="faq.is_global" class="global-badge">Global</span>
            </div>
            
            <div v-if="faq.translations && faq.translations.length > 0" class="translations">
              <div v-for="translation in faq.translations" :key="translation.id" class="translation">
                <h4>{{ translation.question }}</h4>
                <p>{{ translation.answer }}</p>
                <small class="language-tag">{{ translation.language.toUpperCase() }}</small>
              </div>
            </div>

            <div v-else class="no-translations">
              <p class="no-data">No translations available</p>
            </div>
          </div>
        </div>
      </div>

      <div v-else class="empty-faqs">
        <p>No FAQs available for this store</p>
      </div>

      <div class="actions">
        <router-link to="/stores" class="btn-back">← Back to Stores</router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import apiClient from '../api/client'

const route = useRoute()
const store = ref(null)
const loading = ref(false)
const error = ref(null)

const loadStore = async () => {
  loading.value = true
  error.value = null
  try {
    const response = await apiClient.getStore(route.params.id)
    store.value = response.data.store
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

const formatDate = (dateString) => {
  if (!dateString) return 'N/A'
  return new Date(dateString).toLocaleDateString()
}

onMounted(() => {
  loadStore()
})
</script>

<style scoped>
.store-detail-page {
  max-width: 800px;
  margin: 0 auto;
  padding: 40px 20px;
}

.store-detail {
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

.store-header {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 40px;
  text-align: center;
}

.store-icon {
  font-size: 64px;
  margin-bottom: 20px;
}

.store-header h1 {
  margin: 0 0 10px 0;
}

.store-merchant {
  margin: 0;
  opacity: 0.9;
}

.store-info {
  padding: 30px;
}

.store-info h2 {
  margin: 0 0 20px 0;
  color: #333;
}

.info-grid {
  display: grid;
  gap: 15px;
}

.info-item {
  display: flex;
  justify-content: space-between;
  padding: 15px;
  background: #f5f5f5;
  border-radius: 4px;
}

.info-item strong {
  color: #555;
}

.info-item span {
  color: #333;
}

.actions {
  padding: 30px;
  border-top: 1px solid #eee;
}

.faqs-section {
  padding: 30px;
  border-top: 1px solid #eee;
  background: #fafafa;
}

.faqs-section h2 {
  margin: 0 0 20px 0;
  color: #333;
}

.faqs-container {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.faq-item {
  background: white;
  padding: 20px;
  border-radius: 8px;
  border-left: 4px solid #667eea;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
}

.faq-header {
  display: flex;
  gap: 10px;
  margin-bottom: 15px;
  flex-wrap: wrap;
}

.category-badge {
  background: #e3f2fd;
  color: #1976d2;
  padding: 4px 12px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 600;
}

.global-badge {
  background: #fff3e0;
  color: #f57c00;
  padding: 4px 12px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 600;
}

.translations {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.translation {
  padding: 12px;
  background: #f5f5f5;
  border-radius: 6px;
  border-left: 3px solid #667eea;
}

.translation h4 {
  margin: 0 0 8px 0;
  color: #333;
  font-size: 15px;
}

.translation p {
  margin: 0 0 10px 0;
  color: #555;
  line-height: 1.6;
}

.language-tag {
  display: inline-block;
  background: #e0e0e0;
  color: #333;
  padding: 2px 8px;
  border-radius: 4px;
  font-size: 11px;
  font-weight: 600;
}

.no-translations {
  padding: 20px;
  background: #f5f5f5;
  border-radius: 6px;
  text-align: center;
}

.no-data {
  margin: 0;
  color: #999;
  font-size: 14px;
}

.empty-faqs {
  padding: 40px 30px;
  text-align: center;
  background: #fafafa;
  border-top: 1px solid #eee;
  color: #999;
}

.btn-back {
  display: inline-block;
  background-color: #757575;
  color: white;
  padding: 10px 20px;
  border-radius: 4px;
  text-decoration: none;
  transition: background-color 0.3s;
}

.btn-back:hover {
  background-color: #616161;
}

.loading,
.error {
  text-align: center;
  padding: 40px;
  font-size: 18px;
}

.error {
  color: #f44336;
}
</style>
