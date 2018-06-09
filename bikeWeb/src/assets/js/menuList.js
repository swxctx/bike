export const MENULIST = [
    {
        menuName: "立即使用",
        menuIcon: "fz-ad-rizhi",
        menuUrl: "/use"
    },
    {
        menuName: "我的",
        menuIcon: "fz-ad-icon-test",
        menuSubLink: [
            {
                menuName: "个人资料",
                menuUrl: "/profile"
            },
            {
                menuName: "修改密码",
                menuUrl: "/update_pass"
            },
            {
                menuName: "实名认证",
                menuUrl: "/auth"
            },
            {
                menuName: "使用记录",
                menuUrl: "/use_log"
            }
        ]
    },
    {
        menuName: "钱包",
        menuIcon: "fz-ad-guanggao",
        menuSubLink: [
            {
                menuName: "余额明细",
                menuUrl: "/detail"
            },
            {
                menuName: "充值",
                menuUrl: "/topup"
            }
        ]
    },
    {
        menuName: "反馈",
        menuIcon: "fz-ad-rizhi",
        menuUrl: "/feedback"
    },
    {
        menuName: "退出登录",
        menuIcon: "fz-ad-rizhi",
        menuUrl: "/logout"
    }
];