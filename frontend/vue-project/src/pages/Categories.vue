<template>
  <div class="categories-page">
    <div class="header">
      <h1>FAQ Categories</h1>
      <button @click="showCreateModal = true" class="btn-primary">+ Add Category</button>
    </div>

    <div v-if="loading" class="loading">Loading...</div>

    <div v-else-if="error" class="error">{{ error }}</div>

    <div v-else class="categories-grid">
      <div v-for="category in categories" :key="category.id" class="category-card">
        <h3>{{ category.name }}</h3>
        <div class="card-actions">
          <button @click="editCategory(category)" class="btn-edit">Edit</button>
          <button @click="deleteCategory(category.id)" class="btn-delete">Delete</button>
        </div>
      </div>
    </div>

    <!-- Create/Edit Modal -->
    <div v-if="showCreateModal || showEditModal" class="modal-overlay" @click="closeModal">
      <div class="modal" @click.stop>
        <h2>{{ showEditModal ? 'Edit Category' : 'Create Category' }}</h2>
        <form @submit.prevent="saveCategory">
          <div class="form-group">
            <label>Category Name:</label>
            <input v-model="formData.name" type="text" required />
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
import { ref, onMounted } from 'vue'
import { useUserStore } from '../stores/user'
import apiClient from '../api/client'

const userStore = useUserStore()
const categories = ref([])
const loading = ref(false)
const error = ref(null)
const showCreateModal = ref(false)
const showEditModal = ref(false)
const formData = ref({ name: '' })
const editingId = ref(null)

apiClient.setToken(userStore.token)

const loadCategories = async () => {
  loading.value = true
  error.value = null
  try {
    const response = await apiClient.getCategories()
    categories.value = response.data.categories || []
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

const editCategory = (category) => {
  formData.value = { name: category.name }
  editingId.value = category.id
  showEditModal.value = true
}

const saveCategory = async () => {
  try {
    if (showEditModal.value) {
      await apiClient.updateCategory(editingId.value, formData.value)
    } else {
      await apiClient.createCategory(formData.value)
    }
    closeModal()
    loadCategories()
  } catch (err) {
    alert('Error: ' + err.message)
  }
}

const deleteCategory = async (id) => {
  if (!confirm('Are you sure you want to delete this category?')) return

  try {
    await apiClient.deleteCategory(id)
    loadCategories()
  } catch (err) {
    alert('Error: ' + err.message)
  }
}

const closeModal = () => {
  showCreateModal.value = false
  showEditModal.value = false
  formData.value = { name: '' }
  editingId.value = null
}

onMounted(() => {
  loadCategories()
})
</script>

<style scoped>
.categories-page {
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

.categories-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 20px;
}

.category-card {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.category-card h3 {
  margin: 0 0 15px 0;
  color: #333;
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

.btn-edit {
  background-color: #2196F3;
  color: white;
  border: none;
  padding: 8px 16px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
}

.btn-edit:hover {
  background-color: #0b7dda;
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

.btn-delete:hover {
  background-color: #da190b;
}

.btn-secondary {
  background-color: #757575;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
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
  min-width: 400px;
  max-width: 90%;
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

.form-group input {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
  box-sizing: border-box;
}

.modal-actions {
  display: flex;
  gap: 10px;
  justify-content: flex-end;
}
</style>
