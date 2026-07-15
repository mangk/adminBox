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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

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

INSERT INTO `sys_auth` (`cb`, `ub`, `db`, `ct`, `ut`, `dt`, `table_id`, `table_module`, `type`, `key`, `set_value`) VALUES
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-106', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-710', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-701', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-100', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-504', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-301', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-306', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-501', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-200', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-303', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-505', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-707', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-602', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-709', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-705', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-304', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-704', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-102', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-200', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-500', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-304', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-603', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-100', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-300', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-302', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-103', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-404', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-202', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-502', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-503', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-706', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-203', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-402', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-401', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-703', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-204', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-600', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-400', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-700', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-104', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-301', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-601', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-105', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-604', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-305', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-702', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-711', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-403', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-300', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-708', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-302', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-101', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-405', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-201', 1),
(NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-303', 1);

INSERT INTO `sys_department` (`cb`, `ub`, `db`, `ct`, `ut`, `dt`, `pid`, `name`, `description`) VALUES
(NULL, NULL, NULL, NULL, NULL, NULL, 0, 'AdminBox', '');

INSERT INTO `sys_role` (`cb`, `ub`, `db`, `ct`, `ut`, `dt`, `name`, `description`) VALUES
(NULL, NULL, NULL, NULL, NULL, NULL, 'superAdmin', '超级管理员拥有所有权限');

INSERT INTO `sys_user` (`cb`, `ub`, `db`, `ct`, `ut`, `dt`, `uuid`, `username`, `phone`, `email`, `salt`, `password`, `nick_name`, `avatar`, `enable`, `user_config`, `extend_config`, `last_login_at`, `feishu_user_info`) VALUES
(NULL, NULL, NULL, NULL, NULL, NULL, '2efbd13a-d4a1-4af7-b5f5-2ec7d68d048c', 'adminbox', '12345678901', 'super@adminbox.demo', NULL, '$2a$10$VuSoUrm4xO1NWgZ/ENCUL.2ERAWh5dCxQj9fjQIqpneRz5ua0X31W', 'super', '', 1, NULL, NULL, NUll, NULL);

INSERT INTO `sys_user_department` (`sys_user_id`, `sys_department_id`) VALUES
(1, 1);

INSERT INTO `sys_user_role` (`sys_user_id`, `sys_role_id`) VALUES
(1, 1);