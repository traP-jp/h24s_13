CREATE TABLE `users`
(
    `id`   VARCHAR(32) PRIMARY KEY COMMENT 'traQ ID',
    `name` VARCHAR(32) COMMENT '表示名'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin
    COMMENT = 'ユーザー';

CREATE TABLE `user_groups`
(
    `id`   VARCHAR(32) PRIMARY KEY COMMENT 'traQ ID',
    `name` VARCHAR(32) COMMENT 'traQ Group Name'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin
    COMMENT = 'ユーザーグループ';

CREATE TABLE `user_connections`
(
    `id_1`     VARCHAR(32) COMMENT 'traQ ID',
    `id_2`     VARCHAR(32) COMMENT 'traQ ID',
    `strength` DOUBLE COMMENT 'つながりの強さ',
    PRIMARY KEY (`id_1`, `id_2`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin
    COMMENT = 'つながりの強さ';
