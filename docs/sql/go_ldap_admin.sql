/*
 Navicat Premium Data Transfer

 Source Server         : test-eryajf
 Source Server Type    : MySQL
 Source Server Version : 50741
 Source Host           : localhost:3306
 Source Schema         : go_ldap_admin

 Target Server Type    : MySQL
 Target Server Version : 50741
 File Encoding         : 65001

 Date: 24/02/2023 22:29:50
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for apis
-- ----------------------------
DROP TABLE IF EXISTS `apis`;
CREATE TABLE `apis` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `method` varchar(20) DEFAULT NULL COMMENT '''请求方式''',
  `path` varchar(100) DEFAULT NULL COMMENT '''访问路径''',
  `category` varchar(50) DEFAULT NULL COMMENT '''所属类别''',
  `remark` varchar(100) DEFAULT NULL COMMENT '''备注''',
  `creator` varchar(20) DEFAULT NULL COMMENT '''创建人''',
  PRIMARY KEY (`id`),
  KEY `idx_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=55 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of apis
-- ----------------------------
BEGIN;
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (1, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/base/login', 'base', '用户登录', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (2, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/base/logout', 'base', '用户登出', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (3, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/base/refreshToken', 'base', '刷新JWT令牌', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (4, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/base/changePwd', 'base', '通过邮箱修改密码', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (5, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'GET', '/user/info', 'user', '获取当前登录用户信息', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (6, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'GET', '/user/list', 'user', '获取用户列表', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (7, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/user/changePwd', 'user', '更新用户登录密码', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (8, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/user/add', 'user', '创建用户', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (9, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/user/update', 'user', '更新用户', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (10, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/user/delete', 'user', '批量删除用户', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (11, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/user/changeUserStatus', 'user', '更改用户在职状态', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (12, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/user/syncDingTalkUsers', 'user', '从钉钉拉取用户信息', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (13, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/user/syncWeComUsers', 'user', '从企业微信拉取用户信息', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (14, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/user/syncFeiShuUsers', 'user', '从飞书拉取用户信息', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (15, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/user/syncOpenLdapUsers', 'user', '从openldap拉取用户信息', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (16, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/user/syncSqlUsers', 'user', '将数据库中的用户同步到Ldap', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (17, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'GET', '/group/list', 'group', '获取分组列表', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (18, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'GET', '/group/tree', 'group', '获取分组列表树', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (19, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/group/add', 'group', '创建分组', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (20, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/group/update', 'group', '更新分组', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (21, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/group/delete', 'group', '批量删除分组', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (22, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/group/adduser', 'group', '添加用户到分组', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (23, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/group/removeuser', 'group', '将用户从分组移出', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (24, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'GET', '/group/useringroup', 'group', '获取在分组内的用户列表', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (25, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'GET', '/group/usernoingroup', 'group', '获取不在分组内的用户列表', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (26, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/group/syncDingTalkDepts', 'group', '从钉钉拉取部门信息', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (27, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/group/syncWeComDepts', 'group', '从企业微信拉取部门信息', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (28, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/group/syncFeiShuDepts', 'group', '从飞书拉取部门信息', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (29, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/group/syncOpenLdapDepts', 'group', '从openldap拉取部门信息', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (30, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/group/syncSqlGroups', 'group', '将数据库中的分组同步到Ldap', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (31, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'GET', '/role/list', 'role', '获取角色列表', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (32, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/role/add', 'role', '创建角色', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (33, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/role/update', 'role', '更新角色', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (34, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'GET', '/role/getmenulist', 'role', '获取角色的权限菜单', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (35, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/role/updatemenus', 'role', '更新角色的权限菜单', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (36, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'GET', '/role/getapilist', 'role', '获取角色的权限接口', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (37, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/role/updateapis', 'role', '更新角色的权限接口', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (38, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/role/delete', 'role', '批量删除角色', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (39, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'GET', '/menu/list', 'menu', '获取菜单列表', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (40, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'GET', '/menu/tree', 'menu', '获取菜单树', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (41, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/menu/add', 'menu', '创建菜单', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (42, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/menu/update', 'menu', '更新菜单', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (43, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/menu/delete', 'menu', '批量删除菜单', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (44, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'GET', '/api/list', 'api', '获取接口列表', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (45, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'GET', '/api/tree', 'api', '获取接口树', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (46, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/api/add', 'api', '创建接口', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (47, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/api/update', 'api', '更新接口', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (48, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/api/delete', 'api', '批量删除接口', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (49, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'GET', '/fieldrelation/list', 'fieldrelation', '获取字段动态关系列表', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (50, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/fieldrelation/add', 'fieldrelation', '创建字段动态关系', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (51, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/fieldrelation/update', 'fieldrelation', '更新字段动态关系', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (52, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/fieldrelation/delete', 'fieldrelation', '批量删除字段动态关系', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (53, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'GET', '/log/operation/list', 'log', '获取操作日志列表', '系统');
INSERT INTO `apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `method`, `path`, `category`, `remark`, `creator`) VALUES (54, '2023-02-24 19:51:37.398', '2023-02-24 19:51:37.398', NULL, 'POST', '/log/operation/delete', 'log', '批量删除操作日志', '系统');
COMMIT;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_index` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=74 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (64, 'p', 'admin', '/api/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (66, 'p', 'admin', '/api/delete', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (60, 'p', 'admin', '/api/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (62, 'p', 'admin', '/api/tree', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (65, 'p', 'admin', '/api/update', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (7, 'p', 'admin', '/base/changePwd', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (1, 'p', 'admin', '/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (3, 'p', 'admin', '/base/logout', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (5, 'p', 'admin', '/base/refreshToken', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (68, 'p', 'admin', '/fieldrelation/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (70, 'p', 'admin', '/fieldrelation/delete', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (67, 'p', 'admin', '/fieldrelation/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (69, 'p', 'admin', '/fieldrelation/update', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (28, 'p', 'admin', '/group/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (31, 'p', 'admin', '/group/adduser', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (30, 'p', 'admin', '/group/delete', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (24, 'p', 'admin', '/group/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (32, 'p', 'admin', '/group/removeuser', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (37, 'p', 'admin', '/group/syncDingTalkDepts', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (39, 'p', 'admin', '/group/syncFeiShuDepts', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (40, 'p', 'admin', '/group/syncOpenLdapDepts', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (41, 'p', 'admin', '/group/syncSqlGroups', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (38, 'p', 'admin', '/group/syncWeComDepts', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (26, 'p', 'admin', '/group/tree', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (29, 'p', 'admin', '/group/update', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (33, 'p', 'admin', '/group/useringroup', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (35, 'p', 'admin', '/group/usernoingroup', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (73, 'p', 'admin', '/log/operation/delete', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (71, 'p', 'admin', '/log/operation/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (57, 'p', 'admin', '/menu/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (59, 'p', 'admin', '/menu/delete', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (53, 'p', 'admin', '/menu/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (55, 'p', 'admin', '/menu/tree', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (58, 'p', 'admin', '/menu/update', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (44, 'p', 'admin', '/role/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (52, 'p', 'admin', '/role/delete', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (49, 'p', 'admin', '/role/getapilist', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (46, 'p', 'admin', '/role/getmenulist', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (42, 'p', 'admin', '/role/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (45, 'p', 'admin', '/role/update', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (51, 'p', 'admin', '/role/updateapis', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (48, 'p', 'admin', '/role/updatemenus', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (15, 'p', 'admin', '/user/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (13, 'p', 'admin', '/user/changePwd', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (18, 'p', 'admin', '/user/changeUserStatus', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (17, 'p', 'admin', '/user/delete', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (9, 'p', 'admin', '/user/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (11, 'p', 'admin', '/user/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (19, 'p', 'admin', '/user/syncDingTalkUsers', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (21, 'p', 'admin', '/user/syncFeiShuUsers', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (22, 'p', 'admin', '/user/syncOpenLdapUsers', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (23, 'p', 'admin', '/user/syncSqlUsers', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (20, 'p', 'admin', '/user/syncWeComUsers', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (16, 'p', 'admin', '/user/update', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (61, 'p', 'user', '/api/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (63, 'p', 'user', '/api/tree', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (8, 'p', 'user', '/base/changePwd', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (2, 'p', 'user', '/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (4, 'p', 'user', '/base/logout', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (6, 'p', 'user', '/base/refreshToken', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (25, 'p', 'user', '/group/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (27, 'p', 'user', '/group/tree', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (34, 'p', 'user', '/group/useringroup', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (36, 'p', 'user', '/group/usernoingroup', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (72, 'p', 'user', '/log/operation/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (54, 'p', 'user', '/menu/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (56, 'p', 'user', '/menu/tree', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (50, 'p', 'user', '/role/getapilist', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (47, 'p', 'user', '/role/getmenulist', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (43, 'p', 'user', '/role/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (14, 'p', 'user', '/user/changePwd', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (10, 'p', 'user', '/user/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (12, 'p', 'user', '/user/list', 'GET', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for field_relations
-- ----------------------------
DROP TABLE IF EXISTS `field_relations`;
CREATE TABLE `field_relations` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `flag` longtext,
  `attributes` json DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_field_relations_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of field_relations
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for group_users
-- ----------------------------
DROP TABLE IF EXISTS `group_users`;
CREATE TABLE `group_users` (
  `group_id` bigint(20) unsigned NOT NULL,
  `user_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`group_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of group_users
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for groups
-- ----------------------------
DROP TABLE IF EXISTS `groups`;
CREATE TABLE `groups` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `group_name` varchar(128) DEFAULT NULL COMMENT '''分组名称''',
  `remark` varchar(128) DEFAULT NULL COMMENT '''分组中文说明''',
  `creator` varchar(20) DEFAULT NULL COMMENT '''创建人''',
  `group_type` varchar(20) DEFAULT NULL COMMENT '''分组类型：cn、ou''',
  `parent_id` bigint(20) unsigned DEFAULT '0' COMMENT '''父组编号(编号为0时表示根组)''',
  `source_dept_id` varchar(100) DEFAULT NULL COMMENT '''部门编号''',
  `source` varchar(20) DEFAULT NULL COMMENT '''来源：dingTalk、weCom、ldap、platform''',
  `source_dept_parent_id` varchar(100) DEFAULT NULL COMMENT '''父部门编号''',
  `source_user_num` bigint(20) DEFAULT '0' COMMENT '''部门下的用户数量，从第三方获取的数据''',
  `group_dn` varchar(255) NOT NULL COMMENT '''分组dn''',
  `sync_state` tinyint(1) DEFAULT '1' COMMENT '''同步状态:1已同步, 2未同步''',
  PRIMARY KEY (`id`),
  KEY `idx_groups_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of groups
-- ----------------------------
BEGIN;
INSERT INTO `groups` (`id`, `created_at`, `updated_at`, `deleted_at`, `group_name`, `remark`, `creator`, `group_type`, `parent_id`, `source_dept_id`, `source`, `source_dept_parent_id`, `source_user_num`, `group_dn`, `sync_state`) VALUES (1, '2023-02-24 19:51:37.453', '2023-02-24 19:51:37.453', NULL, 'root', 'Base', 'system', '', 0, '0', 'openldap', '0', 0, 'dc=eryajf,dc=net', 1);
INSERT INTO `groups` (`id`, `created_at`, `updated_at`, `deleted_at`, `group_name`, `remark`, `creator`, `group_type`, `parent_id`, `source_dept_id`, `source`, `source_dept_parent_id`, `source_user_num`, `group_dn`, `sync_state`) VALUES (2, '2023-02-24 19:51:37.453', '2023-02-24 19:51:37.453', NULL, 'dingtalkroot', '钉钉根部门', 'system', 'ou', 1, 'dingtalk_1', 'dingtalk', 'dingtalk_0', 0, 'ou=dingtalkroot,dc=eryajf,dc=net', 1);
INSERT INTO `groups` (`id`, `created_at`, `updated_at`, `deleted_at`, `group_name`, `remark`, `creator`, `group_type`, `parent_id`, `source_dept_id`, `source`, `source_dept_parent_id`, `source_user_num`, `group_dn`, `sync_state`) VALUES (3, '2023-02-24 19:51:37.453', '2023-02-24 19:51:37.453', NULL, 'wecomroot', '企业微信根部门', 'system', 'ou', 1, 'wecom_1', 'wecom', 'wecom_0', 0, 'ou=wecomroot,dc=eryajf,dc=net', 1);
INSERT INTO `groups` (`id`, `created_at`, `updated_at`, `deleted_at`, `group_name`, `remark`, `creator`, `group_type`, `parent_id`, `source_dept_id`, `source`, `source_dept_parent_id`, `source_user_num`, `group_dn`, `sync_state`) VALUES (4, '2023-02-24 19:51:37.453', '2023-02-24 19:51:37.453', NULL, 'feishuroot', '飞书根部门', 'system', 'ou', 1, 'feishu_0', 'feishu', 'feishu_0', 0, 'ou=feishuroot,dc=eryajf,dc=net', 1);
COMMIT;

-- ----------------------------
-- Table structure for menus
-- ----------------------------
DROP TABLE IF EXISTS `menus`;
CREATE TABLE `menus` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL COMMENT '''菜单名称(英文名, 可用于国际化)''',
  `title` varchar(50) DEFAULT NULL COMMENT '''菜单标题(无法国际化时使用)''',
  `icon` varchar(50) DEFAULT NULL COMMENT '''菜单图标''',
  `path` varchar(100) DEFAULT NULL COMMENT '''菜单访问路径''',
  `redirect` varchar(100) DEFAULT NULL COMMENT '''重定向路径''',
  `component` varchar(100) DEFAULT NULL COMMENT '''前端组件路径''',
  `sort` int(3) DEFAULT '999' COMMENT '''菜单顺序(1-999)''',
  `status` tinyint(1) DEFAULT '1' COMMENT '''菜单状态(正常/禁用, 默认正常)''',
  `hidden` tinyint(1) DEFAULT '2' COMMENT '''菜单在侧边栏隐藏(1隐藏，2显示)''',
  `no_cache` tinyint(1) DEFAULT '2' COMMENT '''菜单是否被 <keep-alive> 缓存(1不缓存，2缓存)''',
  `always_show` tinyint(1) DEFAULT '2' COMMENT '''忽略之前定义的规则，一直显示根路由(1忽略，2不忽略)''',
  `breadcrumb` tinyint(1) DEFAULT '1' COMMENT '''面包屑可见性(可见/隐藏, 默认可见)''',
  `active_menu` varchar(100) DEFAULT NULL COMMENT '''在其它路由时，想在侧边栏高亮的路由''',
  `parent_id` bigint(20) unsigned DEFAULT '0' COMMENT '''父菜单编号(编号为0时表示根菜单)''',
  `creator` varchar(20) DEFAULT NULL COMMENT '''创建人''',
  PRIMARY KEY (`id`),
  KEY `idx_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of menus
-- ----------------------------
BEGIN;
INSERT INTO `menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `title`, `icon`, `path`, `redirect`, `component`, `sort`, `status`, `hidden`, `no_cache`, `always_show`, `breadcrumb`, `active_menu`, `parent_id`, `creator`) VALUES (1, '2023-02-24 19:51:37.316', '2023-02-24 19:51:37.316', NULL, 'UserManage', '人员管理', 'user', '/personnel', '/personnel/user', 'Layout', 5, 1, 2, 2, 2, 1, '', 0, '系统');
INSERT INTO `menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `title`, `icon`, `path`, `redirect`, `component`, `sort`, `status`, `hidden`, `no_cache`, `always_show`, `breadcrumb`, `active_menu`, `parent_id`, `creator`) VALUES (2, '2023-02-24 19:51:37.316', '2023-02-24 19:51:37.316', NULL, 'User', '用户管理', 'people', 'user', '', '/personnel/user/index', 6, 1, 2, 2, 2, 1, '', 1, '系统');
INSERT INTO `menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `title`, `icon`, `path`, `redirect`, `component`, `sort`, `status`, `hidden`, `no_cache`, `always_show`, `breadcrumb`, `active_menu`, `parent_id`, `creator`) VALUES (3, '2023-02-24 19:51:37.316', '2023-02-24 19:51:37.316', NULL, 'Group', '分组管理', 'peoples', 'group', '', '/personnel/group/index', 7, 1, 2, 1, 2, 1, '', 1, '系统');
INSERT INTO `menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `title`, `icon`, `path`, `redirect`, `component`, `sort`, `status`, `hidden`, `no_cache`, `always_show`, `breadcrumb`, `active_menu`, `parent_id`, `creator`) VALUES (4, '2023-02-24 19:51:37.316', '2023-02-24 19:51:37.316', NULL, 'FieldRelation', '字段关系管理', 'el-icon-s-tools', 'fieldRelation', '', '/personnel/fieldRelation/index', 8, 1, 2, 2, 2, 1, '', 1, '系统');
INSERT INTO `menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `title`, `icon`, `path`, `redirect`, `component`, `sort`, `status`, `hidden`, `no_cache`, `always_show`, `breadcrumb`, `active_menu`, `parent_id`, `creator`) VALUES (5, '2023-02-24 19:51:37.316', '2023-02-24 19:51:37.316', NULL, 'System', '系统管理', 'component', '/system', '/system/role', 'Layout', 9, 1, 2, 2, 2, 1, '', 0, '系统');
INSERT INTO `menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `title`, `icon`, `path`, `redirect`, `component`, `sort`, `status`, `hidden`, `no_cache`, `always_show`, `breadcrumb`, `active_menu`, `parent_id`, `creator`) VALUES (6, '2023-02-24 19:51:37.316', '2023-02-24 19:51:37.316', NULL, 'Role', '角色管理', 'eye-open', 'role', '', '/system/role/index', 10, 1, 2, 2, 2, 1, '', 5, '系统');
INSERT INTO `menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `title`, `icon`, `path`, `redirect`, `component`, `sort`, `status`, `hidden`, `no_cache`, `always_show`, `breadcrumb`, `active_menu`, `parent_id`, `creator`) VALUES (7, '2023-02-24 19:51:37.316', '2023-02-24 19:51:37.316', NULL, 'Menu', '菜单管理', 'tree-table', 'menu', '', '/system/menu/index', 13, 1, 2, 2, 2, 1, '', 5, '系统');
INSERT INTO `menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `title`, `icon`, `path`, `redirect`, `component`, `sort`, `status`, `hidden`, `no_cache`, `always_show`, `breadcrumb`, `active_menu`, `parent_id`, `creator`) VALUES (8, '2023-02-24 19:51:37.316', '2023-02-24 19:51:37.316', NULL, 'Api', '接口管理', 'tree', 'api', '', '/system/api/index', 14, 1, 2, 2, 2, 1, '', 5, '系统');
INSERT INTO `menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `title`, `icon`, `path`, `redirect`, `component`, `sort`, `status`, `hidden`, `no_cache`, `always_show`, `breadcrumb`, `active_menu`, `parent_id`, `creator`) VALUES (9, '2023-02-24 19:51:37.316', '2023-02-24 19:51:37.316', NULL, 'Log', '日志管理', 'example', '/log', '/log/operation-log', 'Layout', 20, 1, 2, 2, 2, 1, '', 0, '系统');
INSERT INTO `menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `title`, `icon`, `path`, `redirect`, `component`, `sort`, `status`, `hidden`, `no_cache`, `always_show`, `breadcrumb`, `active_menu`, `parent_id`, `creator`) VALUES (10, '2023-02-24 19:51:37.316', '2023-02-24 19:51:37.316', NULL, 'OperationLog', '操作日志', 'documentation', 'operation-log', '', '/log/operation-log/index', 21, 1, 2, 2, 2, 1, '', 9, '系统');
COMMIT;

-- ----------------------------
-- Table structure for operation_logs
-- ----------------------------
DROP TABLE IF EXISTS `operation_logs`;
CREATE TABLE `operation_logs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(20) DEFAULT NULL COMMENT '''用户登录名''',
  `ip` varchar(20) DEFAULT NULL COMMENT '''Ip地址''',
  `ip_location` varchar(20) DEFAULT NULL COMMENT '''Ip所在地''',
  `method` varchar(20) DEFAULT NULL COMMENT '''请求方式''',
  `path` varchar(100) DEFAULT NULL COMMENT '''访问路径''',
  `remark` varchar(100) DEFAULT NULL COMMENT '''备注''',
  `status` int(4) DEFAULT NULL COMMENT '''响应状态码''',
  `start_time` varchar(2048) DEFAULT NULL COMMENT '''发起时间''',
  `time_cost` int(6) DEFAULT NULL COMMENT '''请求耗时(ms)''',
  `user_agent` varchar(2048) DEFAULT NULL COMMENT '''浏览器标识''',
  PRIMARY KEY (`id`),
  KEY `idx_operation_logs_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of operation_logs
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for role_menus
-- ----------------------------
DROP TABLE IF EXISTS `role_menus`;
CREATE TABLE `role_menus` (
  `menu_id` bigint(20) unsigned NOT NULL,
  `role_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`menu_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of role_menus
-- ----------------------------
BEGIN;
INSERT INTO `role_menus` (`menu_id`, `role_id`) VALUES (1, 1);
INSERT INTO `role_menus` (`menu_id`, `role_id`) VALUES (2, 1);
INSERT INTO `role_menus` (`menu_id`, `role_id`) VALUES (3, 1);
INSERT INTO `role_menus` (`menu_id`, `role_id`) VALUES (4, 1);
INSERT INTO `role_menus` (`menu_id`, `role_id`) VALUES (5, 1);
INSERT INTO `role_menus` (`menu_id`, `role_id`) VALUES (6, 1);
INSERT INTO `role_menus` (`menu_id`, `role_id`) VALUES (7, 1);
INSERT INTO `role_menus` (`menu_id`, `role_id`) VALUES (8, 1);
INSERT INTO `role_menus` (`menu_id`, `role_id`) VALUES (9, 1);
INSERT INTO `role_menus` (`menu_id`, `role_id`) VALUES (9, 2);
INSERT INTO `role_menus` (`menu_id`, `role_id`) VALUES (10, 1);
INSERT INTO `role_menus` (`menu_id`, `role_id`) VALUES (10, 2);
COMMIT;

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(20) NOT NULL,
  `keyword` varchar(20) NOT NULL,
  `remark` varchar(100) DEFAULT NULL COMMENT '''备注''',
  `status` tinyint(1) DEFAULT '1' COMMENT '''1正常, 2禁用''',
  `sort` int(3) DEFAULT '999' COMMENT '''角色排序(排序越大权限越低, 不能查看比自己序号小的角色, 不能编辑同序号用户权限, 排序为1表示超级管理员)''',
  `creator` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `keyword` (`keyword`),
  KEY `idx_roles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of roles
-- ----------------------------
BEGIN;
INSERT INTO `roles` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `keyword`, `remark`, `status`, `sort`, `creator`) VALUES (1, '2023-02-24 19:51:37.286', '2023-02-24 19:51:37.286', NULL, '管理员', 'admin', '', 1, 1, '系统');
INSERT INTO `roles` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `keyword`, `remark`, `status`, `sort`, `creator`) VALUES (2, '2023-02-24 19:51:37.286', '2023-02-24 19:51:37.286', NULL, '普通用户', 'user', '', 1, 3, '系统');
INSERT INTO `roles` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `keyword`, `remark`, `status`, `sort`, `creator`) VALUES (3, '2023-02-24 19:51:37.286', '2023-02-24 19:51:37.286', NULL, '访客', 'guest', '', 1, 5, '系统');
COMMIT;

-- ----------------------------
-- Table structure for user_roles
-- ----------------------------
DROP TABLE IF EXISTS `user_roles`;
CREATE TABLE `user_roles` (
  `role_id` bigint(20) unsigned NOT NULL,
  `user_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`role_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of user_roles
-- ----------------------------
BEGIN;
INSERT INTO `user_roles` (`role_id`, `user_id`) VALUES (1, 1);
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(50) NOT NULL COMMENT '''用户名''',
  `password` varchar(255) NOT NULL COMMENT '''用户密码''',
  `nickname` varchar(50) DEFAULT NULL COMMENT '''中文名''',
  `given_name` varchar(50) DEFAULT NULL COMMENT '''花名''',
  `mail` varchar(100) DEFAULT NULL COMMENT '''邮箱''',
  `job_number` varchar(20) DEFAULT NULL COMMENT '''工号''',
  `mobile` varchar(15) NOT NULL COMMENT '''手机号''',
  `avatar` varchar(255) DEFAULT NULL COMMENT '''头像''',
  `postal_address` varchar(255) DEFAULT NULL COMMENT '''地址''',
  `departments` varchar(128) DEFAULT NULL COMMENT '''部门''',
  `position` varchar(128) DEFAULT NULL COMMENT '''职位''',
  `introduction` varchar(255) DEFAULT NULL COMMENT '''个人简介''',
  `status` tinyint(1) DEFAULT '1' COMMENT '''状态:1在职, 2离职''',
  `creator` varchar(20) DEFAULT NULL COMMENT '''创建者''',
  `source` varchar(50) DEFAULT NULL COMMENT '''用户来源：dingTalk、wecom、feishu、ldap、platform''',
  `department_id` varchar(100) NOT NULL COMMENT '''部门id''',
  `source_user_id` varchar(100) NOT NULL COMMENT '''第三方用户id''',
  `source_union_id` varchar(100) NOT NULL COMMENT '''第三方唯一unionId''',
  `user_dn` varchar(255) NOT NULL COMMENT '''用户dn''',
  `sync_state` tinyint(1) DEFAULT '1' COMMENT '''同步状态:1已同步, 2未同步''',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `mobile` (`mobile`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `username`, `password`, `nickname`, `given_name`, `mail`, `job_number`, `mobile`, `avatar`, `postal_address`, `departments`, `position`, `introduction`, `status`, `creator`, `source`, `department_id`, `source_user_id`, `source_union_id`, `user_dn`, `sync_state`) VALUES (1, '2023-02-24 19:51:37.344', '2023-02-24 19:51:37.344', NULL, 'admin', 'A3zQm77nc4xFqHRfTUhjFelycRdYrw1B11YXom+rahmssQIrBXsriHTxJQvG5qXi1HqvPpoFEV9483yRCkcqvxr8l68ZyLOpZOPMyCRBMkYIDgMz4b9q7q5aqiyXTE6bhmv0SkBvQf7qYmRMIKu0QKSxaCg4RWKqBOp6WUtQD0s=', '管理员', '最强后台', 'admin@eryajf.net', '0000', '18888888888', 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif', '地球', '研发中心', '打工人', '最强后台的管理员', 1, '系统', '', '', '', '', 'cn=admin,dc=eryajf,dc=net', 1);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;