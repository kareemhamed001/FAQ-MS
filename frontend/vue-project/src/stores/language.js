import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useLanguageStore = defineStore('language', () => {
  const currentLanguage = ref(localStorage.getItem('language') || 'en')
  const supportedLanguages = [
    { code: 'en', name: 'English', flag: '🇺🇸' },
    { code: 'ar', name: 'العربية', flag: '🇸🇦' },
  ]

  const translations = {
    en: {
      // Navigation
      nav_dashboard: 'Dashboard',
      nav_categories: 'Categories',
      nav_faqs: 'FAQs',
      nav_stores: 'Stores',
      nav_logout: 'Logout',
      nav_login: 'Login',

      // Common
      loading: 'Loading...',
      error: 'An error occurred',
      success: 'Success',
      add: 'Add',
      edit: 'Edit',
      delete: 'Delete',
      save: 'Save',
      cancel: 'Cancel',
      back: 'Back',
      search: 'Search',

      // Login
      login_title: 'Login',
      login_email: 'Email',
      login_password: 'Password',
      login_button: 'Login',
      login_register_link: "Don't have an account?",
      login_register_here: 'Register here',

      // Register
      register_title: 'Register',
      register_name: 'Name',
      register_email: 'Email',
      register_password: 'Password',
      register_password_hint: 'Must be at least 8 characters with a number and special character',
      register_account_type: 'Account Type',
      register_merchant: 'Merchant',
      register_customer: 'Customer',
      register_button: 'Register',
      register_login_link: 'Already have an account?',
      register_login_here: 'Login here',

      // Dashboard
      dashboard_welcome: 'Welcome',
      dashboard_email: 'Email',
      dashboard_role: 'Role',
      dashboard_categories: 'FAQ Categories',
      dashboard_categories_desc: 'Manage FAQ categories',
      dashboard_faqs: 'FAQs',
      dashboard_faqs_desc: 'Manage frequently asked questions',
      dashboard_stores: 'Stores',
      dashboard_stores_desc: 'Browse merchant stores',

      // Categories
      categories_title: 'FAQ Categories',
      categories_add: 'Add Category',
      categories_name: 'Category Name',
      categories_edit: 'Edit Category',
      categories_create: 'Create Category',
      categories_edit_modal: 'Edit Category',
      categories_delete_confirm: 'Are you sure you want to delete this category?',

      // FAQs
      faqs_title: 'FAQs Management',
      faqs_add: 'Add FAQ',
      faqs_category: 'Category',
      faqs_store: 'Store (optional, leave empty for global FAQ)',
      faqs_global: 'Global FAQ',
      faqs_translations: 'Translations',
      faqs_language: 'Language (e.g., en)',
      faqs_question: 'Question',
      faqs_answer: 'Answer',
      faqs_add_translation: 'Add Translation',
      faqs_remove_translation: 'Remove',
      faqs_edit: 'Edit FAQ',
      faqs_create: 'Create FAQ',
      faqs_delete_confirm: 'Are you sure you want to delete this FAQ?',

      // Stores
      stores_title: 'Merchant Stores',
      stores_no_stores: 'No stores found',
      stores_merchant_id: 'Merchant ID',
      stores_view_details: 'View Details',
      stores_info: 'Store Information',
      stores_store_id: 'Store ID',
      stores_created: 'Created',
      stores_faqs: 'Frequently Asked Questions',
      stores_no_faqs: 'No FAQs available for this store',
      stores_global: 'Global',
      stores_no_translations: 'No translations available',

      // Navbar
      navbar_faq_system: 'FAQ System',
    },
    ar: {
      // Navigation
      nav_dashboard: 'لوحة التحكم',
      nav_categories: 'الفئات',
      nav_faqs: 'الأسئلة الشائعة',
      nav_stores: 'المتاجر',
      nav_logout: 'تسجيل الخروج',
      nav_login: 'تسجيل الدخول',

      // Common
      loading: 'جاري التحميل...',
      error: 'حدث خطأ',
      success: 'نجح',
      add: 'إضافة',
      edit: 'تعديل',
      delete: 'حذف',
      save: 'حفظ',
      cancel: 'إلغاء',
      back: 'رجوع',
      search: 'بحث',

      // Login
      login_title: 'تسجيل الدخول',
      login_email: 'البريد الإلكتروني',
      login_password: 'كلمة المرور',
      login_button: 'تسجيل الدخول',
      login_register_link: 'ليس لديك حساب؟',
      login_register_here: 'سجل هنا',

      // Register
      register_title: 'إنشاء حساب',
      register_name: 'الاسم',
      register_email: 'البريد الإلكتروني',
      register_password: 'كلمة المرور',
      register_password_hint: 'يجب أن تكون 8 أحرف على الأقل وتحتوي على رقم وحرف خاص',
      register_account_type: 'نوع الحساب',
      register_merchant: 'تاجر',
      register_customer: 'عميل',
      register_button: 'إنشاء حساب',
      register_login_link: 'هل لديك حساب بالفعل؟',
      register_login_here: 'سجل الدخول هنا',

      // Dashboard
      dashboard_welcome: 'أهلا وسهلا',
      dashboard_email: 'البريد الإلكتروني',
      dashboard_role: 'الدور',
      dashboard_categories: 'فئات الأسئلة الشائعة',
      dashboard_categories_desc: 'إدارة فئات الأسئلة الشائعة',
      dashboard_faqs: 'الأسئلة الشائعة',
      dashboard_faqs_desc: 'إدارة الأسئلة الشائعة',
      dashboard_stores: 'المتاجر',
      dashboard_stores_desc: 'استعرض متاجر التجار',

      // Categories
      categories_title: 'فئات الأسئلة الشائعة',
      categories_add: 'إضافة فئة',
      categories_name: 'اسم الفئة',
      categories_edit: 'تعديل الفئة',
      categories_create: 'إنشاء فئة',
      categories_edit_modal: 'تعديل الفئة',
      categories_delete_confirm: 'هل أنت متأكد من حذف هذه الفئة؟',

      // FAQs
      faqs_title: 'إدارة الأسئلة الشائعة',
      faqs_add: 'إضافة سؤال',
      faqs_category: 'الفئة',
      faqs_store: 'المتجر (اختياري، اتركه فارغًا للأسئلة العامة)',
      faqs_global: 'سؤال عام',
      faqs_translations: 'الترجمات',
      faqs_language: 'اللغة (مثلا: ar)',
      faqs_question: 'السؤال',
      faqs_answer: 'الإجابة',
      faqs_add_translation: 'إضافة ترجمة',
      faqs_remove_translation: 'حذف',
      faqs_edit: 'تعديل السؤال',
      faqs_create: 'إنشاء سؤال',
      faqs_delete_confirm: 'هل أنت متأكد من حذف هذا السؤال؟',

      // Stores
      stores_title: 'متاجر التجار',
      stores_no_stores: 'لم يتم العثور على متاجر',
      stores_merchant_id: 'معرّف التاجر',
      stores_view_details: 'عرض التفاصيل',
      stores_info: 'معلومات المتجر',
      stores_store_id: 'معرّف المتجر',
      stores_created: 'تاريخ الإنشاء',
      stores_faqs: 'الأسئلة الشائعة',
      stores_no_faqs: 'لا توجد أسئلة شائعة لهذا المتجر',
      stores_global: 'عام',
      stores_no_translations: 'لا توجد ترجمات متاحة',

      // Navbar
      navbar_faq_system: 'نظام الأسئلة الشائعة',
    },
  }

  function setLanguage(lang) {
    if (supportedLanguages.some((l) => l.code === lang)) {
      currentLanguage.value = lang
      localStorage.setItem('language', lang)
      document.documentElement.lang = lang
      if (lang === 'ar') {
        document.documentElement.dir = 'rtl'
      } else {
        document.documentElement.dir = 'ltr'
      }
    }
  }

  function t(key) {
    return translations[currentLanguage.value]?.[key] || translations.en[key] || key
  }

  const isRTL = computed(() => currentLanguage.value === 'ar')

  return {
    currentLanguage,
    supportedLanguages,
    setLanguage,
    t,
    isRTL,
  }
})
