-- MySQL dump 10.13  Distrib 5.7.35, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: decorshop2
-- ------------------------------------------------------
-- Server version	5.7.35

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `collections`
--

DROP TABLE IF EXISTS `collections`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `collections` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `image` longtext,
  `collection_name` varchar(225) DEFAULT NULL,
  `page_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_collections_deleted_at` (`deleted_at`),
  KEY `fk_pages_collection` (`page_id`),
  CONSTRAINT `fk_pages_collections` FOREIGN KEY (`page_id`) REFERENCES `pages` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `collections`
--

LOCK TABLES `collections` WRITE;
/*!40000 ALTER TABLE `collections` DISABLE KEYS */;
INSERT INTO `collections` VALUES (1,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/collections/table-lamp_large.jpg?v=1592809470','table-lamps',3),(2,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/collections/hurricane_large.jpg?v=1592876146','hurricane-lamps-and-candelabras',3),(3,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/collections/ceilinglights_large.jpg?v=1592809080','ceiling-lights',3),(4,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/collections/23_large.png?v=1592809029','chandeliers',3),(5,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/collections/floorlamps_large.jpg?v=1592809117','floor-lamps',3),(6,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/collections/wall-lamps_large.jpg?v=1477451342','wall-lights',3),(7,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/collections/chairs-and-seating_large.jpg?v=1592891963','chairs-and-seating',1),(8,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/collections/tables_e2083375-eab8-4ced-9770-3ac044025b9f_large.jpg?v=1592893134','tables',1),(9,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/collections/storage_18323929-116f-4f3c-bc71-fa3bde001c5a_large.jpg?v=1592891465','storage',1),(10,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/collections/misc_large.jpg?v=1592894775','miscellaneous',1),(11,'2021-09-18 16:15:42.000',NULL,NULL,'https://cdn.shopify.com/s/files/1/1089/1214/files/bedroom-furniture_480x480.png?v=1592895188','bedroom-furniture',1),(12,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/collections/plant_large.png?v=1592964346','artificial-flowers-and-trees',2),(13,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/collections/60_large.png?v=1592961625','wall-clocks',2),(14,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/collections/cushions_large.png?v=1592956310','cushions-throws-and-quilts',2),(15,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/collections/62_large.png?v=1592961692','decorator-pieces',2),(16,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/collections/doormat1_large.png?v=1592955669','designer-doormats',2),(17,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/products/shopify_38999_271b5f8c-a208-4a60-b3d2-f639cb8864cd_large.jpg?v=1592370310','room-dividers',2),(18,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/products/shopify_c7db46ad886ae97a52c9befcd911e37d_venetian-wall-fountain_large.jpg?v=1611704026','water-features',2),(19,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/collections/rug_large.png?v=1592881829','designer-rugs',2),(20,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/collections/mirrors_large.jpg?v=1592882163','decorative-mirrors',2),(21,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/collections/table-lamp_large.jpg?v=1592809470','lighting',2),(23,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/collections/58_large.png?v=1592959587','round-mirrors',2),(24,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/collections/art_canvas_large.png?v=1592958530','paintings',2),(25,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/collections/56_large.png?v=1592959544','benches',2),(26,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/collections/outdoor-accessories_7eccd0ed-1a9e-47cb-96f0-ca3036e07ed3_large.png?v=1477466490','outdoor-accessories',4),(27,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/collections/outdoor-seating_large.jpg?v=1592877148','outdoor-seating',4),(28,'2021-09-18 16:15:42.000',NULL,NULL,'//cdn.shopify.com/s/files/1/1089/1214/collections/outdoorsetting_large.jpg?v=1592876884','outdoor-tables',4);
/*!40000 ALTER TABLE `collections` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-09-21  9:40:38
