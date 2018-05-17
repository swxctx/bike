import Vue from 'vue'
import Router from 'vue-router'
/*引入页面*/
import Login from '@/components/login/login.vue'
import Home from '@/components/home/home.vue'
import BaiDuMap from "@/components/map/baidumap.vue";
import Menu from '@/components/menu/SideBar.vue'


Vue.use(Router)

/*配置路由*/
export default new Router({
  routes: [
    {
      path: "/",
      name: "导航",
      component: Menu,
      children: [
        {
          path: '/login',
          name: '登录',
          component: Login
        },
        {
          path: '/home',
          name: '首页',
          component: Home
        },
        {
          path: '/map',
          name: '定位',
          component: BaiDuMap
        }
      ]
    }
  ]
})