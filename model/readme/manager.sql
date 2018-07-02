CREATE TABLE `admin` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(30) NOT NULL DEFAULT '' COMMENT '用户名',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '姓名',
  `password` varchar(255) NOT NULL DEFAULT '',
  `is_admin` int(11) NOT NULL DEFAULT '0' COMMENT '权限',
  `last_login_ts` bigint(20) unsigned NOT NULL COMMENT '最后登录时间',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='后台管理用户表';