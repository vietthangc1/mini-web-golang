CREATE TABLE `propertises` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `product_id` bigint unsigned DEFAULT NULL,
  `attribute` longtext NOT NULL,
  `value` longtext NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_propertises_deleted_at` (`deleted_at`),
  KEY `fk_products_propertises` (`product_id`),
  CONSTRAINT `fk_products_propertises` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;