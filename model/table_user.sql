drop table if exists `table_user`;
CREATE TABLE `table_user` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `uuid` BIGINT(20) NOT NULL,
    `username` VARCHAR(64) NOT NULL,
    `password` VARCHAR(256) NOT NULL,
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `username_unique` (`username`),
    UNIQUE KEY `uuid_unique` (`uuid`),
    INDEX `username_index` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
