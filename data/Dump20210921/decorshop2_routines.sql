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
-- Temporary table structure for view `product_detail_info`
--

DROP TABLE IF EXISTS `product_detail_info`;
/*!50001 DROP VIEW IF EXISTS `product_detail_info`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE VIEW `product_detail_info` AS SELECT 
 1 AS `collection_id`,
 1 AS `collection_name`,
 1 AS `collection_image`,
 1 AS `category_id`,
 1 AS `category_name`,
 1 AS `product_id`,
 1 AS `product_name`,
 1 AS `description`,
 1 AS `variance_id`,
 1 AS `color`,
 1 AS `size`,
 1 AS `price`,
 1 AS `inventory`*/;
SET character_set_client = @saved_cs_client;

--
-- Final view structure for view `product_detail_info`
--

/*!50001 DROP VIEW IF EXISTS `product_detail_info`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8 */;
/*!50001 SET character_set_results     = utf8 */;
/*!50001 SET collation_connection      = utf8_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `product_detail_info` AS select `c`.`id` AS `collection_id`,`c`.`collection_name` AS `collection_name`,`c`.`image` AS `collection_image`,`cg`.`id` AS `category_id`,`cg`.`category_name` AS `category_name`,`cp`.`product_id` AS `product_id`,`p`.`name` AS `product_name`,`p`.`description` AS `description`,`pv`.`id` AS `variance_id`,`pv`.`color` AS `color`,`pv`.`size` AS `size`,`pv`.`price` AS `price`,`pv`.`inventory` AS `inventory` from ((((`collections` `c` left join `categories` `cg` on((`c`.`id` = `cg`.`collection_id`))) left join `category_products` `cp` on((`cg`.`id` = `cp`.`category_id`))) left join `products` `p` on((`cp`.`product_id` = `p`.`id`))) left join `product_variances` `pv` on((`p`.`id` = `pv`.`product_id`))) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-09-21  9:40:38
