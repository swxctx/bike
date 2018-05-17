import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/components/login/login.vue'
import Home from '@/components/home/home.vue'
import Menu from '@/components/menu/SideBar.vue'


Vue.use(Router)

export default new Router({
  routes: [
    {
      path: "/",
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
