CREATE TABLE `images` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `id_product` int DEFAULT NULL,
  `name` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;