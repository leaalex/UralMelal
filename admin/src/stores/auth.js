import { defineStore } from 'pinia'
import client from '../api/client'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: JSON.parse(localStorage.getItem('user') || 'null'),
    token: localStorage.getItem('token'),
  }),
  getters: {
    isLoggedIn: (s) => !!s.token,
    isAdmin: (s) => s.user?.role === 'admin',
    isManager: (s) => ['admin', 'manager'].includes(s.user?.role),
    isEditor: (s) => ['admin', 'editor'].includes(s.user?.role),
  },
  actions: {
    async login(username, password) {
      const { data } = await client.post('/auth/login', { username, password })
      this.token = data.token
      this.user = data.user
      localStorage.setItem('token', data.token)
      localStorage.setItem('user', JSON.stringify(data.user))
    },
    logout() {
      this.token = null
      this.user = null
      localStorage.removeItem('token')
      localStorage.removeItem('user')
    },
    async fetchMe() {
      const { data } = await client.get('/auth/me')
      this.user = data
      localStorage.setItem('user', JSON.stringify(data))
    },
  },
})
