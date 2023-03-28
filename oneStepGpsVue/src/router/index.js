import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/pages/Login'
import Index from '@/pages/Index'
import Objects from '@/pages/Objects'
import UserSettings from '@/pages/UserSettings'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/login',
      component: Login
    },
    {
      path: '/',
      name: 'index',
      component: Index
    },
    {
      path: '/map/:id?',
      name: 'map',
      component: Index
    },
    {
      path: '/objects',
      component: Objects
    },
    {
      path: '/userSettings',
      component: UserSettings
    }
  ]
})
