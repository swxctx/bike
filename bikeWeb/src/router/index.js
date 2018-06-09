import Vue from 'vue'
import Router from 'vue-router'
/*引入页面*/
import Login from '@/components/account/login.vue'
import Logout from '@/components/account/logout.vue'
import Home from '@/components/home/home.vue'

import BaiDuMap from "@/components/map/baidumap.vue";
import Menu from '@/components/menu/SideBar.vue'

import FeedBack from '@/components/feedback/feedback.vue'

import Detail from '@/components/bag/detail.vue'
import Topup from '@/components/bag/topup.vue'

import User from '@/components/user/user.vue'
import Auth from '@/components/user/auth.vue'
import Profile from '@/components/user/profile.vue'
import UseLog from '@/components/user/useLog.vue'

import UseBike from '@/components/bike/bike.vue'

import StartAuth from '@/components/user/startAuth.vue'

import UpdateProfile from '@/components/user/update_profile.vue'
import UpdatePass from '@/components/account/updatePass.vue'



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
          },
          {
            path: "/feedback",
            name: "反馈",
            component: FeedBack,
          },
          {
            path: "/use",
            name: "开始使用",
            component: UseBike,
          },
          {
            path: "/logout",
            name: "退出登录",
            component: Logout,
          },
          {
            path: "/start_auth",
            name: "开始认证",
            component: StartAuth,
          },
          {
            path: "/update_profile",
            name: "资料修改",
            component: UpdateProfile,
          },
          {
            path: "/update_pass",
            name: "密码修改",
            component: UpdatePass,
          }
        ]
      }
  ]
})