export const MENULIST = [
    {
        menuName: "我的账号",
        menuIcon: "fz-ad-icon-test",
        menuSubLink: [
            {
                menuName: "密码修改",
                menuUrl: "/user_password"
            },
            {
                menuName: "添加用户",
                menuUrl: "/add_user"
            }
        ]
    },
    {
        menuName: "用户管理",
        menuIcon: "fz-ad-guanggao",
        menuUrl: "/user_list"
    },
    {
        menuName: "单车管理",
        menuIcon: "fz-ad-toufang",
        menuSubLink: [
            {
                menuName: "单车列表",
                menuUrl: "/bike_list"
            },
            {
                menuName: "告警信息",
                menuUrl: "/bike_alarm"
            }
        ]
    },
    {
        menuName: "数据统计",
        menuIcon: "fz-ad-statistics",
        menuSubLink: [
            {
                menuName: "单车数据统计",
                menuUrl: "/bike_stat"
            },
            {
                menuName: "用户数据统计",
                menuUrl: "/user_stat"
            }
        ]
    },
    {
        menuName: "退出登录",
        menuIcon: "fz-ad-guanli",
        menuUrl: "/logout"
    }
];