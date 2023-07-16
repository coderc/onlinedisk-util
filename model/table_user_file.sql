drop table if exists `table_user_file`;
CREATE TABLE `table_user_file` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT(20) NOT NULL,
    `file_id` BIGINT(20) NOT NULL,
    `file_name` VARCHAR(256) NOT NULL COMMENT '文件名 用户设定',
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `user_id_index` (`user_id`),
    INDEX `file_id_index` (`file_id`),
    FOREIGN KEY (`user_id`) REFERENCES `table_user` (`uuid`),
    FOREIGN KEY (`file_id`) REFERENCES `table_file` (`uuid`),
    UNIQUE KEY `user_id_file_id_unique` (`user_id`, `file_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
