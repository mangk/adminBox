/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


DROP TABLE IF EXISTS `sys_api`;
CREATE TABLE `sys_api` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `cb` int DEFAULT NULL COMMENT '创建者',
  `ub` int DEFAULT NULL COMMENT '更新者',
  `db` int DEFAULT NULL COMMENT '删除者',
  `ct` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `dt` datetime DEFAULT NULL COMMENT '删除时间',
  `menu_id` int DEFAULT NULL COMMENT '菜单Id',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'api名称',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'api简介',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'api路径',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT 'POST' COMMENT '方法:创建POST(默认)|查看GET|更新PUT|删除DELETE',
  PRIMARY KEY (`id`),
  KEY `idx_sys_api_dt` (`dt`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `sys_auth`;
CREATE TABLE `sys_auth` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `cb` int DEFAULT NULL COMMENT '创建者',
  `ub` int DEFAULT NULL COMMENT '更新者',
  `db` int DEFAULT NULL COMMENT '删除者',
  `ct` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `dt` datetime DEFAULT NULL COMMENT '删除时间',
  `table_id` int DEFAULT NULL COMMENT '表ID',
  `table_module` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '表模块',
  `type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '类型: menu,api,action',
  `key` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '数据标记: menu: id, api: id, action: mapKey',
  `set_value` int DEFAULT NULL COMMENT '设置值',
  PRIMARY KEY (`id`),
  KEY `idx_sys_auth_dt` (`dt`)
) ENGINE=InnoDB AUTO_INCREMENT=58 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `sys_casbin_rule`;
CREATE TABLE `sys_casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_sys_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `sys_department`;
CREATE TABLE `sys_department` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `cb` int DEFAULT NULL COMMENT '创建者',
  `ub` int DEFAULT NULL COMMENT '更新者',
  `db` int DEFAULT NULL COMMENT '删除者',
  `ct` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `dt` datetime DEFAULT NULL COMMENT '删除时间',
  `pid` int DEFAULT NULL COMMENT '父级Id',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '角色名称',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '角色简介',
  PRIMARY KEY (`id`),
  KEY `idx_sys_department_dt` (`dt`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `sys_file`;
CREATE TABLE `sys_file` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `cb` int DEFAULT NULL COMMENT '创建者',
  `ub` int DEFAULT NULL COMMENT '更新者',
  `db` int DEFAULT NULL COMMENT '删除者',
  `ct` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `dt` datetime DEFAULT NULL COMMENT '删除时间',
  `group_id` int DEFAULT NULL COMMENT '分组id',
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '文件名',
  `url` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '文件地址',
  `tag` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '文件标签',
  `key` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '编号',
  `uuid` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `size` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_sys_file_upload_dt` (`dt`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `sys_file_group`;
CREATE TABLE `sys_file_group` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `cb` int DEFAULT NULL COMMENT '创建者',
  `ub` int DEFAULT NULL COMMENT '更新者',
  `db` int DEFAULT NULL COMMENT '删除者',
  `ct` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `dt` datetime DEFAULT NULL COMMENT '删除时间',
  `parent_id` bigint DEFAULT NULL COMMENT '父级id',
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '路径',
  `desc` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '描述',
  PRIMARY KEY (`id`),
  KEY `idx_zomb_sys_file_path_dt` (`dt`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `cb` int DEFAULT NULL COMMENT '创建者',
  `ub` int DEFAULT NULL COMMENT '更新者',
  `db` int DEFAULT NULL COMMENT '删除者',
  `ct` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `dt` datetime DEFAULT NULL COMMENT '删除时间',
  `pid` int DEFAULT NULL COMMENT '父级Id',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '菜单名称',
  `path` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '路由地址',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否隐藏',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '模版地址',
  `sort` int DEFAULT NULL COMMENT '排序',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '附加属性',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '附加属性',
  `auto_close` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `sc_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '附加属性',
  `action_list` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '附加属性',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_sys_menu_name` (`name`),
  KEY `idx_sys_menu_dt` (`dt`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `cb` int DEFAULT NULL COMMENT '创建者',
  `ub` int DEFAULT NULL COMMENT '更新者',
  `db` int DEFAULT NULL COMMENT '删除者',
  `ct` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `dt` datetime DEFAULT NULL COMMENT '删除时间',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '角色名称',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '角色简介',
  PRIMARY KEY (`id`),
  KEY `idx_sys_role_dt` (`dt`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `cb` int DEFAULT NULL COMMENT '创建者',
  `ub` int DEFAULT NULL COMMENT '更新者',
  `db` int DEFAULT NULL COMMENT '删除者',
  `ct` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `dt` datetime DEFAULT NULL COMMENT '删除时间',
  `uuid` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户登录名',
  `phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户手机号',
  `email` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户邮箱',
  `salt` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '密码混淆',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '系统用户' COMMENT '用户昵称',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户头像',
  `enable` tinyint(1) DEFAULT '1' COMMENT '用户是否有效',
  `user_config` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '用户配置文件',
  `extend_config` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '扩展配置，保存自定义使用的配置',
  `last_login_at` datetime DEFAULT NULL COMMENT '最后登录时间',
  `feishu_user_info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`),
  KEY `idx_sys_user_uuid` (`uuid`),
  KEY `idx_sys_user_username` (`username`),
  KEY `idx_sys_user_dt` (`dt`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `sys_user_department`;
CREATE TABLE `sys_user_department` (
  `sys_user_id` int DEFAULT NULL,
  `sys_department_id` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role` (
  `sys_user_id` int DEFAULT NULL,
  `sys_role_id` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `sys_api` (`id`, `cb`, `ub`, `db`, `ct`, `ut`, `dt`, `menu_id`, `name`, `description`, `path`, `method`) VALUES
(1, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'example', '', '/example', 'POST');

INSERT INTO `sys_auth` (`id`, `cb`, `ub`, `db`, `ct`, `ut`, `dt`, `table_id`, `table_module`, `type`, `key`, `set_value`) VALUES
(1, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-100', 1),
(2, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-100', 1),
(3, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-101', 1),
(4, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-102', 1),
(5, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-103', 1),
(6, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-104', 1),
(7, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-105', 1),
(8, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-106', 1),
(9, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-200', 1),
(10, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-200', 1),
(11, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-201', 1),
(12, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-202', 1),
(13, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-203', 1),
(14, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-204', 1),
(15, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-300', 1),
(16, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-300', 1),
(17, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-301', 1),
(18, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-301', 1),
(19, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-302', 1),
(20, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-302', 1),
(21, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-303', 1),
(22, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-303', 1),
(23, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-304', 1),
(24, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-304', 1),
(25, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-305', 1),
(26, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-306', 1),
(27, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-400', 1),
(28, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-401', 1),
(29, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-402', 1),
(30, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-403', 1),
(31, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-404', 1),
(32, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-405', 1),
(33, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-500', 1),
(34, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-501', 1),
(35, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-502', 1),
(36, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-503', 1),
(37, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-504', 1),
(38, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-505', 1),
(39, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-600', 1),
(40, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-601', 1),
(41, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-602', 1),
(42, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-603', 1),
(43, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-604', 1),
(44, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-700', 1),
(45, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-701', 1),
(46, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-702', 1),
(47, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-703', 1),
(48, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-704', 1),
(49, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-705', 1),
(50, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-706', 1),
(51, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-707', 1),
(52, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-708', 1),
(53, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-709', 1),
(54, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-710', 1),
(55, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-711', 1),
(56, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '1', 1),
(57, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '1', 1);

INSERT INTO `sys_department` (`id`, `cb`, `ub`, `db`, `ct`, `ut`, `dt`, `pid`, `name`, `description`) VALUES
(1, NULL, NULL, NULL, NULL, NULL, NULL, 0, 'adminBox', '');

INSERT INTO `sys_menu` (`id`, `cb`, `ub`, `db`, `ct`, `ut`, `dt`, `pid`, `name`, `path`, `hidden`, `component`, `sort`, `title`, `keep_alive`, `default_menu`, `icon`, `auto_close`, `sc_path`, `action_list`) VALUES
(1, NULL, NULL, NULL, NULL, NULL, NULL, 0, 'example', 'example', 0, 'views/util/serverComponent.vue', 0, 'Example Vue Template', 0, 0, 'Apple', 0, '/example', NULL);

INSERT INTO `sys_role` (`id`, `cb`, `ub`, `db`, `ct`, `ut`, `dt`, `name`, `description`) VALUES
(1, NULL, NULL, NULL, NULL, NULL, NULL, 'adminBox', '超级管理员拥有所有权限');

INSERT INTO `sys_user` (`id`, `cb`, `ub`, `db`, `ct`, `ut`, `dt`, `uuid`, `username`, `phone`, `email`, `salt`, `password`, `nick_name`, `avatar`, `enable`, `user_config`, `extend_config`, `last_login_at`, `feishu_user_info`) VALUES
(1, NULL, NULL, NULL, NULL, NULL, NULL, 'f9209ecc-7b8e-4357-b2f4-389b6a3a1f0a', 'adminbox', '12345678901', 'adminbox@adminbox.com', NULL, '$2a$10$Sm8wwe/It2Y4J2xx3o1uku1wDR95b58fmJbO5q.FPPmZBjaBuoxTG', 'adminbox', '', 1, NULL, NULL, NULL, NULL);

INSERT INTO `sys_user_department` (`sys_user_id`, `sys_department_id`) VALUES
(1, 1);

INSERT INTO `sys_user_role` (`sys_user_id`, `sys_role_id`) VALUES
(1, 1);



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;