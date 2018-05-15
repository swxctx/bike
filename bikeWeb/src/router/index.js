import Vue from 'vue'
import Router from 'vue-router'
/*引入页面*/
import Login from '@/components/login/login.vue'
import Home from '@/components/home/home.vue'
import BaiDuMap from "@/components/map/baidumap.vue";

Vue.use(Router)

/*配置路由*/
export default new Router({
  routes: [
    {
      path: '/login',
      name: 'Login',
      component: Login
    },
    {
      path: '/home',
      name: 'Home',
      component: Home
    },
    {
      path: '/',
      name: 'Map',
      component: BaiDuMap
    }
  ]
})