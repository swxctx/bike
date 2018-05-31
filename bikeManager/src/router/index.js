import Vue from 'vue'
import Router from 'vue-router'
import Welcome from '@/components/welcome/welcome.vue'
import Menu from '@/components/menu/SideBar.vue'
import Home from '@/components/home/home.vue'
import Logout from '@/components/account/logout.vue'
import AddUser from '@/components/account/addUser.vue'
import Login from '@/components/account/login.vue'
import UserList from '@/components/user/user.vue'
import BikeList from '@/components/bike/bike.vue'
import BikeAlarm from '@/components/bike/alarm.vue'


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
        },
        {
          path: '/logout',
          name: '退出登录',
          component: Logout
        },
        {
          path: '/add_user',
          name: '添加用户',
          component: AddUser
        },
        {
          path: '/user_list',
          name: '用户列表',
          component: UserList
        },
        {
          path: '/bike_list',
          name: '单车列表',
          component: BikeList
        },
        {
          path: '/bike_alarm',
          name: '告警信息',
          component: BikeAlarm
        }
      ]
    }
  ]
})
