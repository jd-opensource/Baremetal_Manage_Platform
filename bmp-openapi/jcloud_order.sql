[mysqldump -h 10.208.12.72 -u root -pLpK9Jq12Zf jcloud_order] is a warning command
-- MySQL dump 10.16  Distrib 10.1.17-MariaDB, for Linux (x86_64)
--
-- Host: 10.208.12.72    Database: jcloud_order
-- ------------------------------------------------------
-- Server version	5.7.40

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
-- Table structure for table `ads_user_resource_info`
--

DROP TABLE IF EXISTS `ads_user_resource_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ads_user_resource_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `pin` varchar(64) DEFAULT NULL COMMENT 'pin',
  `type` int(11) DEFAULT NULL COMMENT '类型',
  `state` char(1) DEFAULT NULL COMMENT '类型',
  `resource_id` varchar(64) DEFAULT NULL COMMENT '资源ID',
  `source_id` varchar(64) DEFAULT NULL COMMENT '三方id',
  `order_number` varchar(64) DEFAULT NULL COMMENT '订单号',
  `create_time` char(19) DEFAULT NULL COMMENT '创建时间',
  `update_time` char(19) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=565 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ads_user_resource_info`
--

LOCK TABLES `ads_user_resource_info` WRITE;
/*!40000 ALTER TABLE `ads_user_resource_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `ads_user_resource_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `business_payment`
--

DROP TABLE IF EXISTS `business_payment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `business_payment` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增数据库主键',
  `business_code` varchar(50) NOT NULL COMMENT '业务主体Id',
  `business_name` varchar(50) NOT NULL COMMENT '业务主体名称',
  `payment_method` varchar(100) NOT NULL COMMENT '支付方式（uc的dictionary表中的key字段）',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态：1-上线，2-下线',
  `creator` varchar(50) DEFAULT NULL COMMENT '创建者',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `index_key` (`business_code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8 COMMENT='业务主体和支付方式关系表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `business_payment`
--

LOCK TABLES `business_payment` WRITE;
/*!40000 ALTER TABLE `business_payment` DISABLE KEYS */;
INSERT INTO `business_payment` VALUES (18,'612','360度','companyBankpay','企业网银',1,NULL,NULL,NULL),(19,'612','360度','personalBankPay','个人网银',1,NULL,NULL,NULL),(20,'612','360度','offLinePay','线下汇款充值',1,NULL,NULL,NULL),(23,'612','360度','jdPay','京东支付',1,NULL,NULL,NULL),(24,'612','360度','weiXinPay','微信支付',1,NULL,NULL,NULL),(25,'612','360度','balancePay','余额支付',1,NULL,NULL,NULL),(26,'543','江苏京东','companyBankpay','企业网银',1,NULL,NULL,NULL),(27,'543','江苏京东','personalBankPay','个人网银',1,NULL,NULL,NULL),(28,'543','江苏京东','jdPay','京东支付',1,NULL,NULL,NULL),(29,'656','尚科','companyBankpay','企业网银',1,NULL,NULL,NULL),(30,'656','尚科','personalBankPay','个人网银',1,NULL,NULL,NULL),(31,'656','尚科','jdPay','京东支付',1,NULL,NULL,NULL);
/*!40000 ALTER TABLE `business_payment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `business_service`
--

DROP TABLE IF EXISTS `business_service`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `business_service` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增数据库主键',
  `business_code` varchar(50) NOT NULL COMMENT '业务主体Id',
  `business_name` varchar(50) NOT NULL COMMENT '业务主体名称',
  `site_type_id` int(4) NOT NULL COMMENT '站点类型Id',
  `site_type_name` varchar(50) NOT NULL COMMENT '站点类型',
  `app_code` varchar(50) NOT NULL COMMENT '业务线code',
  `app_name` varchar(50) NOT NULL COMMENT '业务线名称',
  `service_code` varchar(50) NOT NULL COMMENT '产品code',
  `service_name` varchar(50) NOT NULL COMMENT '产品名称',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态：1-上线，2-下线',
  `creator` varchar(50) DEFAULT NULL COMMENT '创建者',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `index_key` (`site_type_id`,`app_code`,`service_code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8 COMMENT='业务主体和产品关系表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `business_service`
--

LOCK TABLES `business_service` WRITE;
/*!40000 ALTER TABLE `business_service` DISABLE KEYS */;
INSERT INTO `business_service` VALUES (18,'612','360度',0,'主站','domain','域名','domain','域名','域名服务',1,NULL,'2017-04-18 13:01:04','2017-04-18 13:01:10'),(19,'612','360度',0,'主站','jcloud','基础云','balance','负载均衡','云主机、子网、路由器、公网IP、负载均衡、云数据库、云硬盘、云存储、防火墙、监控报警、DDoS防护,直播',1,NULL,'2015-10-27 16:10:04','2015-10-27 16:10:04'),(20,'612','360度',0,'主站','jcloud','基础云','vm','云主机','云主机、子网、路由器、公网IP、负载均衡、云数据库、云硬盘、云存储、防火墙、监控报警、DDoS防护,直播',1,NULL,'2015-10-27 16:10:04','2015-10-27 16:10:04'),(21,'612','360度',0,'主站','jcloud','基础云','ip','公网IP','云主机、子网、路由器、公网IP、负载均衡、云数据库、云硬盘、云存储、防火墙、监控报警、DDoS防护,直播',1,NULL,'2016-01-20 14:29:23','2016-01-20 14:29:25'),(22,'612','360度',0,'主站','jcloud','基础云','database','云数据库MySQL','云主机、子网、路由器、公网IP、负载均衡、云数据库、云硬盘、云存储、防火墙、监控报警、DDoS防护,直播',1,NULL,'2016-01-22 15:45:45','2017-09-08 17:28:09'),(23,'612','360度',0,'主站','jcloud','基础云','disk','云硬盘','云主机、子网、路由器、公网IP、负载均衡、云数据库、云硬盘、云存储、防火墙、监控报警、DDoS防护,直播',1,NULL,'2016-01-25 11:11:31','2016-01-25 11:11:33'),(24,'612','360度',0,'主站','jcloud','基础云','storage','云存储','云主机、子网、路由器、公网IP、负载均衡、云数据库、云硬盘、云存储、防火墙、监控报警、DDoS防护,直播',1,NULL,'2016-02-29 16:19:56','2016-02-29 16:19:56'),(25,'612','360度',0,'主站','0003','数据云','DCS','数知','为开发者提供分布式、全托管、批量和实时数据计算服务。 ',1,NULL,'2016-03-29 15:13:13','2016-03-29 15:13:13'),(26,'612','360度',0,'主站','jcloud','基础云','redis','云缓存','云主机、子网、路由器、公网IP、负载均衡、云数据库、云硬盘、云存储、防火墙、监控报警、DDoS防护,直播',1,NULL,'2016-06-07 16:59:20','2016-06-07 16:59:20'),(27,'612','360度',0,'主站','jcloud','基础云','db_ro','云数据库MySQL-只读实例','云主机、子网、路由器、公网IP、负载均衡、云数据库、云硬盘、云存储、防火墙、监控报警、DDoS防护,直播',1,NULL,'2016-07-27 19:04:52','2017-09-08 17:04:52'),(28,'612','360度',0,'主站','jcloud','基础云','cdn','CDN','云主机、子网、路由器、公网IP、负载均衡、云数据库、云硬盘、云存储、防火墙、监控报警、DDoS防护,直播',1,NULL,'2016-08-25 16:09:45','2016-08-25 16:09:45'),(29,'612','360度',0,'主站','jcloud','基础云','nat_gw','NAT网关','云主机、子网、路由器、公网IP、负载均衡、云数据库、云硬盘、云存储、防火墙、监控报警、DDoS防护,直播',1,NULL,'2016-08-30 10:28:02','2016-08-30 10:28:02'),(30,'612','360度',0,'主站','jcloud','基础云','vpn','VPN','云主机、子网、路由器、公网IP、负载均衡、云数据库、云硬盘、云存储、防火墙、监控报警、DDoS防护,直播',1,NULL,'2016-08-30 10:28:02','2016-08-30 10:28:02'),(31,'612','360度',0,'主站','jcloud','基础云','ddos','BGP高防','云主机、子网、路由器、公网IP、负载均衡、云数据库、云硬盘、云存储、防火墙、监控报警、DDoS防护,直播',1,NULL,'2017-01-17 17:17:43','2017-01-17 17:17:44'),(32,'612','360度',0,'主站','jcloud','基础云','live','直播','云主机、子网、路由器、公网IP、负载均衡、云数据库、云硬盘、云存储、防火墙、监控报警、DDoS防护,直播',1,NULL,'2017-02-08 14:24:27','2017-02-08 14:24:27'),(34,'612','360度',0,'云市场','jcloudmark','云市场','web_self','云市场产品','自营建站',1,NULL,'2017-04-18 13:01:16','2017-04-18 13:01:18'),(35,'612','360度',0,'云市场','jcloudmark','云市场','mirror_self','自营镜像','自营镜像',1,NULL,'2017-04-18 13:01:22','2017-04-18 13:01:33'),(36,'543','江苏京东',0,'云市场','jcloudmark','云市场','web_other','云市场产品','第三方建站',1,NULL,'2017-04-18 13:01:28','2017-04-18 13:01:31'),(37,'543','江苏京东',0,'云市场','jcloudmark','云市场','mirror_other','第三方镜像','第三方镜像',1,NULL,'2017-04-18 13:01:37','2017-04-18 13:01:40'),(39,'612','360度',0,'主站','jcloud','基础云','mongodb','MongoDB','可信赖的基础服务\r\n可信赖的基础服务\r\n可信赖的基础服务',1,NULL,'2017-05-12 17:35:53','2017-05-12 17:35:53'),(40,'612','360度',0,'主站','jcloud','基础云','webanti','网站高防','网站高防',1,NULL,'2017-06-09 18:31:29','2017-06-09 18:31:37'),(41,'612','360度',0,'主站','jcloud','基础云','sqlserver','云数据库SQL Server','云数据库SQL Server',1,NULL,'2017-07-24 14:24:27','2017-07-24 14:24:27'),(42,'612','360度',0,'主站','jcloud','基础云','ipanti','IP高防','IP高防',1,NULL,'2017-11-30 17:46:56','2017-11-30 17:46:56'),(43,'612','360度',0,'主站','jcloud','基础云','nativecontainer','容器服务','容器服务',1,NULL,'2017-12-25 11:24:59','2017-12-25 11:24:59');
/*!40000 ALTER TABLE `business_service` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `failure_task`
--

DROP TABLE IF EXISTS `failure_task`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `failure_task` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `params` varchar(2000) NOT NULL COMMENT '参数',
  `failure_type` tinyint(1) NOT NULL COMMENT '失败原因：1-服务异常，2-网络异常',
  `bean_name` varchar(100) NOT NULL COMMENT 'bean对象名称',
  `operation` varchar(200) NOT NULL COMMENT '调用失败的方法',
  `protocol` varchar(100) NOT NULL COMMENT '接口协议，如jsf、http',
  `url` varchar(500) DEFAULT NULL COMMENT '接口地址',
  `status` tinyint(1) NOT NULL COMMENT '状态：1-成功，2-失败',
  `retry_time` int(10) DEFAULT NULL COMMENT '重试次数',
  `msg` varchar(2000) DEFAULT NULL COMMENT '失败信息',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  `last_time` datetime DEFAULT NULL COMMENT '最后一次重试时间',
  `batch_id` varchar(100) DEFAULT NULL COMMENT '批次id',
  `sequence_no` int(10) DEFAULT NULL COMMENT '序号（针对batch_id字段）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=806 DEFAULT CHARSET=utf8 COMMENT='失败任务表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `failure_task`
--

LOCK TABLES `failure_task` WRITE;
/*!40000 ALTER TABLE `failure_task` DISABLE KEYS */;
/*!40000 ALTER TABLE `failure_task` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kafka_consumer_log`
--

DROP TABLE IF EXISTS `kafka_consumer_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kafka_consumer_log` (
  `id` int(12) NOT NULL AUTO_INCREMENT COMMENT '自增数据库主键',
  `topic` varchar(255) DEFAULT NULL COMMENT '主题',
  `partition` int(10) NOT NULL COMMENT '分区',
  `offset` int(10) DEFAULT NULL COMMENT '站点',
  `kafka_key` varchar(255) DEFAULT NULL COMMENT '消息的key值',
  `kafka_value` varchar(2000) NOT NULL COMMENT '消息的value',
  `md5` varchar(255) NOT NULL COMMENT '对消息value的md5 值',
  `batch_id` varchar(64) DEFAULT NULL COMMENT '批次号',
  `sequence_no` int(10) DEFAULT NULL COMMENT '订单信息序号',
  `status` tinyint(2) NOT NULL COMMENT '消息处理的状态 1、成功 2、失败',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `md5` (`md5`)
) ENGINE=InnoDB AUTO_INCREMENT=843348 DEFAULT CHARSET=utf8 COMMENT='kafka消费者日志';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kafka_consumer_log`
--

LOCK TABLES `kafka_consumer_log` WRITE;
/*!40000 ALTER TABLE `kafka_consumer_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `kafka_consumer_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kafka_producer_log`
--

DROP TABLE IF EXISTS `kafka_producer_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `kafka_producer_log` (
  `id` int(12) NOT NULL AUTO_INCREMENT COMMENT '自增数据库主键',
  `topic` varchar(255) DEFAULT NULL COMMENT '主题',
  `kafka_value` varchar(2000) NOT NULL COMMENT '消息的value',
  `md5` varchar(255) NOT NULL COMMENT '对消息value的md5 值',
  `batch_id` varchar(64) DEFAULT NULL COMMENT '批次号',
  `sequence_no` int(10) DEFAULT NULL COMMENT '订单信息序号',
  `status` tinyint(2) NOT NULL COMMENT '消息处理的状态 1、成功 2、失败',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `md5` (`md5`)
) ENGINE=InnoDB AUTO_INCREMENT=909203 DEFAULT CHARSET=utf8 COMMENT='kafka生产者日志';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kafka_producer_log`
--

LOCK TABLES `kafka_producer_log` WRITE;
/*!40000 ALTER TABLE `kafka_producer_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `kafka_producer_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order`
--

DROP TABLE IF EXISTS `order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增数据库主键',
  `order_number` varchar(50) NOT NULL COMMENT '订单编号',
  `parent_id` bigint(20) DEFAULT NULL COMMENT '父订单id，无父单时为-1',
  `process_type` tinyint(1) DEFAULT '1' COMMENT '流程类型：1、主站_标准流程_V1.0（当前标准下单接口）2、服务api_预付费下单免支付_V1.0（政务云需求场景，取消对接财务系统）3、服务api_标准下单流程_V1.0（OEM需求场景，取消对接财务系统）',
  `self_support_type` tinyint(1) DEFAULT '1' COMMENT '自营类型：1-自营，2-非自营，3-两者并存',
  `split_type` tinyint(1) DEFAULT NULL COMMENT '拆分类型，是否有子订单：0-无子单，1-有子单',
  `pin` varchar(30) NOT NULL COMMENT '用户pin',
  `proposer` varchar(200) DEFAULT NULL COMMENT '申请人',
  `payer` varchar(200) DEFAULT NULL COMMENT '付款人',
  `site_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '站点类型 0：主站,1：国际站,2：云市场,3：API',
  `order_type` tinyint(1) DEFAULT NULL COMMENT '订单类型：1-新购，2-续费，3-配置变更',
  `charge_mode` tinyint(1) DEFAULT NULL COMMENT '订单计费类型：1-按配置，2-按用量，3-按包年/包月，4-一次性付费',
  `app_code` varchar(50) DEFAULT NULL COMMENT '应用code：domain-域名，0003-数据云，jcloud-基础云，jcloudmark-云市场',
  `service_code` varchar(50) DEFAULT NULL COMMENT '产品 code',
  `item_name` varchar(255) DEFAULT NULL COMMENT '商品名称，用于订单列表显示',
  `status` tinyint(1) DEFAULT NULL COMMENT '订单状态：1-已支付，2-失败，3-未支付，4-处理中，5-已取消',
  `currency` tinyint(1) NOT NULL DEFAULT '1' COMMENT '币种 1：人民币,2：美元',
  `pay_type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '付费类型 1：预付费,2：后付费',
  `activity_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '活动订单类型：0-正常订单，1-活动订单',
  `activity_info` varchar(500) DEFAULT NULL COMMENT '活动信息，json串 [{&quot;promotionType&quot;:2,&quot;activityCode&quot;:123},{&quot;promotionType&quot;:3,&quot;activityCode&quot;:456}] ',
  `total_fee` decimal(14,2) NOT NULL COMMENT '订单总额（原价，不考虑折扣、优惠）',
  `actual_fee` decimal(14,2) DEFAULT NULL COMMENT '实际总额（如未改价则与计算总价一致，否则不一致）',
  `is_actual_fee_modified` tinyint(1) NOT NULL DEFAULT '0' COMMENT '应付价格是否被修改过：0-未修改，1-已修改；',
  `discount_fee` decimal(14,2) DEFAULT NULL COMMENT '优惠总额',
  `paid_fee` decimal(14,2) DEFAULT NULL COMMENT '已支付总额',
  `favorable_detail_json` varchar(512) DEFAULT NULL COMMENT '优惠明细，json结构，包括折扣、代金券、优惠券、商品券等促销信息的id、促销额、实际优惠额。',
  `voucher_pay` decimal(14,2) DEFAULT '0.00' COMMENT '代金券支付金额',
  `promotion_favorable_info` varchar(2000) DEFAULT NULL COMMENT '促销优惠明细',
  `balance_pay` decimal(14,2) DEFAULT NULL COMMENT '余额支付金额',
  `money_pay` decimal(14,2) DEFAULT NULL COMMENT '现金支付金额',
  `payment_channel` tinyint(1) DEFAULT NULL COMMENT '支付方式:0-余额支付，1-企业网银，2-个人网银，3-微信支付，4-京东支付，5-线下汇款',
  `created_time` datetime DEFAULT NULL COMMENT '订单创建时间',
  `updated_time` datetime DEFAULT NULL COMMENT '订单更新时间',
  `pay_time` datetime DEFAULT NULL COMMENT '支付时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `origin_order_number` varchar(50) DEFAULT NULL COMMENT '原订单编号，针对升降配、续费',
  `source_id` varchar(50) DEFAULT NULL COMMENT '来源标识',
  `task_id` varchar(40) DEFAULT NULL COMMENT '任务id：对应于中间层任务id，打包购买时的一个事务Id',
  `receipt_org_number` varchar(50) DEFAULT NULL COMMENT '收款机构编号',
  `payment_number` varchar(50) DEFAULT NULL COMMENT '支付单号',
  `return_url` varchar(500) DEFAULT NULL COMMENT '支付成功后的回调地址',
  `source` varchar(100) DEFAULT NULL COMMENT '订单来源，PC业务订单为空，OpenAPI订单固定为‘API’',
  `mod_count` int(10) DEFAULT NULL COMMENT '更改次数',
  `refund_fee` decimal(14,2) DEFAULT NULL COMMENT '已退款金额',
  `unique_code` varchar(100) DEFAULT NULL COMMENT 'MD5加密串',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_order_number` (`order_number`),
  UNIQUE KEY `unique_code` (`unique_code`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_pin_app_code` (`pin`,`app_code`),
  KEY `idx_pay_time` (`pay_time`),
  KEY `idx_source_id` (`source_id`),
  KEY `idx_created_time` (`created_time`)
) ENGINE=InnoDB AUTO_INCREMENT=2876954 DEFAULT CHARSET=utf8 COMMENT='订单';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order`
--

LOCK TABLES `order` WRITE;
/*!40000 ALTER TABLE `order` DISABLE KEYS */;
/*!40000 ALTER TABLE `order` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_configuration`
--

DROP TABLE IF EXISTS `order_configuration`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_configuration` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `process_type` int(4) NOT NULL COMMENT '流程类型',
  `site_type` int(4) NOT NULL COMMENT '站点类型',
  `app_code` varchar(50) NOT NULL COMMENT '业务线',
  `service_code` varchar(50) DEFAULT NULL COMMENT '产品code',
  `config_type` int(10) NOT NULL COMMENT '配置类型：1-签名key；2-配额接口域名',
  `test_value` varchar(200) DEFAULT NULL COMMENT '开发测试环境用',
  `online_value` varchar(200) DEFAULT NULL COMMENT '预发线上环境用',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态：1-开启，2-不开启',
  `remark` varchar(200) DEFAULT NULL COMMENT '备注',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_configuration` (`process_type`,`site_type`,`app_code`,`service_code`,`config_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=153 DEFAULT CHARSET=utf8 COMMENT='订单配置';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_configuration`
--

LOCK TABLES `order_configuration` WRITE;
/*!40000 ALTER TABLE `order_configuration` DISABLE KEYS */;
INSERT INTO `order_configuration` VALUES (105,1,0,'jcloud','jcloud',1,'3b41befe77a4c9cc62dd5141712a6246','0ee22ec2386f317f282701abc7fcb6cf',1,NULL,'2017-05-12 17:14:52','2017-05-12 17:14:52'),(106,1,0,'0003','0003',1,'b2cb9c22fe664a0c2600303dfe9dfecf','1f2da7d57abf331bf107efa58576743c',1,NULL,'2017-05-12 17:14:52','2017-05-12 17:14:52'),(107,1,0,'domain','domain',1,'9a539fc8de8605ca10af5d1e97d3cc67','caa6fe99f3e0fa064d18058297904f8f',1,NULL,'2017-05-12 17:14:52','2017-05-12 17:14:52'),(108,1,0,'jcloudmark','jcloudmark',1,'481b78e3e9dbd184f80859eec9ef65e1','451f679e2a2575dbca2ac21e19433e99',1,NULL,'2017-05-12 17:14:52','2017-05-12 17:14:52'),(110,1,0,'0003','DCS',3,'1','1',1,NULL,'2017-05-10 16:00:58','2017-05-10 16:00:58'),(111,1,0,'domain','domain',3,'1','1',1,NULL,'2017-05-10 16:00:58','2017-05-10 16:00:58'),(113,1,0,'jcloud','cdn',3,'1','1',1,NULL,'2017-05-10 16:00:58','2017-09-11 21:56:31'),(114,1,0,'jcloud','database',3,'1','1',1,NULL,'2017-05-10 16:00:58','2017-10-20 06:53:17'),(115,1,0,'jcloud','db_ro',3,'1','1',1,NULL,'2017-05-10 16:00:58','2017-10-20 06:54:13'),(116,1,0,'jcloud','ddos',3,'1','1',1,NULL,'2017-05-10 16:00:58','2017-09-28 18:50:13'),(119,1,0,'jcloud','live',3,'1','1',1,NULL,'2017-05-10 16:00:58','2017-09-28 14:38:03'),(120,1,0,'jcloud','nat_gw',3,'0','0',1,NULL,'2017-05-10 16:00:58','2017-05-10 16:00:58'),(122,1,0,'jcloud','storage',3,'1','1',1,NULL,'2017-05-10 16:00:58','2017-10-12 20:04:12'),(124,1,0,'jcloud','vpn',3,'0','0',1,NULL,'2017-05-10 16:00:58','2017-05-10 16:00:58'),(125,1,0,'jcloud','webanti',3,'1','1',1,NULL,'2017-05-10 16:00:58','2017-05-10 16:00:58'),(126,1,0,'jcloudmark','mirror_other',3,'1','1',1,NULL,'2017-05-10 16:00:58','2017-05-10 16:00:58'),(127,1,0,'jcloudmark','mirror_self',3,'1','1',1,NULL,'2017-05-10 16:00:58','2017-05-10 16:00:58'),(128,1,0,'jcloudmark','web_other',3,'1','1',1,NULL,'2017-05-10 16:00:58','2017-05-10 16:00:58'),(129,1,0,'jcloudmark','web_self',3,'1','1',1,NULL,'2017-05-10 16:00:58','2017-05-10 16:00:58'),(131,1,0,'jcloud','sqlserver',3,'1','1',1,NULL,'2017-07-24 14:00:58','2017-07-24 14:00:58'),(134,1,0,'jcloud','mongodb',3,'1','1',1,NULL,'2017-08-30 01:26:17','2017-08-30 01:26:17'),(135,1,0,'jcloud','redis',3,'1','1',1,NULL,'2017-08-30 11:45:44','2017-08-30 11:45:44'),(136,1,0,'jcloud','vm',2,'vpc-biz.jcloud.com','vpc-biz.jcloud.com',1,NULL,'2017-09-07 17:42:01','2017-09-07 17:42:01'),(137,1,0,'jcloud','disk',2,'vpc-biz.jcloud.com','vpc-biz.jcloud.com',1,NULL,'2017-09-07 17:42:11','2017-09-07 17:42:11'),(138,1,0,'jcloud','ip',2,'vpc-biz.jcloud.com','vpc-biz.jcloud.com',1,NULL,'2017-09-07 17:42:18','2017-09-07 17:42:18'),(139,1,0,'jcloud','balance',2,'vpc-biz.jcloud.com','vpc-biz.jcloud.com',1,NULL,'2017-09-07 17:42:27','2017-09-07 17:42:27'),(146,1,0,'jcloud','vm',3,'0','0',1,NULL,'2017-09-14 03:07:40','2017-09-14 03:07:40'),(147,1,0,'jcloud','disk',3,'0','0',1,NULL,'2017-09-14 03:07:57','2017-09-14 03:07:57'),(148,1,0,'jcloud','ip',3,'0','0',1,NULL,'2017-09-14 03:08:07','2017-09-14 03:08:07'),(149,1,0,'jcloud','balance',3,'0','0',1,NULL,'2017-09-14 03:08:16','2017-09-14 03:08:16'),(150,1,0,'jcloud','ipanti',3,'1','1',1,NULL,'2017-11-30 17:46:56','2017-11-30 17:46:56'),(151,4,0,'jcloud','jcloud',1,'4441befe77a4c9cc62dd1514712a6211','4ee22ec2386f317f323201abc7fcb611',1,NULL,'2017-12-19 16:50:03','2017-12-19 16:50:03'),(152,1,0,'jcloud','nativecontainer',3,'1','1',1,NULL,'2017-12-25 11:24:59','2017-12-25 11:24:59');
/*!40000 ALTER TABLE `order_configuration` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_extra`
--

DROP TABLE IF EXISTS `order_extra`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_extra` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `pin` varchar(30) NOT NULL COMMENT '用户pin',
  `order_number` varchar(50) NOT NULL COMMENT '订单编号',
  `provider_pin` varchar(100) NOT NULL COMMENT '服务商pin',
  `provider_name` varchar(200) NOT NULL COMMENT '服务商名称',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=387 DEFAULT CHARSET=utf8 COMMENT='订单扩展表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_extra`
--

LOCK TABLES `order_extra` WRITE;
/*!40000 ALTER TABLE `order_extra` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_extra` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_item`
--

DROP TABLE IF EXISTS `order_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_item` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增数据库主键',
  `order_number` varchar(50) DEFAULT NULL COMMENT '订单编号',
  `order_item_number` varchar(50) DEFAULT NULL COMMENT '订单商品编号',
  `pin` varchar(30) DEFAULT NULL COMMENT '用户pin',
  `charge_mode` tinyint(1) DEFAULT NULL COMMENT '订单计费类型：1-按配置，2-按用量，3-按包年/包月，4-一次性付费（用来支持单个订单多种计费类型的情况）',
  `charge_duration` int(10) DEFAULT NULL COMMENT '计费时长',
  `app_code` varchar(50) DEFAULT NULL COMMENT '应用code：domain-域名，0003-数据云，jcloud-基础云，jcloudmark-云市场',
  `service_code` varchar(50) DEFAULT NULL COMMENT '产品code',
  `region` varchar(50) DEFAULT NULL COMMENT '资源中心',
  `network_operator` tinyint(1) DEFAULT NULL COMMENT '网络类型 0: non-BGP, 1: BGP',
  `item_id` varchar(50) DEFAULT NULL COMMENT '商品id',
  `sku_id` varchar(50) DEFAULT NULL COMMENT 'sku code',
  `item_name` varchar(100) DEFAULT NULL COMMENT '商品名称',
  `item_type` varchar(50) DEFAULT NULL COMMENT '1-普通商品，2-活动商品',
  `promotion_info` varchar(500) DEFAULT NULL COMMENT '促销明细，json数组： [{&quot;promotionType&quot;:1,&quot;activityCode&quot;:123},{&quot;promotionType&quot;:2,&quot;activityCode&quot;:}] ',
  `resize_formula_type` tinyint(4) DEFAULT NULL COMMENT '变配明细，变配订单该字段不能为空：1-升配，2-降配，3-配置转包年包月',
  `extra_info` varchar(500) DEFAULT NULL COMMENT '商品规格参数',
  `quantity` int(11) DEFAULT NULL COMMENT '数量',
  `formula` varchar(1000) DEFAULT NULL COMMENT '计算公式（配置细项）',
  `price_snapshot` varchar(1000) DEFAULT NULL COMMENT '价格快照',
  `total_fee` decimal(14,2) DEFAULT NULL COMMENT '计算总价（原价，不考虑折扣、优惠）',
  `actual_fee` decimal(14,2) DEFAULT NULL COMMENT '实际总额（如未改价则与计算总价一致，否则不一致）',
  `discount_fee` decimal(14,2) DEFAULT NULL COMMENT '优惠总额',
  `start_time` datetime DEFAULT NULL COMMENT '订单开始计费时间',
  `end_time` datetime DEFAULT NULL COMMENT '订单停止计费时间',
  `favorable_detail_json` varchar(512) DEFAULT NULL COMMENT '优惠明细(按订单分摊后)，json字符串结构（id、value、fee），包括折扣、代金券、优惠券、商品券等促销信息的id、促销额、实际优惠额',
  `voucher_pay` decimal(14,2) DEFAULT '0.00' COMMENT '代金券支付金额',
  `promotion_favorable_item_info` varchar(2000) DEFAULT NULL COMMENT '促销优惠明细',
  `balance_pay` decimal(14,2) DEFAULT NULL COMMENT '余额支付金额(按订单分摊)',
  `money_pay` decimal(14,2) DEFAULT NULL COMMENT '现金支付金额(按订单分摊)',
  `operator_type` tinyint(1) DEFAULT NULL COMMENT '操作者类型：0-用户；1-管理员（运营）；-1 系统',
  `operator_name` varchar(20) DEFAULT NULL COMMENT '操作者名称',
  `status` tinyint(1) DEFAULT NULL COMMENT '订单商品状态：0-创建中，1-成功，2-失败',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(100) DEFAULT NULL COMMENT '备注',
  `unit` varchar(100) DEFAULT NULL COMMENT '计费时长单位：1-小时，2-天，3-月，4-年',
  `refund_fee` decimal(14,2) DEFAULT NULL COMMENT '已退款金额',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_item_number` (`order_item_number`),
  KEY `idx_number_service_code` (`order_number`,`service_code`),
  KEY `idx_item_id_pin` (`item_id`,`pin`)
) ENGINE=InnoDB AUTO_INCREMENT=2681011 DEFAULT CHARSET=utf8 COMMENT='订单商品';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_item`
--

LOCK TABLES `order_item` WRITE;
/*!40000 ALTER TABLE `order_item` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_jdv_info`
--

DROP TABLE IF EXISTS `order_jdv_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_jdv_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `order_number` varchar(50) NOT NULL COMMENT '主单的订单编号',
  `pin` varchar(50) NOT NULL COMMENT '用户pin',
  `jdv_content` varchar(500) NOT NULL COMMENT '来源渠道信息字符串',
  `jdv_hash` varchar(500) DEFAULT NULL COMMENT '域名希值',
  `utm_source` varchar(100) DEFAULT NULL COMMENT '来源域名',
  `utm_campaign` varchar(100) DEFAULT NULL COMMENT '联盟id',
  `utm_media` varchar(100) DEFAULT NULL COMMENT '推广类型',
  `utm_term` varchar(100) DEFAULT NULL COMMENT '推广关键词',
  `rt` varchar(100) DEFAULT NULL COMMENT 'unix时间戳（毫秒）',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `IDX_COMBINED` (`order_number`,`pin`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1338 DEFAULT CHARSET=utf8 COMMENT='渠道来源信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_jdv_info`
--

LOCK TABLES `order_jdv_info` WRITE;
/*!40000 ALTER TABLE `order_jdv_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_jdv_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_receipts`
--

DROP TABLE IF EXISTS `order_receipts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_receipts` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增数据库主键',
  `pin` varchar(50) DEFAULT NULL COMMENT 'pin',
  `order_number` varchar(50) DEFAULT NULL COMMENT '订单编号',
  `trade_type` tinyint(1) DEFAULT NULL COMMENT '交易类型 1充值、2消费',
  `payment_channel` tinyint(1) DEFAULT NULL COMMENT '支付方式:0-余额支付，1-企业网银，2-个人网银，3-微信支付，4-京东支付，5-线下汇款',
  `payment_number` varchar(50) DEFAULT NULL COMMENT '支付单号',
  `amount` decimal(14,2) DEFAULT NULL COMMENT '金额',
  `payment_time` datetime DEFAULT NULL COMMENT '支付时间',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_time` datetime DEFAULT NULL COMMENT '更新时间',
  `funds_type` tinyint(4) DEFAULT NULL COMMENT '资金类型：1：收入 2：支出',
  `user_group` tinyint(4) DEFAULT NULL COMMENT '用户标识 1:其他非云鼎, 2:云鼎',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=39460 DEFAULT CHARSET=utf8 COMMENT='订单实收';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_receipts`
--

LOCK TABLES `order_receipts` WRITE;
/*!40000 ALTER TABLE `order_receipts` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_receipts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_record`
--

DROP TABLE IF EXISTS `order_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_record` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增数据库主键',
  `order_number` varchar(50) DEFAULT NULL COMMENT '订单编号',
  `process_type` tinyint(1) DEFAULT '1' COMMENT '流程类型：1、主站_标准流程_V1.0（当前标准下单接口）2、服务api_预付费下单免支付_V1.0（政务云需求场景，取消对接财务系统）3、服务api_标准下单流程_V1.0（OEM需求场景，取消对接财务系统）',
  `self_support_type` tinyint(1) DEFAULT '1' COMMENT '自营类型：1-自营，2-非自营，3-两者并存',
  `pin` varchar(30) DEFAULT NULL COMMENT '用户pin',
  `site_type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '站点类型 1：主站,2：国际站',
  `order_type` tinyint(1) DEFAULT NULL COMMENT '订单类型：1-新购，2-续费，3-配置变更',
  `charge_mode` tinyint(1) DEFAULT NULL COMMENT '订单计费类型：1-按配置，2-按用量，3-按包年/包月，4-一次性付费',
  `app_code` varchar(50) DEFAULT NULL COMMENT '应用code：domain-域名，0003-数据云，jcloud-基础云，jcloudmark-云市场',
  `item_name` varchar(255) DEFAULT NULL COMMENT '商品名称，用于订单列表显示',
  `status` tinyint(1) DEFAULT NULL COMMENT '订单状态：1-已支付，2-失败，3-未支付，4-处理中，5-已取消，6-部分退款，7-全部退款',
  `currency` tinyint(1) NOT NULL DEFAULT '1' COMMENT '币种 1：人民币,2：美元',
  `pay_type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '付费类型 1：预付费,2：后付费',
  `total_fee` decimal(14,2) DEFAULT NULL COMMENT '订单总额（原价，不考虑折扣、优惠）',
  `actual_fee` decimal(14,2) DEFAULT NULL COMMENT '实际总额（如未改价则与计算总价一致，否则不一致）',
  `is_actual_fee_modified` tinyint(1) NOT NULL DEFAULT '0' COMMENT '应付价格是否被修改过：0-未修改，1-已修改；',
  `discount_fee` decimal(14,2) DEFAULT NULL COMMENT '优惠总额',
  `paid_fee` decimal(14,2) DEFAULT NULL COMMENT '已支付总额',
  `favorable_detail_json` varchar(512) DEFAULT NULL COMMENT '优惠明细，json结构，包括折扣、代金券、优惠券、商品券等促销信息的id、促销额、实际优惠额',
  `voucher_pay` decimal(14,2) DEFAULT '0.00' COMMENT '代金券支付金额',
  `balance_pay` decimal(14,2) DEFAULT NULL COMMENT '余额支付金额',
  `money_pay` decimal(14,2) DEFAULT NULL COMMENT '现金支付金额',
  `payment_channel` tinyint(1) DEFAULT NULL COMMENT '支付方式:0-余额支付，1-企业网银，2-个人网银，3-微信支付，4-京东支付，5-线下汇款',
  `created_time` datetime DEFAULT NULL COMMENT '订单创建时间',
  `updated_time` datetime DEFAULT NULL COMMENT '订单更新时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `origin_order` varchar(50) DEFAULT NULL COMMENT '原订单编号，针对升降配、续费',
  `source_id` varchar(50) DEFAULT NULL COMMENT '来源标识',
  `task_id` varchar(40) DEFAULT NULL COMMENT '任务id：对应于中间层任务id，打包购买时的一个事务Id',
  `op_type` tinyint(1) DEFAULT NULL COMMENT '操作类型：0-订单提交，1-取消订单，2-支付完成，3-订单失败，4-生成完成，6-退款申请，7-退款完成',
  `operator_type` tinyint(1) DEFAULT NULL COMMENT '操作者类型：0-用户，1-管理员（运营），-1 系统',
  `operator_name` varchar(20) DEFAULT NULL COMMENT '操作者名称',
  `ip` varchar(200) DEFAULT NULL COMMENT 'ip地址',
  `device_info` varchar(2000) DEFAULT NULL COMMENT '设备信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1880153 DEFAULT CHARSET=utf8 COMMENT='订单记录';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_record`
--

LOCK TABLES `order_record` WRITE;
/*!40000 ALTER TABLE `order_record` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_record` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_record_2017`
--

DROP TABLE IF EXISTS `order_record_2017`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_record_2017` (
  `id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '自增数据库主键',
  `order_number` varchar(50) DEFAULT NULL COMMENT '订单编号',
  `process_type` tinyint(1) DEFAULT '1' COMMENT '流程类型：1、主站_标准流程_V1.0（当前标准下单接口）2、服务api_预付费下单免支付_V1.0（政务云需求场景，取消对接财务系统）3、服务api_标准下单流程_V1.0（OEM需求场景，取消对接财务系统）',
  `self_support_type` tinyint(1) DEFAULT '1' COMMENT '自营类型：1-自营，2-非自营，3-两者并存',
  `pin` varchar(30) DEFAULT NULL COMMENT '用户pin',
  `site_type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '站点类型 1：主站,2：国际站',
  `order_type` tinyint(1) DEFAULT NULL COMMENT '订单类型：1-新购，2-续费，3-配置变更',
  `charge_mode` tinyint(1) DEFAULT NULL COMMENT '订单计费类型：1-按配置，2-按用量，3-按包年/包月，4-一次性付费',
  `app_code` varchar(50) DEFAULT NULL COMMENT '应用code：domain-域名，0003-数据云，jcloud-基础云，jcloudmark-云市场',
  `item_name` varchar(255) DEFAULT NULL COMMENT '商品名称，用于订单列表显示',
  `status` tinyint(1) DEFAULT NULL COMMENT '订单状态：1-已支付，2-失败，3-未支付，4-处理中，5-已取消，6-部分退款，7-全部退款',
  `currency` tinyint(1) NOT NULL DEFAULT '1' COMMENT '币种 1：人民币,2：美元',
  `pay_type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '付费类型 1：预付费,2：后付费',
  `total_fee` decimal(14,2) DEFAULT NULL COMMENT '订单总额（原价，不考虑折扣、优惠）',
  `actual_fee` decimal(14,2) DEFAULT NULL COMMENT '实际总额（如未改价则与计算总价一致，否则不一致）',
  `is_actual_fee_modified` tinyint(1) NOT NULL DEFAULT '0' COMMENT '应付价格是否被修改过：0-未修改，1-已修改；',
  `discount_fee` decimal(14,2) DEFAULT NULL COMMENT '优惠总额',
  `paid_fee` decimal(14,2) DEFAULT NULL COMMENT '已支付总额',
  `favorable_detail_json` varchar(512) DEFAULT NULL COMMENT '优惠明细，json结构，包括折扣、代金券、优惠券、商品券等促销信息的id、促销额、实际优惠额',
  `balance_pay` decimal(14,2) DEFAULT NULL COMMENT '余额支付金额',
  `money_pay` decimal(14,2) DEFAULT NULL COMMENT '现金支付金额',
  `payment_channel` tinyint(1) DEFAULT NULL COMMENT '支付方式:0-余额支付，1-企业网银，2-个人网银，3-微信支付，4-京东支付，5-线下汇款',
  `created_time` datetime DEFAULT NULL COMMENT '订单创建时间',
  `updated_time` datetime DEFAULT NULL COMMENT '订单更新时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `origin_order` varchar(50) DEFAULT NULL COMMENT '原订单编号，针对升降配、续费',
  `source_id` varchar(50) DEFAULT NULL COMMENT '来源标识',
  `task_id` varchar(40) DEFAULT NULL COMMENT '任务id：对应于中间层任务id，打包购买时的一个事务Id',
  `op_type` tinyint(1) DEFAULT NULL COMMENT '操作类型：0-订单提交，1-取消订单，2-支付完成，3-订单失败，4-生成完成，6-退款申请，7-退款完成',
  `operator_type` tinyint(1) DEFAULT NULL COMMENT '操作者类型：0-用户，1-管理员（运营），-1 系统',
  `operator_name` varchar(20) DEFAULT NULL COMMENT '操作者名称',
  `ip` varchar(200) DEFAULT NULL COMMENT 'ip地址',
  `device_info` varchar(2000) DEFAULT NULL COMMENT '设备信息'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_record_2017`
--

LOCK TABLES `order_record_2017` WRITE;
/*!40000 ALTER TABLE `order_record_2017` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_record_2017` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_signature`
--

DROP TABLE IF EXISTS `order_signature`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_signature` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `process_type` int(4) NOT NULL COMMENT '流程类型',
  `site_type` int(4) NOT NULL COMMENT '站点类型',
  `app_code` varchar(50) NOT NULL COMMENT '业务线',
  `service_code` varchar(50) DEFAULT NULL COMMENT '产品code',
  `test_sign_key` varchar(100) NOT NULL COMMENT '开发测试签名key',
  `online_sign_key` varchar(100) NOT NULL COMMENT '线上签名key',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态：1-在线，2-下线',
  `created_time` datetime DEFAULT NULL,
  `updated_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `business_type` (`process_type`,`site_type`,`app_code`,`service_code`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_signature`
--

LOCK TABLES `order_signature` WRITE;
/*!40000 ALTER TABLE `order_signature` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_signature` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_snapshot`
--

DROP TABLE IF EXISTS `order_snapshot`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_snapshot` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增数据库主键',
  `batch_id` varchar(100) NOT NULL COMMENT '快照批次id',
  `pin` varchar(50) NOT NULL COMMENT '用户pin',
  `snapshot` varchar(2000) NOT NULL COMMENT '订单快照',
  `ip` varchar(100) DEFAULT NULL COMMENT 'ip',
  `device` varchar(2000) DEFAULT NULL COMMENT '设备信息',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=588307 DEFAULT CHARSET=utf8 COMMENT='记录订单快照表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_snapshot`
--

LOCK TABLES `order_snapshot` WRITE;
/*!40000 ALTER TABLE `order_snapshot` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_snapshot` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_temp`
--

DROP TABLE IF EXISTS `order_temp`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_temp` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `batch_id` varchar(100) NOT NULL COMMENT '批次id',
  `pin` varchar(50) NOT NULL COMMENT '用户pin',
  `snapshot` varchar(2000) NOT NULL COMMENT '订单信息json串',
  `sequence_no` int(10) NOT NULL COMMENT '订单信息序号',
  `ip` varchar(100) DEFAULT NULL COMMENT 'ip',
  `device` varchar(2000) DEFAULT NULL COMMENT '设备信息',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  `unique_code` varchar(100) DEFAULT NULL COMMENT 'MD5加密串',
  PRIMARY KEY (`id`),
  KEY `IDX_BATCHID` (`batch_id`) USING BTREE,
  KEY `IDX_COMBINED` (`batch_id`,`pin`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=41471 DEFAULT CHARSET=utf8 COMMENT='临时订单表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_temp`
--

LOCK TABLES `order_temp` WRITE;
/*!40000 ALTER TABLE `order_temp` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_temp` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payment_journal`
--

DROP TABLE IF EXISTS `payment_journal`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `payment_journal` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `pin` varchar(50) DEFAULT NULL COMMENT 'pin',
  `order_number` varchar(50) NOT NULL COMMENT '订单号',
  `trade_type` tinyint(1) DEFAULT NULL COMMENT '交易类型 1充值、2消费',
  `payment_number` varchar(50) NOT NULL COMMENT '支付单号',
  `amount` decimal(14,2) NOT NULL COMMENT '订单金额',
  `payment_channel` tinyint(1) DEFAULT NULL COMMENT '支付方式:0-余额支付，1-企业网银，2-个人网银，3-微信支付，4-京东支付，5-线下汇款',
  `seller_account` varchar(255) DEFAULT NULL COMMENT '卖家支付帐号',
  `bank_code` varchar(255) DEFAULT NULL COMMENT '银行code',
  `remittance_code` varchar(255) DEFAULT NULL COMMENT '汇款识别码',
  `app_code` varchar(255) DEFAULT NULL COMMENT '应用code',
  `service_code` varchar(255) DEFAULT NULL COMMENT '服务code',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_payment_number` (`payment_number`) USING BTREE,
  KEY `idx_order_number` (`order_number`)
) ENGINE=InnoDB AUTO_INCREMENT=59508 DEFAULT CHARSET=utf8 COMMENT='支付流水';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payment_journal`
--

LOCK TABLES `payment_journal` WRITE;
/*!40000 ALTER TABLE `payment_journal` DISABLE KEYS */;
/*!40000 ALTER TABLE `payment_journal` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payment_result`
--

DROP TABLE IF EXISTS `payment_result`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `payment_result` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `pin` varchar(50) DEFAULT NULL COMMENT 'pin',
  `payment_number` varchar(50) NOT NULL COMMENT '本系统的订单编号',
  `payment_channel` tinyint(1) DEFAULT NULL COMMENT '支付方式:0-余额支付，1-企业网银，2-个人网银，3-微信支付，4-京东支付，5-线下汇款',
  `trade_type` tinyint(1) DEFAULT NULL COMMENT '交易类型 1充值、2消费',
  `trade_status` varchar(20) DEFAULT NULL COMMENT '交易状态',
  `trade_number` varchar(64) DEFAULT NULL COMMENT '支付渠道的交易号',
  `subject` varchar(256) DEFAULT NULL COMMENT '商品名称',
  `total_fee` decimal(14,2) NOT NULL COMMENT '支付金额',
  `fee_type` tinyint(1) DEFAULT NULL COMMENT '币种：1-人民币，2-美元',
  `bank_code` varchar(50) DEFAULT NULL COMMENT '付款银行',
  `payment_time` datetime DEFAULT NULL COMMENT '支付完成时间',
  `sent_account_status` tinyint(1) DEFAULT NULL COMMENT '是否已发送对账通知：1-未发送，2-已发送，3-发送失败',
  `full_text` varchar(4000) DEFAULT NULL COMMENT '结果通知全文',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_payment_number` (`payment_number`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=38588 DEFAULT CHARSET=utf8 COMMENT='支付结果';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payment_result`
--

LOCK TABLES `payment_result` WRITE;
/*!40000 ALTER TABLE `payment_result` DISABLE KEYS */;
/*!40000 ALTER TABLE `payment_result` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `refund_order`
--

DROP TABLE IF EXISTS `refund_order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `refund_order` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增数据库主键',
  `order_number` varchar(50) DEFAULT NULL COMMENT '订单编号',
  `order_item_number` varchar(50) DEFAULT NULL COMMENT '订单商品编号',
  `pin` varchar(30) DEFAULT NULL COMMENT '用户pin',
  `refund_type` tinyint(1) DEFAULT NULL COMMENT '退款类型：1-用户退款，2-系统退款',
  `total_fee` decimal(14,2) DEFAULT NULL COMMENT '退款总额',
  `favorable_detail_json` varchar(512) DEFAULT NULL COMMENT '促销退款明细，json结构，包括代金券、商品券等促销信息的id、促销额、退款金额',
  `balance_fee` decimal(14,2) DEFAULT NULL COMMENT '余额退款金额',
  `money_fee` decimal(14,2) DEFAULT NULL COMMENT '现金退款金额',
  `payment_channel` tinyint(1) DEFAULT NULL COMMENT '支付方式:0-余额支付，1-企业网银，2-个人网银，3-微信支付，4-京东支付，5-线下汇款',
  `status` tinyint(1) DEFAULT NULL COMMENT '退款状态：0-申请退款，1-退款审核中，2-已退款，3-已取消，4-退款失败，5-财务审核驳回，6-财务审批通过',
  `created_time` datetime DEFAULT NULL COMMENT '订单创建时间',
  `updated_time` datetime DEFAULT NULL COMMENT '订单更新时间',
  `remark` varchar(2000) DEFAULT NULL COMMENT '备注',
  `actual_order_number` varchar(2000) DEFAULT NULL COMMENT '实际退款订单号',
  `ref_id` varchar(50) DEFAULT NULL COMMENT 'UC退款业务ID',
  `order_type` tinyint(1) DEFAULT NULL COMMENT '1:余额提现，2:订单退款',
  `erp_order_id` varchar(50) DEFAULT NULL COMMENT 'uc充值单号',
  `refund_route` tinyint(1) DEFAULT NULL COMMENT '余额提现为空  1、原路退 2、退款到余额',
  PRIMARY KEY (`id`),
  KEY `idx_number_ref` (`order_number`,`ref_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6374 DEFAULT CHARSET=utf8 COMMENT='退款单';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `refund_order`
--

LOCK TABLES `refund_order` WRITE;
/*!40000 ALTER TABLE `refund_order` DISABLE KEYS */;
/*!40000 ALTER TABLE `refund_order` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `refund_order_detail`
--

DROP TABLE IF EXISTS `refund_order_detail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `refund_order_detail` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `order_number` varchar(50) NOT NULL COMMENT '订单号',
  `refund_order_id` bigint(20) NOT NULL COMMENT '退款订单号',
  `sub_order_number` varchar(50) DEFAULT NULL COMMENT '子订单号',
  `item_id` varchar(50) DEFAULT NULL COMMENT '商品id',
  `total_fee` decimal(14,2) NOT NULL COMMENT '退款总金额',
  `balance_fee` decimal(14,2) DEFAULT NULL COMMENT '余额退款金额',
  `money_fee` decimal(14,2) DEFAULT NULL COMMENT '现金退款金额',
  `status` tinyint(1) DEFAULT NULL COMMENT '退款状态',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `creator` varchar(50) NOT NULL COMMENT '创建人',
  `updated_time` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=493 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `refund_order_detail`
--

LOCK TABLES `refund_order_detail` WRITE;
/*!40000 ALTER TABLE `refund_order_detail` DISABLE KEYS */;
/*!40000 ALTER TABLE `refund_order_detail` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `refund_order_ots`
--

DROP TABLE IF EXISTS `refund_order_ots`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `refund_order_ots` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增数据库主键',
  `order_number` varchar(50) NOT NULL COMMENT '订单号',
  `refund_fee` decimal(14,2) NOT NULL COMMENT '退款金额',
  `status` tinyint(1) NOT NULL COMMENT '状态0:受理状态;-1:错误状态 ;2:成功状态 1.失败状态',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `refund_apply_id` varchar(100) DEFAULT NULL COMMENT '台账退款业务id',
  `order_type` tinyint(1) DEFAULT NULL COMMENT '订单类型，1：余额提现，2：订单退款',
  `info` varchar(100) DEFAULT NULL COMMENT '错误信息',
  `payment_channel` tinyint(2) DEFAULT NULL COMMENT '支付渠道',
  `pin` varchar(30) DEFAULT NULL COMMENT '用户pin',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_refund_apply_id` (`refund_apply_id`),
  KEY `idx_payment_channel` (`payment_channel`)
) ENGINE=InnoDB AUTO_INCREMENT=1062 DEFAULT CHARSET=utf8 COMMENT='退款ots表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `refund_order_ots`
--

LOCK TABLES `refund_order_ots` WRITE;
/*!40000 ALTER TABLE `refund_order_ots` DISABLE KEYS */;
/*!40000 ALTER TABLE `refund_order_ots` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `reverse_sterilisation`
--

DROP TABLE IF EXISTS `reverse_sterilisation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `reverse_sterilisation` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增数据库主键',
  `pin` varchar(30) DEFAULT NULL COMMENT '用户pin',
  `order_number` varchar(50) DEFAULT NULL COMMENT '充值单ID',
  `payment_number` varchar(50) DEFAULT NULL COMMENT '支付单ID',
  `refund_number` varchar(50) DEFAULT NULL COMMENT '退款单ID',
  `amount` decimal(14,2) DEFAULT NULL COMMENT '金额',
  `status` tinyint(1) DEFAULT NULL COMMENT '冲销状态 1冲销成功 2 余额不足冲销失败',
  `platform_account_status` tinyint(1) DEFAULT NULL COMMENT '台账状态  1成功 2失败',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_time` datetime DEFAULT NULL COMMENT '修改时间',
  `remark` varchar(50) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='逆向冲销';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reverse_sterilisation`
--

LOCK TABLES `reverse_sterilisation` WRITE;
/*!40000 ALTER TABLE `reverse_sterilisation` DISABLE KEYS */;
/*!40000 ALTER TABLE `reverse_sterilisation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `settlement_detail`
--

DROP TABLE IF EXISTS `settlement_detail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `settlement_detail` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增数据库主键',
  `pin` varchar(50) DEFAULT NULL COMMENT 'pin',
  `order_number` varchar(50) DEFAULT NULL COMMENT '订单编号',
  `trade_type` tinyint(1) DEFAULT NULL COMMENT '交易类型 1充值、2消费',
  `payment_channel` tinyint(1) DEFAULT NULL COMMENT '支付方式:0-余额支付，1-企业网银，2-个人网银，3-微信支付，4-京东支付，5-线下汇款',
  `payment_number` varchar(255) DEFAULT NULL COMMENT '支付单号',
  `amount` decimal(14,2) DEFAULT NULL COMMENT '订单金额',
  `receivable` decimal(14,2) DEFAULT NULL COMMENT '应付金额',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_time` datetime DEFAULT NULL COMMENT '更新时间',
  `status` tinyint(1) DEFAULT '1' COMMENT '标识对账状态  1、成功  2、失败',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=38402 DEFAULT CHARSET=utf8 COMMENT='对账凭证';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `settlement_detail`
--

LOCK TABLES `settlement_detail` WRITE;
/*!40000 ALTER TABLE `settlement_detail` DISABLE KEYS */;
/*!40000 ALTER TABLE `settlement_detail` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_config`
--

DROP TABLE IF EXISTS `sys_config`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_config` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `key` varchar(100) NOT NULL COMMENT '键',
  `value` varchar(2000) NOT NULL COMMENT '值',
  `sys` varchar(100) DEFAULT NULL COMMENT '系统',
  `type` varchar(100) NOT NULL COMMENT '类别',
  `description` varchar(2000) DEFAULT NULL COMMENT '描述',
  `created_time` datetime NOT NULL COMMENT '创建时间',
  `updated_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COMMENT='系统配置表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_config`
--

LOCK TABLES `sys_config` WRITE;
/*!40000 ALTER TABLE `sys_config` DISABLE KEYS */;
INSERT INTO `sys_config` VALUES (1,'cancel_interval','15','jcloud_order','order_auto_cancel','订单自动取消时间间隔','2017-09-19 14:45:28','2017-09-19 14:45:32'),(2,'cancel_switch','1','jcloud_order','order_auto_cancel','订单自动取消开关，0为关闭，1为开启','2017-09-21 10:32:16','2017-09-21 10:49:42'),(3,'quota_switch','1','jcloud_order','order_check_quota','配额：0为关闭，1为开启','2017-09-21 10:32:16','2017-09-21 10:32:16'),(4,'signature_time','300000','jcloud_order','check_signature_time','签名时间戳有效时间，单位是ms','2017-09-21 10:32:16','2017-09-21 15:45:37'),(5,'pin','boss-01,JaneCloud002,JaneCloud003,jd_pengchang,JaneCloud007,boss-05,JaneCloud008,dujuangrace,jessesky,黑莓cc,jd_57b822db30d1c,陈令奎,bjhehuan,beesharp,skty,125752875-47861837,datajingdo_m,jd_411f536a7746e,R_gabriel,sunfire_999,JD电商云,jos-api,apatriot,auhiug,tx4808,laodang1990,yuanshuaimingtest,wangchaoyang27,lizhenzhen2015,jcloud2016,gingindini,sunimoon,jd_430e83f009d83,sqjiankong','jcloud_order','order_test_pin','测试人员的pin','2017-10-24 18:32:27','2017-10-24 18:32:27'),(6,'pin','JaneCloud005,JaneCloud004,数据云AGS,1439195762-738219,jd_4af74c4322d89,choui,仇俊宝,quning_jd,hello_yunhai,排骨虾,jd_7bf1c61b1973d,tianlei2939,jcloud_001,dataJcloudTest1,jcloud_005,pmo_cloud,xdata_ml,huyp891,jcloudiaas2,jcloudiaas3,jcloudiaas4,jcloudiaas5,商业合作部测试账号,sunbo9,yuanyuanzhang18,jcloud_iaas,区块链项目测试,Govcloud,2017徐小白来了,miamiamia1986,inf_dev,ringcjs','jcloud_order','order_test_pin','测试人员的pin','2017-10-24 18:32:27','2017-10-24 18:32:27'),(7,'pin','电销测试账号,jd_cdntest,jccstest_dx,jccstest_kf,iaas-ops,qwertyuiop8787,jcloud_ami,jd_4db1955cf8390,jd_数据库测试_00A,流岚鱼,sec365buy,wdtxliqf,jcloud_025,xingyixun,jd_18e0c2452e2c5,jd_7f0131d8b9767,jd_67668e1a8e701,jd_8669e5c5f7809,lookatmebay1,areatest,xueliang001,15110055080_p,satsuki_li,zhuhc999,pid-plat,pmo_jd_a,bz_jcloud_test,PMO_JCLOUD,jcloudhillhouse,云安全外测,jcloud-b2b-04','jcloud_order','order_test_pin','测试人员的pin','2017-10-24 18:32:27','2017-10-24 18:32:27'),(8,'pin','jcloud-b2c-03,liuxws,jd_52ca057742db6,jd_saas_pre,刘刘看看,jcloud_security,coolcole,changchuan1,JDCSolution,network_test','jcloud_order','order_test_pin','测试人员的pin','2017-10-24 18:32:27','2017-10-24 18:32:27'),(9,'notice.user.week','guojing13,zhenglei6,cuienduo,zhaojunqiang,liweijun3,yuweiwei1','jcloud_order','ews.mail','周报统计人员名单','2017-12-26 14:37:04','2017-12-26 14:37:04'),(10,'hdsh.voucherInfo.user','huangwanyue,lihuirong,zhaoyu7,daijinquan,liujia51,houchao,yuweiwei1,shiyening,cuienduo,zhaojunqiang,zhenglei6','jcloud_order','ews.mail','华东-上海代金券信息统计接收人','2017-12-27 14:41:21','2017-12-27 14:41:21');
/*!40000 ALTER TABLE `sys_config` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-08-02 13:37:25
