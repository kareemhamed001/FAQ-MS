const API_BASE = 'http://localhost:8080'

export class ApiClient {
  constructor(token = null, language = 'en') {
    this.token = token
    this.language = language
  }

  setToken(token) {
    this.token = token
  }

  setLanguage(language) {
    this.language = language
  }

  async request(endpoint, options = {}) {
    const headers = {
      'Content-Type': 'application/json',
      'Accept-Language': this.language,
      ...options.headers,
    }

    if (this.token) {
      headers['Authorization'] = `Bearer ${this.token}`
    }

    const response = await fetch(`${API_BASE}${endpoint}`, {
      ...options,
      headers,
    })

    const data = await response.json()

    if (!response.ok) {
      throw new Error(data.error?.message || 'Request failed')
    }

    return data
  }

  // FAQ Categories
  async getCategories() {
    return this.request('/api/faq-categories/')
  }

  async getCategory(id) {
    return this.request(`/api/faq-categories/${id}`)
  }

  async createCategory(categoryData) {
    return this.request('/api/faq-categories/', {
      method: 'POST',
      body: JSON.stringify(categoryData),
    })
  }

  async updateCategory(id, categoryData) {
    return this.request(`/api/faq-categories/${id}`, {
      method: 'PUT',
      body: JSON.stringify(categoryData),
    })
  }

  async deleteCategory(id) {
    return this.request(`/api/faq-categories/${id}`, {
      method: 'DELETE',
    })
  }

  // FAQs
  async getFAQs(params = {}) {
    const query = new URLSearchParams()
    if (params.search) query.append('search', params.search)
    if (params.page) query.append('page', params.page)
    if (params.page_size) query.append('page_size', params.page_size)
    if (params.sort) query.append('sort', params.sort)

    const qs = query.toString()
    const url = qs ? `/api/faqs/?${qs}` : '/api/faqs/'
    return this.request(url)
  }

  async getFAQ(id, params = {}) {
    const query = new URLSearchParams()
    if (params.include_all_translations) query.append('include_all_translations', 'true')
    const qs = query.toString()
    const url = qs ? `/api/faqs/${id}?${qs}` : `/api/faqs/${id}`
    return this.request(url)
  }

  async createFAQ(faqData) {
    return this.request('/api/faqs/', {
      method: 'POST',
      body: JSON.stringify(faqData),
    })
  }

  async updateFAQ(id, faqData) {
    return this.request(`/api/faqs/${id}`, {
      method: 'PUT',
      body: JSON.stringify(faqData),
    })
  }

  async deleteFAQ(id) {
    return this.request(`/api/faqs/${id}`, {
      method: 'DELETE',
    })
  }

  // Stores
  async getStores() {
    return this.request('/api/stores/')
  }

  async getStore(id) {
    return this.request(`/api/stores/${id}`)
  }
}

export default new ApiClient()
