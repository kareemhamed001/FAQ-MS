import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '../stores/user'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/dashboard',
    },
    {
      path: '/login',
      name: 'Login',
      component: () => import('../pages/Login.vue'),
      meta: { public: true },
    },
    {
      path: '/register',
      name: 'Register',
      component: () => import('../pages/Register.vue'),
      meta: { public: true },
    },
    {
      path: '/dashboard',
      name: 'Dashboard',
      component: () => import('../pages/Dashboard.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/categories',
      name: 'Categories',
      component: () => import('../pages/Categories.vue'),
      meta: { requiresAuth: true, role: 'admin' },
    },
    {
      path: '/faqs',
      name: 'FAQs',
      component: () => import('../pages/FAQs.vue'),
      meta: { requiresAuth: true, roles: ['admin', 'merchant'] },
    },
    {
      path: '/stores',
      name: 'Stores',
      component: () => import('../pages/Stores.vue'),
    },
    {
      path: '/stores/:id',
      name: 'StoreDetail',
      component: () => import('../pages/StoreDetail.vue'),
    },
  ],
})

// Navigation guard
router.beforeEach((to, from, next) => {
  const userStore = useUserStore()

  if (to.meta.requiresAuth && !userStore.isAuthenticated) {
    next('/login')
  } else if (to.meta.role && userStore.user?.role !== to.meta.role) {
    next('/dashboard')
  } else if (to.meta.roles && !to.meta.roles.includes(userStore.user?.role)) {
    next('/dashboard')
  } else if ((to.path === '/login' || to.path === '/register') && userStore.isAuthenticated) {
    next('/dashboard')
  } else {
    next()
  }
})

export default router
