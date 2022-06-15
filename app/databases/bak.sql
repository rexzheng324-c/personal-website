-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`         varchar(128) NOT NULL COMMENT 'id',
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `nickname`   varchar(256) NOT NULL COMMENT '用户昵称',
    `role`       tinyint(3)   NOT NULL COMMENT '用户角色',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4;

-- ----------------------------
-- Table structure for basic_auth
-- ----------------------------
DROP TABLE IF EXISTS `basic_auth`;
CREATE TABLE `basic_auth`
(
    `id`         varchar(128) NOT NULL COMMENT 'id',
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `user_id`    varchar(128) NOT NULL COMMENT '用户 id',
    `username`   varchar(128) NOT NULL COMMENT '用户名',
    `password`   varchar(500) NOT NULL COMMENT '用户密码',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uk_username` (`username`)
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4;

-- ----------------------------
-- Table structure for blog
-- ----------------------------
DROP TABLE IF EXISTS `blog`;
CREATE TABLE `blog`
(
    `id`         varchar(128) NOT NULL COMMENT 'id',
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `user_id`    varchar(128) NOT NULL COMMENT '用户 id',
    `title`      varchar(128) NOT NULL COMMENT '标题',
    `content`    text         NOT NULL COMMENT '内容',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4