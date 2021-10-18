CREATE productsCREATE TABLE `products` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name_product` varchar(300) DEFAULT NULL,
  `price` int DEFAULT NULL,
  `id_category` int DEFAULT NULL,
  `created_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;