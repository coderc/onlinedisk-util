drop table if exists `table_file`;
CREATE TABLE `table_file` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `uuid` BIGINT(20) NOT NULL,
    `sha1` VARCHAR(256) NOT NULL,
    `user_id` BIGINT(20) NOT NULL,
    `name` VARCHAR(256) NOT NULL,
    `path` VARCHAR(256) NOT NULL,
    `size` BIGINT(20) NOT NULL,
    `status` INT(11) NOT NULL DEFAULT 0 COMMENT '0: ok, 1: delete, 2: unable',
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `ex1` INT(11) NOT NULL DEFAULT 0 COMMENT '扩展字段1',
    `ex2` VARCHAR(256) NOT NULL DEFAULT '' COMMENT '扩展字段2',
    PRIMARY KEY (`id`),
    UNIQUE KEY `sha1_unique` (`sha1`),
    UNIQUE KEY `uuid_unique` (`uuid`),
    INDEX `name_index` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
