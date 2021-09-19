CREATE TABLE IF NOT EXISTS `users` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `full_name` varchar(150) NOT NULL,
    `email` varchar(150) NOT NULL,
    `password` varchar(150) NOT NULL,
    `role` varchar(150) NOT NULL,
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
);