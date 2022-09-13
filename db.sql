DROP DATABASE IF EXISTS `uima`;

CREATE DATABASE `uima`;

USE `uima`;

-- 用户
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`(
    `id`        INT NOT NULL AUTO_INCREMENT,
    `phone`     VARCHAR(11) NOT NULL UNIQUE COMMENT "电话号码",
    `nickname`  VARCHAR(255) NULL COMMENT "用户名",
    `password`  VARCHAR(255) NOT NULL COMMENT "密码",
    `avatar`    VARCHAR(255) NULL COMMENT "头像",
    `gender`    VARCHAR(255) NULL COMMENT "性别",
    `email`     VARCHAR(255) NULL COMMENT "邮箱",
    `realname`  VARCHAR(255) NULL COMMENT "真实姓名",
    `idcard`    VARCHAR(255) NULL COMMENT "身份证号",
    `sha`       VARCHAR(255) NULL,  -- 用于设置头像
    `path`      VARCHAR(255) NULL,   -- 用于设置头像
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 商店
DROP TABLE IF EXISTS `shops`;
CREATE TABLE `shops`(
    `id`            INT NOT NULL AUTO_INCREMENT,
    `picture`       VARCHAR(255) NULL COMMENT "商店图片",
    `shop_name`     VARCHAR(255) NOT NULL COMMENT "商店名称",
    `fiery_num`     VARCHAR(255) NOT NULL  COMMENT "火热指数",
    `opening_time`  VARCHAR(255) NULL COMMENT "营业时间",
    `current_num`   VARCHAR(255) NULL,
    `hot_line`      VARCHAR(255) NULL,
    `sha`           VARCHAR(255) NULL,
    `path`          VARCHAR(255) NULL,
    `service_intro` TEXT NULL COMMENT "服务介绍",
    `vip_service`   TEXT NULL COMMENT "vip服务",
    `should_know`   TEXT NULL COMMENT "需知",
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 精彩放送
DROP TABLE IF EXISTS `broadcasts`;
CREATE TABLE `broadcasts`(
    `id`            INT NOT NULL AUTO_INCREMENT,
    `content`       TEXT NULL,
    `title`         VARCHAR(255) NULL,
    `create_time`   VARCHAR(255) NULL,
    `picture`       VARCHAR(255) NULL,
    `sha`           VARCHAR(255) NULL,  -- 用于设置图片
    `path`          VARCHAR(255) NULL,   -- 用于设置图片
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `scenic_spots`;
CREATE TABLE `scenic_spots`(
    `id`            INT NOT NULL AUTO_INCREMENT,
    `name`          VARCHAR(255) NOT NULL,
    `picture`       VARCHAR(255) NOT NULL,
    `download_time` VARCHAR(255) NOT NULL,
    `sha`           VARCHAR(255) NULL,  -- 用于设置图片
    `path`          VARCHAR(255) NULL,   -- 用于设置图片
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 订单
DROP TABLE IF EXISTS `script_orders`;
CREATE TABLE `script_orders`(
    `id`            INT NOT NULL AUTO_INCREMENT,
    `user_id`       INT NOT NULL COMMENT "用户id",
    `script_id`     INT NOT NULL COMMENT "剧本id",
    `script_name`   VARCHAR(255) NOT NULL COMMENT "剧本名称",
    `type`          VARCHAR(255) NOT NULL COMMENT "类型(店铺/线上/线下)/剧本",
    `price`         INT NOT NULL COMMENT "价格",
    `createtime`    VARCHAR(255) NOT NULL COMMENT "创建订单时间",
    `paymenttime`   VARCHAR(255) NOT NULL COMMENT "付款时间",
    `avatar`        VARCHAR(255) NULL COMMENT "订单图片",
    `information`   VARCHAR(255) NOT NULL COMMENT "订单内容",
    `status`        VARCHAR(255) NOT NULL COMMENT "订单状态",
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- 剧本
DROP TABLE IF EXISTS `scripts`;
CREATE TABLE `scripts`(
    `id`            INT NOT NULL  AUTO_INCREMENT,
    `time`          VARCHAR(255)  NOT NULL COMMENT "剧本时长",
    `script_name`   VARCHAR(255)  NOT NULL COMMENT "剧本名称",
    `place`         VARCHAR(255)  NOT NULL COMMENT "地点",
    `brief_intro`   VARCHAR(255)  NOT NULL COMMENT "剧本简介",
    `introduction`  VARCHAR(1000) NOT NULL COMMENT "剧本介绍",
    `avatar`        VARCHAR(255)  COMMENT "剧本封面",
    `price`         INT NOT NULL  COMMENT "价格",
    `tag1`          VARCHAR(100)  COMMENT "标签一",
    `tag2`          VARCHAR(100)  COMMENT "标签二",
    `tag3`          VARCHAR(100)  COMMENT "标签三",
    `tag4`          VARCHAR(100)  COMMENT "标签四",
    `tag5`          VARCHAR(100)  COMMENT "标签五",
    `sha`           VARCHAR(255)  NULL,  -- 用于设置图片
    `path`          VARCHAR(255)  NULL,  -- 用于设置图片
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 地点详情
DROP TABLE IF EXISTS `places`;
CREATE TABLE `places`(
    `id`            INT NOT NULL  AUTO_INCREMENT,
    `name`          VARCHAR(255)  NOT NULL UNIQUE COMMENT "地点名称",
    `data`          VARCHAR(1000) NOT NULL COMMENT "地点资料",
    `area`          VARCHAR(255)  NOT NULL COMMENT "地点面积",
    `visitor`       VARCHAR(255)  NOT NULL COMMENT "最大游客量",
    `entertainment` VARCHAR(255)  NOT NULL COMMENT "娱乐项目数量",
    `scenic_spot`   VARCHAR(255)  NOT NULL COMMENT "特色景点数量",
    `picture1`      VARCHAR(255)  COMMENT "地点轮换封面1",
    `sha1`          VARCHAR(255)  NULL,  -- 用于设置图片
    `path1`         VARCHAR(255)  NULL,  -- 用于设置图片
    `picture2`      VARCHAR(255)  COMMENT "地点轮换封面2",
    `sha2`          VARCHAR(255)  NULL,  -- 用于设置图片
    `path2`         VARCHAR(255)  NULL,  -- 用于设置图片
    `picture3`      VARCHAR(255)  COMMENT "地点轮换封面3",
    `sha3`          VARCHAR(255)  NULL,  -- 用于设置图片
    `path3`         VARCHAR(255)  NULL,  -- 用于设置图片
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 剧本收藏
DROP TABLE IF EXISTS `script_collections`;
CREATE TABLE `script_collections`(
    `id`          INT NOT NULL AUTO_INCREMENT,
    `users_id`    INT NOT NULL COMMENT "用户id",
    `scripts_id`  INT NOT NULL COMMENT "剧本id",
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 剧本预约
DROP TABLE IF EXISTS `script_appointments`;
CREATE TABLE `script_appointments`(
    `id`            INT NOT NULL AUTO_INCREMENT,
    `users_id`      INT NOT NULL COMMENT "用户id",
    `scripts_id`    INT NOT NULL COMMENT "剧本id",
    `scripts_name`  VARCHAR(255) NOT NULL COMMENT "剧本名称",
    `scripts_cover` VARCHAR(255) NOT NULL COMMENT "剧本封面",
    `time`          VARCHAR(255) COMMENT "预约时间",
    `status`        VARCHAR(255) NOT NULL COMMENT "预约状态(预约中/已预约/已完成)",
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;