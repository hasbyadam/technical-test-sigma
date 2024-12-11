-- MySQL dump 10.13  Distrib 8.0.19, for Win64 (x86_64)
--
-- Host: localhost    Database: sigma_test
-- ------------------------------------------------------
-- Server version	8.0.40

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Temporary view structure for view `newview`
--

DROP TABLE IF EXISTS `newview`;
/*!50001 DROP VIEW IF EXISTS `newview`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `newview` AS SELECT 
 1 AS `created_by`,
 1 AS `credit_limit_id`,
 1 AS `tenor`,
 1 AS `limit_balance`*/;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transactions` (
  `id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `contract_number` varchar(100) DEFAULT NULL,
  `otr` int DEFAULT NULL,
  `admin_fee` int DEFAULT NULL,
  `installment` int DEFAULT NULL,
  `interest` float DEFAULT NULL,
  `asset_name` varchar(100) DEFAULT NULL,
  `created_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `created_at` int NOT NULL,
  `is_settled` tinyint(1) DEFAULT NULL,
  `credit_limit_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `transactions_unique` (`contract_number`),
  KEY `transactions_users_FK` (`created_by`),
  KEY `transactions_user_credit_limits_FK` (`credit_limit_id`),
  CONSTRAINT `transactions_user_credit_limits_FK` FOREIGN KEY (`credit_limit_id`) REFERENCES `user_credit_limits` (`id`),
  CONSTRAINT `transactions_users_FK` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactions`
--

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
INSERT INTO `transactions` VALUES ('60f5e91a-4905-48d6-bff2-f2c942145e52','4',1,1,10000,1,'test','71d510ba-092a-4bce-a8a8-f691ad790dd6',1,1,'82e47570-fcd4-4115-9e80-a9b13c9bb124'),('83661d00-5465-4482-a1cc-6e1b0021ea1e','CN-1733935847037',1,1,10000,1,'Test','71d510ba-092a-4bce-a8a8-f691ad790dd6',1733935847,0,'82e47570-fcd4-4115-9e80-a9b13c9bb124'),('886f5b90-8ddb-45f0-838e-7e9f4ad47ae4','CN-1733935776118',1,1,10000,1,'Test','71d510ba-092a-4bce-a8a8-f691ad790dd6',1733935776,0,'82e47570-fcd4-4115-9e80-a9b13c9bb124'),('9cb9f9dd-6120-46ff-90ad-dfc8318cea99','2',1,1,1000,1,'test','71d510ba-092a-4bce-a8a8-f691ad790dd6',1,1,'1f43444b-7573-4f11-bb1f-b6132bafa78a'),('de87044b-2508-4e71-ac20-c76c63e36250','1',1,1,500,1,'test','71d510ba-092a-4bce-a8a8-f691ad790dd6',1,0,'1f43444b-7573-4f11-bb1f-b6132bafa78a'),('ea8ec609-adec-478a-9f36-78fa78675f7d','3',1,1,5000,1,'test','71d510ba-092a-4bce-a8a8-f691ad790dd6',1,0,'2ea5d44a-9bea-421f-80e7-123e2463f1a9');
/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_credit_limits`
--

DROP TABLE IF EXISTS `user_credit_limits`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_credit_limits` (
  `id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `user_id` varchar(100) DEFAULT NULL,
  `tenor` int DEFAULT NULL,
  `credit` int DEFAULT NULL,
  `created_at` int DEFAULT NULL,
  `updated_at` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_credit_limits_unique` (`user_id`,`tenor`),
  CONSTRAINT `user_credit_limits_users_FK` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_credit_limits`
--

LOCK TABLES `user_credit_limits` WRITE;
/*!40000 ALTER TABLE `user_credit_limits` DISABLE KEYS */;
INSERT INTO `user_credit_limits` VALUES ('1f43444b-7573-4f11-bb1f-b6132bafa78a','71d510ba-092a-4bce-a8a8-f691ad790dd6',1,10000,NULL,NULL),('2ea5d44a-9bea-421f-80e7-123e2463f1a9','71d510ba-092a-4bce-a8a8-f691ad790dd6',2,20000,NULL,NULL),('82e47570-fcd4-4115-9e80-a9b13c9bb124','71d510ba-092a-4bce-a8a8-f691ad790dd6',3,50000,NULL,NULL);
/*!40000 ALTER TABLE `user_credit_limits` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary view structure for view `user_limit_balances`
--

DROP TABLE IF EXISTS `user_limit_balances`;
/*!50001 DROP VIEW IF EXISTS `user_limit_balances`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `user_limit_balances` AS SELECT 
 1 AS `user_id`,
 1 AS `credit_limit_id`,
 1 AS `tenor`,
 1 AS `limit_balance`*/;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `full_name` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `legal_name` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `nik` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `birth_place` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `birth_date` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `salary` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `ktp` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `ktp_selfie` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_at` int DEFAULT NULL,
  `created_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `is_verified` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email_unique` (`email`),
  UNIQUE KEY `nik_unique` (`nik`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('40daa61f-751b-431a-b56d-bd1e60a86f0f','$2a$10$RJFVpqU.WO0lnW//WkdqquY6qpYMFDF.y3dbFDVHVrVJSVs5p52..','hasby3@mail.com',NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,1733935873,'40daa61f-751b-431a-b56d-bd1e60a86f0f',0),('71d510ba-092a-4bce-a8a8-f691ad790dd6','$2a$10$J4zN.t2Gvr.5wxTUwTrrV.URsR5tAvThk71b2ZsJ2Rl5jLtth5kVK','hasby@mail.com','','','','','','','','',1733914344,'71d510ba-092a-4bce-a8a8-f691ad790dd6',0),('e7f5bf2c-868b-4613-9278-44b9dec0c433','test','test','test','test','test','test','test','test','test','test',1733914344,'test',1);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'sigma_test'
--

--
-- Final view structure for view `newview`
--

/*!50001 DROP VIEW IF EXISTS `newview`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `newview` AS select `t`.`created_by` AS `created_by`,`t`.`credit_limit_id` AS `credit_limit_id`,`l`.`tenor` AS `tenor`,(`l`.`credit` - sum((case when (`t`.`is_settled` = false) then `t`.`installment` else 0 end))) AS `limit_balance` from (`transactions` `t` left join `user_credit_limits` `l` on((`l`.`id` = `t`.`credit_limit_id`))) group by `t`.`credit_limit_id`,`l`.`tenor`,`t`.`created_by` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `user_limit_balances`
--

/*!50001 DROP VIEW IF EXISTS `user_limit_balances`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `user_limit_balances` AS select `l`.`user_id` AS `user_id`,`l`.`id` AS `credit_limit_id`,`l`.`tenor` AS `tenor`,(`l`.`credit` - sum((case when (`t`.`is_settled` = false) then `t`.`installment` else 0 end))) AS `limit_balance` from (`user_credit_limits` `l` left join `transactions` `t` on(((`t`.`credit_limit_id` = `l`.`id`) and (`t`.`created_by` = `l`.`user_id`)))) group by `l`.`id`,`l`.`tenor`,`l`.`user_id` */;
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

-- Dump completed on 2024-12-12  1:18:54
