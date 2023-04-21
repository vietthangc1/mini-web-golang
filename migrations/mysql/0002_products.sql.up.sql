CREATE TABLE `products` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `sku` varchar(150) NOT NULL,
  `name` varchar(150) NOT NULL,
  `price` double NOT NULL,
  `number` bigint NOT NULL,
  `description` longtext,
  `cate1` longtext,
  `cate2` longtext,
  `cate3` longtext,
  `cate4` longtext,
  `user_email` varchar(150) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_products_deleted_at` (`deleted_at`),
  KEY `fk_users_products` (`user_email`),
  CONSTRAINT `fk_users_products` FOREIGN KEY (`user_email`) REFERENCES `users` (`email`),
  CONSTRAINT `chk_products_name` CHECK ((`name` <> _utf8mb3'')),
  CONSTRAINT `chk_products_number` CHECK ((`number` > 0)),
  CONSTRAINT `chk_products_price` CHECK ((`price` > 0)),
  CONSTRAINT `chk_products_sku` CHECK ((`sku` <> _utf8mb3''))
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;