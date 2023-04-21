CREATE TABLE `logs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_email` longtext,
  `table_model` longtext,
  `entity_id` bigint unsigned DEFAULT NULL,
  `old_value` longtext,
  `new_value` longtext,
  `timestamp` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;