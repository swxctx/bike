export const MENULIST = [
    {
        menuName: "我的账号",
        menuIcon: "fz-ad-icon-test",
        menuSubLink: [
            {
                menuName: "资料修改",
                menuUrl: "/user_info"
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
        menuSubLink: [
            {
                menuName: "用户列表",
                menuUrl: "/"
            },
            {
                menuName: "用户封禁",
                menuUrl: "/ad/check"
            }
        ]
    },
    {
        menuName: "单车管理",
        menuIcon: "fz-ad-toufang",
        menuSubLink: [
            {
                menuName: "单车列表",
                menuUrl: "/puton/area"
            },
            {
                menuName: "告警信息",
                menuUrl: "/puton/regular"
            }
        ]
    },
    {
        menuName: "数据统计",
        menuIcon: "fz-ad-statistics",
        menuSubLink: [
            {
                menuName: "单车数据统计",
                menuUrl: "/status/adhost"
            },
            {
                menuName: "用户数据统计",
                menuUrl: "/status/channel"
            }
        ]
    },
    {
        menuName: "退出登录",
        menuIcon: "fz-ad-guanli",
        menuUrl: "/logout"
    }
];