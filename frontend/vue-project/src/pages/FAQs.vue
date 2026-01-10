<template>
  <div class="faqs-page">
    <div class="header">
      <h1>FAQs Management</h1>
      <button @click="showCreateModal = true" class="btn-primary">+ Add FAQ</button>
    </div>

    <div class="search-box">
      <input 
        v-model="searchQuery" 
        @input="handleSearchInput"
        type="text" 
        placeholder="Search FAQs by question, answer, or category..." 
        class="search-input"
      />
      <span v-if="searchQuery || total" class="search-results">{{ total }} result(s)</span>
    </div>

    <div v-if="loading" class="loading">Loading...</div>

    <div v-else-if="error" class="error">{{ error }}</div>

    <div v-else class="faqs-list">
      <div v-if="faqs.length === 0" class="no-results">
        <p>No FAQs found{{ searchQuery ? ' matching your search' : '' }}</p>
      </div>
      <div v-for="faq in faqs" :key="faq.id" class="faq-card">
        <div class="faq-header">
          <h3>{{ getFirstTranslation(faq.translations).question || 'FAQ #' + faq.id }}</h3>
          <span class="category-badge">{{ getCategoryName(faq.category_id) }}</span>
        </div>
        <p class="faq-answer">{{ getFirstTranslation(faq.translations).answer || 'No answer provided' }}</p>
        <div class="faq-meta">
          <span v-if="faq.store_id">Store ID: {{ faq.store_id }}</span>
          <span v-else class="global-badge">Global FAQ</span>
        </div>
        <div class="card-actions">
          <button @click="editFAQ(faq)" class="btn-edit">Edit</button>
          <button @click="deleteFAQ(faq.id)" class="btn-delete">Delete</button>
        </div>
      </div>

      <div v-if="faqs.length > 0" class="pagination">
        <button :disabled="page <= 1" @click="changePage(page - 1)">Prev</button>
        <span>Page {{ page }} of {{ totalPages }}</span>
        <button :disabled="page >= totalPages" @click="changePage(page + 1)">Next</button>
      </div>
    </div>

    <!-- Create/Edit Modal -->
    <div v-if="showCreateModal || showEditModal" class="modal-overlay" @click="closeModal">
      <div class="modal" @click.stop>
        <h2>{{ showEditModal ? 'Edit FAQ' : 'Create FAQ' }}</h2>
        <form @submit.prevent="saveFAQ">
          <div class="form-group">
            <label>Category:</label>
            <select v-model="formData.category_id" required>
              <option value="">Select category</option>
              <option v-for="cat in categories" :key="cat.id" :value="cat.id">
                {{ cat.name }}
              </option>
            </select>
          </div>

          <div class="form-group">
            <label>Store (optional, leave empty for global FAQ):</label>
            <select v-model="formData.store_id" :disabled="!isAdmin">
              <option :value="null">Global FAQ</option>
              <option v-for="store in availableStores" :key="store.id" :value="store.id">
                {{ store.name }}
              </option>
            </select>
            <small v-if="!isAdmin" class="hint">Only admins can create global FAQs</small>
          </div>

          <div class="form-group">
            <label>Translations:</label>
            <div class="translations-container">
              <div v-if="showEditModal" class="all-languages-info">
                <p>All translations for this FAQ:</p>
              </div>
              <div v-for="(translation, index) in formData.translations" :key="index" class="translation-group">
                <div class="translation-header">
                  <select v-model="translation.language" class="language-select" required>
                    <option value="">Select Language</option>
                    <option value="en">English</option>
                    <option value="ar">العربية (Arabic)</option>
                    <option value="fr">Français (French)</option>
                    <option value="es">Español (Spanish)</option>
                    <option value="de">Deutsch (German)</option>
                  </select>
                  <button type="button" @click="removeTranslation(index)" class="btn-remove-small">✕ Remove</button>
                </div>
                <input v-model="translation.question" placeholder="Question" required />
                <textarea v-model="translation.answer" placeholder="Answer" rows="3" required></textarea>
              </div>
              <button type="button" @click="addTranslation" class="btn-secondary">+ Add Translation</button>
            </div>
          </div>

          <div class="modal-actions">
            <button type="submit" class="btn-primary">Save</button>
            <button type="button" @click="closeModal" class="btn-secondary">Cancel</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useUserStore } from '../stores/user'
