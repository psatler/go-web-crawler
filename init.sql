CREATE DATABASE /*!32312 IF NOT EXISTS*/ `demodb` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `demodb`;

--
-- Table structure for table `ibovespa`
--

DROP TABLE IF EXISTS `ibovespa`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ibovespa`
(
  `paperName` varchar
(255) DEFAULT NULL,
  `companyName` varchar
(255) DEFAULT NULL,
  `dailyRate` varchar
(20) DEFAULT NULL,
  `marketValue` float DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;