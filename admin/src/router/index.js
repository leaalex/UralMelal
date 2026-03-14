import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  { path: '/login', name: 'Login', component: () => import('../views/LoginView.vue'), meta: { guest: true } },
  {
    path: '/',
    component: () => import('../layouts/AdminLayout.vue'),
    meta: { requiresAuth: true },
    children: [
      { path: '', name: 'Dashboard', component: () => import('../views/DashboardView.vue') },
      { path: 'catalog', name: 'Catalog', component: () => import('../views/CatalogView.vue') },
      { path: 'catalog/import', name: 'CatalogImport', component: () => import('../views/CatalogImportView.vue') },
      { path: 'categories', name: 'Categories', component: () => import('../views/CategoriesView.vue') },
      { path: 'content', name: 'Content', component: () => import('../views/ContentView.vue') },
      { path: 'leads', name: 'Leads', component: () => import('../views/LeadsView.vue') },
      { path: 'users', name: 'Users', component: () => import('../views/UsersView.vue'), meta: { adminOnly: true } },
    ],
  },
]

const router = createRouter({ history: createWebHistory(import.meta.env.BASE_URL), routes })

router.beforeEach((to, _from, next) => {
  const token = localStorage.getItem('token')
  if (to.meta.requiresAuth && !token) next({ name: 'Login' })
  else if (to.meta.guest && token) next({ name: 'Dashboard' })
  else if (to.meta.adminOnly && JSON.parse(localStorage.getItem('user') || '{}')?.role !== 'admin') next({ name: 'Dashboard' })
  else next()
})

export default router
