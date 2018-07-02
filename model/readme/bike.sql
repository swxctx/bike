CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL DEFAULT '' COMMENT '用户昵称, 不可重复',
  `phone` char(11) NOT NULL DEFAULT '' COMMENT '手机号码, 不可重复',
  `password` varchar(32) NOT NULL DEFAULT '' COMMENT 'md5之后的密码字符串',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱地址',
  `status` TINYINT(4) NOT NULL DEFAULT '' COMMENT '用户状态',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=258 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

CREATE TABLE `user_meta` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(10) unsigned NOT NULL,
  `sex` char(1) NOT NULL DEFAULT '' COMMENT '性别, 男: 1, 女: 2',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `bio` varchar(500) NOT NULL DEFAULT '' COMMENT '描述',
  `card_name` varchar(48) NOT NULL DEFAULT '' COMMENT '真实姓名',
  `card_id` varchar(48) NOT NULL DEFAULT '' COMMENT '身份证号',
  `longi_tude` varchar(48) NOT NULL DEFAULT '' COMMENT '经度',
  `lati_tude` varchar(48) NOT NULL DEFAULT '' COMMENT '纬度',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=257 DEFAULT CHARSET=utf8mb4 COMMENT='用户元信息表';

CREATE TABLE `bike` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `by_id` int(11) unsigned NOT NULL DEFAULT '' COMMENT '单车编号',
  `longi_tude` varchar(48) NOT NULL DEFAULT '' COMMENT '经度',
  `lati_tude` varchar(48) NOT NULL DEFAULT '' COMMENT '纬度',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=258 DEFAULT CHARSET=utf8mb4 COMMENT='单车表';

CREATE TABLE `use_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(10) unsigned NOT NULL,
  `by_id` int(11) unsigned NOT NULL DEFAULT '' COMMENT '单车编号',
  `longi_tude` varchar(48) NOT NULL DEFAULT '' COMMENT '经度',
  `lati_tude` varchar(48) NOT NULL DEFAULT '' COMMENT '纬度',
  `start_ts` bigint(20) unsigned NOT NULL COMMENT '开始使用时间',
  `end_ts` bigint(20) unsigned NOT NULL COMMENT '结束使用时间',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=258 DEFAULT CHARSET=utf8mb4 COMMENT='用户使用记录表';

CREATE TABLE `user_bag` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(10) unsigned NOT NULL,
  `balance` int(11) unsigned NOT NULL DEFAULT '' COMMENT '余额',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=258 DEFAULT CHARSET=utf8mb4 COMMENT='用户钱包表';


CREATE TABLE `top_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(10) unsigned NOT NULL,
  `count` int(11) unsigned NOT NULL DEFAULT '' COMMENT '充值金额',
  `top_ts` bigint(20) unsigned NOT NULL COMMENT '充值时间',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=258 DEFAULT CHARSET=utf8mb4 COMMENT='用户充值记录';
