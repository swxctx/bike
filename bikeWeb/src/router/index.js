import Vue from 'vue'
import Router from 'vue-router'
/*引入页面*/
import Login from '@/components/account/login.vue'
import Logout from '@/components/account/logout.vue'
import Home from '@/components/home/home.vue'

import BaiDuMap from "@/components/map/baidumap.vue";
import Menu from '@/components/menu/SideBar.vue'

import FeedBack from '@/components/feedback/feedback.vue'
import Help from '@/components/help/help.vue'

import Detail from '@/components/bag/detail.vue'
import Topup from '@/components/bag/topup.vue'

import User from '@/components/user/user.vue'
import Auth from '@/components/user/auth.vue'
import Profile from '@/components/user/auth.vue'
import UseLog from '@/components/user/useLog.vue'

import UseBike from '@/components/bike/bike.vue'

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
            path: "/detail",
            name: "明细",
            component: Detail,
          },
          {
            path: "/topup",
            name: "充值",
            component: Topup,
          },
          {
            path: "/user",
            name: "用户",
            component: User,
            children: [
              {
                path: "/auth",
                name: "认证",
                component: Auth,
              },
              {
                path: "/profile",
                name: "资料",
                component: Profile,
              },
              {
                path: "/use_log",
                name: "使用记录",
                component: UseLog,
              }
            ]
          },
          {
            path: "/feedback",
            name: "反馈",
            component: FeedBack,
          },
          {
            path: "/help",
            name: "帮助",
            component: Help,
          },
          {
            path: "/use",
            name: "开始使用",
            component: UseBike,
          },
        ]
      }
  ]
})