import { useLanguageStore } from '../stores/language'
import apiClient from '../api/client'

const userStore = useUserStore()
const languageStore = useLanguageStore()
const faqs = ref([])
const categories = ref([])
const stores = ref([])
const loading = ref(false)
const error = ref(null)
const showCreateModal = ref(false)
const showEditModal = ref(false)
const searchQuery = ref('')
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
let searchDebounce = null
const formData = ref({
  category_id: '',
  store_id: null,
  translations: [{ language: 'en', question: '', answer: '' }],
})
const editingId = ref(null)

// Computed: Filter stores based on user role
const availableStores = computed(() => {
  if (userStore.user?.role === 'admin') {
    return stores.value
  }
  // Merchants can only select their own store
  return stores.value.slice(0, 1)
})

const isAdmin = computed(() => userStore.user?.role === 'admin')

const totalPages = computed(() => {
  if (!total.value || !pageSize.value) return 1
  return Math.max(1, Math.ceil(total.value / pageSize.value))
})

apiClient.setToken(userStore.token)

// Helper: Get first translation or fallback
const getFirstTranslation = (translations) => {
  if (!translations || translations.length === 0) {
    return { question: 'No question', answer: 'No answer' }
  }
  return translations[0]
}

const loadFAQs = async () => {
  loading.value = true
  error.value = null
  try {
  const response = await apiClient.getFAQs({
    search: searchQuery.value.trim(),
    page: page.value,
    page_size: pageSize.value,
  })
  faqs.value = response.data.faqs || []
  total.value = response.data.total || 0
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

const loadCategories = async () => {
  try {
    const response = await apiClient.getCategories()
    categories.value = response.data.categories || []
  } catch (err) {
    console.error('Failed to load categories:', err)
  }
}

const loadStores = async () => {
  try {
    const response = await apiClient.getStores()
    stores.value = response.data.stores || []
  } catch (err) {
    console.error('Failed to load stores:', err)
  }
}

const handleSearchInput = () => {
  if (searchDebounce) clearTimeout(searchDebounce)
  searchDebounce = setTimeout(() => {
    page.value = 1
    loadFAQs()
  }, 300)
}

const changePage = (nextPage) => {
  if (nextPage < 1 || nextPage > totalPages.value) return
  page.value = nextPage
  loadFAQs()
}

const getCategoryName = (categoryId) => {
  const category = categories.value.find((c) => c.id === categoryId)
  return category ? category.name : 'Unknown'
}

const addTranslation = () => {
  formData.value.translations.push({ language: '', question: '', answer: '' })
}

const removeTranslation = (index) => {
  formData.value.translations.splice(index, 1)
}

const editFAQ = async (faq) => {
  loading.value = true
  try {
    const response = await apiClient.getFAQ(faq.id, { include_all_translations: true })
    const detailedFAQ = response.data.faq
    formData.value = {
      category_id: detailedFAQ.category_id,
      store_id: detailedFAQ.store_id || null,
      translations: detailedFAQ.translations && detailedFAQ.translations.length > 0
        ? [...detailedFAQ.translations]
        : [{ language: 'en', question: '', answer: '' }],
    }
    editingId.value = faq.id
    showEditModal.value = true
  } catch (err) {
    alert('Error loading FAQ: ' + err.message)
  } finally {
    loading.value = false
  }
}

const saveFAQ = async () => {
  try {
    if (showEditModal.value) {
      await apiClient.updateFAQ(editingId.value, formData.value)
    } else {
      await apiClient.createFAQ(formData.value)
    }
    closeModal()
    loadFAQs()
  } catch (err) {
    alert('Error: ' + err.message)
  }
}

const deleteFAQ = async (id) => {
  if (!confirm('Are you sure you want to delete this FAQ?')) return

  try {
    await apiClient.deleteFAQ(id)
    loadFAQs()
  } catch (err) {
    alert('Error: ' + err.message)
  }
}

const closeModal = () => {
  showCreateModal.value = false
  showEditModal.value = false
  formData.value = {
    category_id: '',
    store_id: null,
    translations: [{ language: 'en', question: '', answer: '' }],
  }
  editingId.value = null
}

onMounted(() => {
  loadFAQs()
  loadCategories()
  loadStores()
})
</script>

<style scoped>
.faqs-page {
  max-width: 1200px;
  margin: 0 auto;
  padding: 40px 20px;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 30px;
}

.search-box {
  margin-bottom: 20px;
  display: flex;
  gap: 10px;
  align-items: center;
}

.search-input {
  flex: 1;
  padding: 12px 16px;
  border: 2px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
  transition: border-color 0.3s;
}

.search-input:focus {
  outline: none;
  border-color: #4CAF50;
}

.search-results {
  color: #666;
  font-size: 14px;
  padding: 0 10px;
}

.no-results {
  text-align: center;
  padding: 40px;
  color: #999;
  font-size: 16px;
}

.faqs-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.pagination {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
  margin-top: 10px;
}

.pagination button {
  padding: 8px 14px;
  border: 1px solid #ddd;
  background: #fff;
  border-radius: 4px;
  cursor: pointer;
  transition: background 0.2s, border-color 0.2s;
}

.pagination button:disabled {
  cursor: not-allowed;
  opacity: 0.6;
}

.pagination button:not(:disabled):hover {
  background: #f5f5f5;
  border-color: #ccc;
}

.faq-card {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.faq-header {
  display: flex;
  justify-content: space-between;
  align-items: start;
  margin-bottom: 10px;
}

.faq-header h3 {
  margin: 0;
  color: #333;
  flex: 1;
}

.category-badge {
  background: #e3f2fd;
  color: #1976d2;
  padding: 4px 12px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 500;
}

.global-badge {
  background: #fff3e0;
  color: #f57c00;
  padding: 4px 12px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 500;
}

.faq-answer {
  color: #666;
  margin: 10px 0;
}

.faq-meta {
  font-size: 14px;
  color: #888;
  margin-bottom: 15px;
}

.card-actions {
  display: flex;
  gap: 10px;
}

.btn-primary {
  background-color: #4CAF50;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
}

.btn-primary:hover {
  background-color: #45a049;
}

.btn-secondary {
  background-color: #757575;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
  margin-top: 10px;
}

.btn-edit {
  background-color: #2196F3;
  color: white;
  border: none;
  padding: 8px 16px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
}

.btn-delete {
  background-color: #f44336;
  color: white;
  border: none;
  padding: 8px 16px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
}

.btn-remove {
  background-color: #ff5722;
  color: white;
  border: none;
  padding: 5px 10px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 12px;
  margin-top: 5px;
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

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal {
  background: white;
  padding: 30px;
  border-radius: 8px;
  min-width: 500px;
  max-width: 90%;
  max-height: 90vh;
  overflow-y: auto;
}

.modal h2 {
  margin: 0 0 20px 0;
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  margin-bottom: 5px;
  color: #555;
  font-weight: 500;
}

.form-group input,
.form-group select,
.form-group textarea {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
  box-sizing: border-box;
  margin-bottom: 5px;
}

.translation-group {
  border: 1px solid #eee;
  padding: 15px;
  border-radius: 4px;
  margin-bottom: 10px;
  background: #fafafa;
}

.translation-header {
  display: flex;
  gap: 10px;
  margin-bottom: 10px;
  align-items: center;
}

.language-select {
  flex: 1;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
  background-color: white;
}

.language-select:focus {
  outline: none;
  border-color: #4CAF50;
}

.btn-remove-small {
  background-color: #f44336;
  color: white;
  border: none;
  padding: 8px 12px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
  white-space: nowrap;
}

.btn-remove-small:hover {
  background-color: #da190b;
}

.translations-container {
  border: 1px solid #e0e0e0;
  padding: 15px;
  border-radius: 4px;
  background: #f5f5f5;
}

.all-languages-info {
  margin-bottom: 15px;
  padding: 10px;
  background: #e3f2fd;
  border-left: 4px solid #1976d2;
  border-radius: 2px;
}

.all-languages-info p {
  margin: 0;
  color: #1976d2;
  font-weight: 500;
  font-size: 13px;
}

.modal-actions {
  display: flex;
  gap: 10px;
  justify-content: flex-end;
  margin-top: 20px;
}
</style>
