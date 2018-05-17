import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/components/login/login.vue'
import Welcome from '@/components/welcome/welcome.vue'
import Menu from '@/components/menu/SideBar.vue'
import Home from '@/components/home/home.vue'



Vue.use(Router)

export default new Router({
  routes: [
    {
      path: "/",
      name: "登录页",
      component: Welcome
    },
    {
      path: "/menu",
      name: "导航",
      component: Menu,
      children:[
        {
          path: '/login',
          name: '登录',
          component: Login
        },
        {
          path: '/home',
          name: '首页',
          component: Home
        }
      ]
    }
  ]
})
