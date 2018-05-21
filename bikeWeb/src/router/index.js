import Vue from 'vue'
import Router from 'vue-router'
/*引入页面*/
import Login from '@/components/login/login.vue'
import Home from '@/components/home/home.vue'
import BaiDuMap from "@/components/map/baidumap.vue";
import Menu from '@/components/menu/SideBar.vue'
import FeedBack from '@/components/feedback/feedback.vue'


Vue.use(Router)

/*配置路由*/
export default new Router({
  routes: [
      {
        path: "/",
        name: "首页",
        component: Home,
      },
      {
        path: "/menu",
        name: "导航",
        component: Menu,
        children: [
          {
            path: '/login',
            name: '登录',
            component: Login
          },
          {
            path: "/map",
            name: "定位",
            component: BaiDuMap,
          },
          {
            path: "/feedback",
            name: "定位",
            component: FeedBack,
          },
        ]
      }
  ]
})