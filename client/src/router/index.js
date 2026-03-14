import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  { path: '/', name: 'Home', component: () => import('../views/HomeView.vue') },
  { path: '/catalog', name: 'Catalog', component: () => import('../views/CatalogView.vue') },
  { path: '/catalog/:id', name: 'Product', component: () => import('../views/ProductView.vue') },
  { path: '/contacts', name: 'Contacts', component: () => import('../views/ContactsView.vue') },
]

export default createRouter({ history: createWebHistory(), routes })
