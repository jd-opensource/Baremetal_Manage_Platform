-- MySQL dump 10.14  Distrib 5.5.68-MariaDB, for Linux (x86_64)
--
-- Host: 10.208.12.72    Database: bmp
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
-- Table structure for table `apikey`
--

DROP TABLE IF EXISTS `apikey`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `apikey` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `apikey_id` varchar(255) NOT NULL DEFAULT '' COMMENT '秘钥对UUID',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '秘钥对名称',
  `read_only` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否支持只读，read_only =1 的时候说明这个key是只读key，不能访问写方法。',
  `token` varchar(255) NOT NULL DEFAULT '' COMMENT '32位字符令牌',
  `type` varchar(10) NOT NULL DEFAULT '' COMMENT 'Token类型system/user',
  `user_id` varchar(255) NOT NULL DEFAULT '' COMMENT '所属用户',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新者',
  `created_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间戳',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  `source` varchar(255) NOT NULL DEFAULT '' COMMENT 'create from console/operation',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=90 DEFAULT CHARSET=utf8 COMMENT='api秘钥';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `apikey`
--

LOCK TABLES `apikey` WRITE;
/*!40000 ALTER TABLE `apikey` DISABLE KEYS */;
INSERT INTO `apikey` VALUES (1,'apikey-xdghjr4cvikabs5g95mejh24fx5d','default-console-api',0,'Ym1wLWNvbnNvbGUtYXBp','system','','admin','admin',1667274361,1667274361,0,0,''),(2,'apikey-23df9fjs9dxc9dchhd925vhlnusg','default-operation-api',0,'Ym1wLW9wZXJhdGlvbi1hcGk=','system','','admin','admin',1667274361,1667274361,0,0,''),(3,'apikey-sinojr4cvicabs5g95mejh24v587','default-alert-api',0,'Ac13LWNvdrpxkGUtxxfcp','alert','','admin','admin',1667274361,1667274361,0,0,'');
/*!40000 ALTER TABLE `apikey` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `audit_logs`
--

DROP TABLE IF EXISTS `audit_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `audit_logs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `logid` varchar(255) NOT NULL DEFAULT '' COMMENT 'log uuid',
  `sn` varchar(255) NOT NULL DEFAULT '',
  `device_id` varchar(255) NOT NULL DEFAULT '',
  `instance_id` varchar(255) NOT NULL DEFAULT '',
  `operation` varchar(255) NOT NULL DEFAULT '' COMMENT 'action',
  `operate_user_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'user_id',
  `operate_user_name` varchar(255) NOT NULL DEFAULT '' COMMENT 'user_name',
  `result` varchar(255) NOT NULL DEFAULT '' COMMENT 'success/fail',
  `fail_reason` varchar(255) NOT NULL DEFAULT '' COMMENT 'reason for fail',
  `created_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `i_sn` (`sn`),
  KEY `i_device_id` (`device_id`),
  KEY `i_instance_id` (`instance_id`),
  KEY `i_operation` (`operation`)
) ENGINE=InnoDB AUTO_INCREMENT=105921 DEFAULT CHARSET=utf8 COMMENT='audit_logs';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `audit_logs`
--

LOCK TABLES `audit_logs` WRITE;
/*!40000 ALTER TABLE `audit_logs` DISABLE KEYS */;
/*!40000 ALTER TABLE `audit_logs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `collect`
--

DROP TABLE IF EXISTS `collect`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `collect` (
  `id` int(32) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `collect_method` tinyint(32) unsigned NOT NULL,
  `kernel_name` varchar(255) NOT NULL DEFAULT '' COMMENT 'kernel_name',
  `initramfs_name` varchar(255) NOT NULL DEFAULT '' COMMENT 'initramfs_name',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `rule_id_idx` (`collect_method`),
  KEY `sn_idx` (`kernel_name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `collect`
--

LOCK TABLES `collect` WRITE;
/*!40000 ALTER TABLE `collect` DISABLE KEYS */;
INSERT INTO `collect` VALUES (1,1,'v2.0.4-centos_7_9-2024080513-vmlinuz','v2.0.4-centos_7_9-2024080513-initramfs.gz',0),(2,2,'v2.0.7-ubuntu_2204-2024082914-vmlinuz','v2.0.7-ubuntu_2204-2024082914-initramfs.gz',0);
/*!40000 ALTER TABLE `collect` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `command`
--

DROP TABLE IF EXISTS `command`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `command` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `request_id` varchar(64) NOT NULL DEFAULT '' COMMENT '请求ID',
  `sn` varchar(64) NOT NULL DEFAULT '' COMMENT '设备SN',
  `instance_id` varchar(36) NOT NULL DEFAULT '' COMMENT '实例Id',
  `action` varchar(32) NOT NULL DEFAULT '' COMMENT '操作',
  `type` varchar(16) NOT NULL DEFAULT '' COMMENT '操作类型：agent, driver, network',
  `status` varchar(16) NOT NULL DEFAULT '' COMMENT '状态: wait,running,finish,error',
  `parent_command_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '父指令Id',
  `execute_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '执行次数',
  `timeout_time` datetime DEFAULT NULL COMMENT 'timeout time',
  `timeout_policy` varchar(16) NOT NULL DEFAULT '' COMMENT 'timeout policy',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新者',
  `created_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间戳',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  `task` varchar(128) DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `i_command_request_id` (`request_id`) USING BTREE,
  KEY `i_command_sn` (`sn`) USING BTREE,
  KEY `i_command_instance_id` (`instance_id`) USING BTREE,
  KEY `i_command_status` (`status`) USING BTREE,
  KEY `i_parent_command_id` (`parent_command_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3503 DEFAULT CHARSET=utf8 COMMENT='指令';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `command`
--

LOCK TABLES `command` WRITE;
/*!40000 ALTER TABLE `command` DISABLE KEYS */;
/*!40000 ALTER TABLE `command` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cps_fault_rules`
--

DROP TABLE IF EXISTS `cps_fault_rules`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cps_fault_rules` (
  `id` int(32) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(52) NOT NULL DEFAULT '' COMMENT '故障名称',
  `parts` varchar(12) NOT NULL DEFAULT '' COMMENT '故障配件',
  `fault_type` varchar(10) NOT NULL DEFAULT '' COMMENT '故障类型',
  `condition` varchar(20) NOT NULL DEFAULT '' COMMENT '判定条件',
  `threshold` varchar(52) NOT NULL DEFAULT '' COMMENT '判定阀值',
  `level` varchar(16) NOT NULL DEFAULT '' COMMENT '故障等级',
  `desc` varchar(152) NOT NULL DEFAULT '' COMMENT '故障描述',
  `is_push` tinyint(1) NOT NULL DEFAULT '0' COMMENT '故障是否推送 0-否 1-是',
  `is_use` tinyint(1) NOT NULL DEFAULT '0' COMMENT '规则是否启动 0-否 1-是',
  `is_default_push` tinyint(1) NOT NULL DEFAULT '0' COMMENT '故障是否默认推送 0-否 1-是',
  `is_default_use` tinyint(1) NOT NULL DEFAULT '0' COMMENT '规则是否默认启动 0-否 1-是',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除 0-否 1-是',
  `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `owned` varchar(12) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `is_push_idx` (`is_push`),
  KEY `is_use_idx` (`is_use`),
  KEY `parts_idx` (`parts`),
  KEY `owned_idx` (`owned`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cps_fault_rules`
--

LOCK TABLES `cps_fault_rules` WRITE;
/*!40000 ALTER TABLE `cps_fault_rules` DISABLE KEYS */;
INSERT INTO `cps_fault_rules` VALUES (1,'power-supply-failure-detected','电源','电力故障','Power Supply','Failure detected','Critical','检测到电源故障(power-supply-failure-detected)',1,1,1,1,0,'2022-12-02 03:30:49','power'),(2,'power-supply-predictive-failure','电源','电力故障','Power Supply','Predictive failure','Warning','预期电源供电即将异常(power-supply-predictive-failure)',1,1,0,1,0,'2022-12-02 03:30:49','power'),(3,'power-supply-ac-lost','电源','电力故障','Power Supply','Power Supply AC lost','Critical','电源动力丢失(power-supply-ac-lost)',1,1,1,1,0,'2022-12-02 03:30:49','power'),(4,'power-supply-redundant-lost','电源','电力故障','Power Supply','Redundant Lost','Critical','电源冗余丢失(power-supply-redundant-lost',1,1,1,1,0,'2022-12-02 03:30:49','power'),(5,'temperature-upper-non-critical','温度','温控故障','Temperature','Upper Non-Critical','Warning','非严重级别温度过高(temperature-upper-non-critical)',1,1,0,1,0,'2022-12-02 03:30:49','other'),(6,'temperature-upper-critical','温度','温控故障','Temperature','Upper Critical','Critical','严重级别温度过高(temperature-upper-critical)',1,1,1,1,0,'2022-12-02 03:30:49','other'),(7,'processor-ierr','CPU','CPU故障','Processor','IERR','Critical','CPU内部错误(processor-ierr)',1,1,1,1,0,'2022-12-02 03:30:49','cpu'),(8,'processor-thermal-trip','CPU','CPU故障','Processor','Thermal Trip','Critical','CPU过热强制关机(processor-thermal-trip)',1,1,1,1,0,'2022-12-02 03:30:49','cpu'),(9,'processor-configuration-error','CPU','CPU故障','Processor','Configuration Error','Critical','CPU配置错误(processor-configuration-error)',1,1,1,1,0,'2022-12-02 03:30:49','cpu'),(10,'processor-uncorrectable-machine-check','CPU','CPU故障','Processor','Uncorrectable machine check exception','Critical','CPU不可纠正的自检故障(processor-uncorrectable-machine-check)',1,1,1,1,0,'2022-12-02 03:30:49','cpu'),(11,'processor-correctable-machine-check','CPU','CPU故障','Processor','Correctable machine check error','Warning','CPU可纠正的自检故障(processor-correctable-machine-check)',0,1,0,1,0,'2022-12-02 03:30:49','cpu'),(12,'memory-uncorrectable-ecc','内存','内存故障','Memory','Uncorrectable ECC','Critical','内存不可纠正的ECC故障(memory-uncorrectable-ecc)',1,1,1,1,0,'2022-12-02 03:30:49','mem'),(13,'memory-correctable-ecc','内存','内存故障','Memory','Correctable ECC','Warning','内存可纠正的ECC故障(memory-correctable-ecc)',0,1,0,1,0,'2022-12-02 03:30:49','mem'),(14,'memory-configuration-error','内存','内存故障','Memory','Configuration Error','Critical','内存配置错误(memory-configuration-error)',1,1,1,1,0,'2022-12-02 03:30:49','mem'),(15,'memory-critical-overtemperature','内存','内存故障','Memory','Critical Overtemperature','Critical','严重级别内存温度过高(memory-critical-overtemperature)',1,1,1,1,0,'2022-12-02 03:30:49','mem'),(16,'memory-device-disabled','内存','内存故障','Memory','Memory Device Disabled','Critical','内存设备禁用故障(memory-device-disabled)',1,1,1,1,0,'2022-12-02 03:30:49','mem'),(17,'disk-predictive-failure','硬盘','硬盘故障','Drive Slot / Bay','Predictive Failure','Warning','预期磁盘即将故障(disk-predictive-failure)',0,1,0,1,0,'2022-12-02 03:30:49','disk'),(18,'disk-drive-fault','硬盘','硬盘故障','Drive Slot / Bay','Drive Fault','Critical','磁盘驱动器故障(disk-drive-fault)',0,1,1,1,0,'2022-12-02 03:30:49','disk'),(19,'disk-in-critical-array','硬盘','硬盘故障','Drive Slot / Bay','In Critical Array','Critical','磁盘所在阵列严重级别故障(disk-in-critical-array)',0,1,1,1,0,'2022-12-02 03:30:49','disk'),(20,'disk-in-failed-array','硬盘','硬盘故障','Drive Slot / Bay','In Failed Array','Critical','磁盘所在阵列失效(disk-in-failed-array)',0,1,1,1,0,'2022-12-02 03:30:49','disk'),(21,'voltage-upper-non-critical','电压','电力故障','Voltage','Upper Non-Critical','Warning','非严重级别电压过高(voltage-upper-non-critical)',0,1,0,1,0,'2022-12-02 03:30:49','power'),(22,'voltage-upper-critical','电压','电力故障','Voltage','Upper Critical','Critical','严重级别电压过高(voltage-upper-critical)',1,1,1,1,0,'2022-12-02 03:30:49','power'),(23,'voltage-upper-non-recoverable','电压','电力故障','Voltage','Upper Non-Recoverable','Warning','电压不可恢复过高(voltage-upper-non-recoverable)',0,1,0,1,0,'2022-12-02 03:30:49','power'),(24,'voltage-lower-non-critical','电压','电力故障','Voltage','Lower Non-Critical','Warning','非严重级别电压过低(voltage-lower-non-critical)',0,1,0,1,0,'2022-12-02 03:30:49','power'),(25,'voltage-lower-critical','电压','电力故障','Voltage','Lower Critical','Critical','严重级别电压过低(voltage-lower-critical)',1,1,1,1,0,'2022-12-02 03:30:49','power'),(26,'voltage-lower-non-recoverable','电压','电力故障','Voltage','Lower Non-Recoverable','Warning','电压不可恢复过低(voltage-lower-non-recoverable)',0,1,0,1,0,'2022-12-02 03:30:49','power'),(27,'fan-lower-critical','风扇','温控故障','Fan','Lower Critical','Critical','严重级别风扇转速过低(fan-lower-critical)',1,1,1,1,0,'2022-12-02 03:30:49','other'),(28,'pcie-bus-fatal-error','PCIE总线','其他故障','Critical Interrupt','Bus Fatal Error','Critical','PCI 总线致命错误(pcie-bus-fatal-error)',1,1,1,1,0,'2022-12-02 03:30:49','other'),(29,'disk-drive-present-deassert','硬盘','硬盘故障','Drive Slot / Bay','Drive Present () | Deasserted','Critical','磁盘驱动器不存在(disk-drive-present-deassert)',0,1,1,1,0,'2022-12-02 03:30:49','disk'),(31,'disk-drive-present-deassert','硬盘','disk error','Drive','Drive','Critical','disk',0,1,1,1,1,'2022-12-02 03:30:49','disk');
/*!40000 ALTER TABLE `cps_fault_rules` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cps_log_content_collection`
--

DROP TABLE IF EXISTS `cps_log_content_collection`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cps_log_content_collection` (
  `id` int(32) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `content` varchar(256) NOT NULL DEFAULT '',
  `fault_conf_id` int(32) unsigned NOT NULL,
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT ': 0,; 1,',
  `sn` varchar(32) NOT NULL DEFAULT '' COMMENT 'sn',
  `count` int(32) unsigned NOT NULL COMMENT 'count',
  `level` varchar(16) NOT NULL DEFAULT '',
  `event_time` timestamp NULL DEFAULT NULL,
  `collect_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_isdel` (`is_del`),
  KEY `rule_id_idx` (`fault_conf_id`),
  KEY `sn_idx` (`sn`),
  KEY `level_idx` (`level`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cps_log_content_collection`
--

LOCK TABLES `cps_log_content_collection` WRITE;
/*!40000 ALTER TABLE `cps_log_content_collection` DISABLE KEYS */;
/*!40000 ALTER TABLE `cps_log_content_collection` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cps_log_items`
--

DROP TABLE IF EXISTS `cps_log_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cps_log_items` (
  `id` int(32) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `collection_id` int(32) unsigned NOT NULL COMMENT 'id',
  `collect_time` timestamp NULL DEFAULT NULL,
  `event_time` timestamp NULL DEFAULT NULL,
  `is_dealed` tinyint(1) NOT NULL DEFAULT '0' COMMENT ': 0,, 1,',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT ': 0,; 1,',
  `sn` varchar(32) NOT NULL DEFAULT '' COMMENT 'sn',
  PRIMARY KEY (`id`),
  KEY `idx_isdel` (`is_del`),
  KEY `collection_id_idx` (`collection_id`),
  KEY `is_dealed_idx` (`is_dealed`),
  KEY `sn_idx` (`sn`)
) ENGINE=InnoDB AUTO_INCREMENT=231 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cps_log_items`
--

LOCK TABLES `cps_log_items` WRITE;
/*!40000 ALTER TABLE `cps_log_items` DISABLE KEYS */;
/*!40000 ALTER TABLE `cps_log_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `device`
--

DROP TABLE IF EXISTS `device`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `device` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '设备ID编号',
  `sn` varchar(255) NOT NULL DEFAULT '' COMMENT '设备SN',
  `device_id` varchar(255) NOT NULL DEFAULT '' COMMENT '设备uuid',
  `instance_id` varchar(255) NOT NULL DEFAULT '' COMMENT '实例id',
  `user_id` varchar(255) NOT NULL DEFAULT '' COMMENT '实例id',
  `user_name` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
  `idc_id` varchar(255) NOT NULL DEFAULT '' COMMENT '机型uuid',
  `device_type_id` varchar(255) NOT NULL DEFAULT '' COMMENT '设备类型uuid',
  `manage_status` varchar(16) NOT NULL DEFAULT '' COMMENT '设备状态: in=已入库|putway=已上架|putawaying=上架中|putawayfail=上架失败|created=已创建|removed=已移除',
  `device_series` varchar(255) NOT NULL DEFAULT '' COMMENT '机型类型，如computer，storage，gpu，other',
  `reason` varchar(255) NOT NULL DEFAULT '' COMMENT '设备状态变更失败原因',
  `cabinet` varchar(64) NOT NULL DEFAULT '' COMMENT '机柜编码',
  `brand` varchar(255) NOT NULL DEFAULT '' COMMENT '品牌',
  `model` varchar(255) NOT NULL DEFAULT '' COMMENT '型号',
  `u_position` varchar(64) NOT NULL DEFAULT '' COMMENT 'U位',
  `ilo_ip` varchar(32) NOT NULL DEFAULT '' COMMENT '带外管理IP',
  `ilo_user` varchar(32) NOT NULL DEFAULT '' COMMENT '带外账号',
  `ilo_password` varchar(255) NOT NULL DEFAULT '' COMMENT '带外账号密码',
  `mac1` varchar(64) NOT NULL DEFAULT '' COMMENT 'MAC1（eth0）',
  `mac2` varchar(64) NOT NULL DEFAULT '' COMMENT 'MAC2（eth2）',
  `switch_ip1` varchar(255) NOT NULL DEFAULT '' COMMENT '交换机1ip',
  `switch_port1` varchar(255) NOT NULL DEFAULT '' COMMENT '交换机1port',
  `switch_ip2` varchar(255) NOT NULL DEFAULT '' COMMENT '交换机2ip',
  `switch_port2` varchar(255) NOT NULL DEFAULT '' COMMENT '交换机2port',
  `switch_user1` varchar(255) NOT NULL DEFAULT '' COMMENT '交换机1登录账号，如果为空，取所在机房的值',
  `switch_password1` varchar(255) NOT NULL DEFAULT '' COMMENT '交换机1登录密码',
  `switch_user2` varchar(255) NOT NULL DEFAULT '' COMMENT '交换机2登录账号，如果为空，取所在机房的值',
  `switch_password2` varchar(255) NOT NULL DEFAULT '' COMMENT '交换机2登录密码',
  `description` varchar(1024) NOT NULL DEFAULT '' COMMENT '描述',
  `switch_ip` varchar(255) NOT NULL DEFAULT '' COMMENT '网口交换机IP',
  `mask` varchar(255) NOT NULL DEFAULT '' COMMENT '子网掩码',
  `gateway` varchar(255) NOT NULL DEFAULT '' COMMENT '网关地址',
  `private_ipv4` varchar(255) NOT NULL DEFAULT '' COMMENT '内网IPV4',
  `private_ipv6` varchar(255) NOT NULL DEFAULT '' COMMENT '内网IPV6',
  `gateway6` varchar(255) NOT NULL DEFAULT '' COMMENT 'IPV6网关地址',
  `adapter_id` int(11) NOT NULL DEFAULT '0' COMMENT 'adapter_id',
  `raid_driver` varchar(255) NOT NULL DEFAULT '' COMMENT 'raid工具：（megacli64等）',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新者',
  `created_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间戳',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  `private_eth1_ipv4` varchar(255) NOT NULL DEFAULT '' COMMENT 'eth1 IPV4',
  `private_eth1_ipv6` varchar(255) NOT NULL DEFAULT '' COMMENT 'eth1 IPV6',
  `mask_eth1` varchar(255) NOT NULL DEFAULT '' COMMENT '子网掩码 for eth1',
  `current_boot_mode` varchar(255) NOT NULL DEFAULT '' COMMENT 'uefi/bios',
  `collect_fail_reason` varchar(255) NOT NULL DEFAULT '',
  `collect_status` varchar(255) NOT NULL DEFAULT '' COMMENT '1->collected success, 2->not collected, 3-> collecting, 4-> collected failed',
  `architecture` varchar(255) NOT NULL DEFAULT '' COMMENT 'from agent.response when collect,x86_64 or aarch64 or something else',
  `collect_method` tinyint(4) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `sn_idx` (`sn`),
  KEY `isdel_idx` (`is_del`)
) ENGINE=InnoDB AUTO_INCREMENT=118 DEFAULT CHARSET=utf8 COMMENT='设备信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `device`
--

LOCK TABLES `device` WRITE;
/*!40000 ALTER TABLE `device` DISABLE KEYS */;
/*!40000 ALTER TABLE `device` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `device_hints`
--

DROP TABLE IF EXISTS `device_hints`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `device_hints` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `volume_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'volume uuid',
  `sn` varchar(64) NOT NULL COMMENT '设备SN',
  `size` bigint(20) NOT NULL DEFAULT '0' COMMENT '以GiB为单位的设备尺寸',
  `rotational` tinyint(4) NOT NULL DEFAULT '0' COMMENT '转动，可以区分HDDs(旋转-磁盘)和SSDs(非旋转-固态硬盘)',
  `wwn` varchar(128) DEFAULT NULL COMMENT '唯一的存储标识符',
  `name` varchar(64) NOT NULL COMMENT '设备名称，例如/dev/md0',
  `wwn_vendor_extension` varchar(128) DEFAULT NULL COMMENT '唯一的供应商存储标识符',
  `wwn_with_extension` varchar(128) DEFAULT NULL COMMENT '带有供应商附加扩展的唯一的存储标识符',
  `model` varchar(64) DEFAULT NULL COMMENT '设备标识符',
  `serial` varchar(64) DEFAULT NULL COMMENT '磁盘序列号',
  `hctl` varchar(64) DEFAULT NULL COMMENT 'SCSI地址（主机名、channel通道、Target和Lun）',
  `by_path` varchar(128) DEFAULT NULL COMMENT '磁盘by_path路径',
  `vendor` varchar(64) DEFAULT NULL COMMENT '设备供应商',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除(0-未删, 1-已删)',
  `volume_type` varchar(255) NOT NULL DEFAULT '' COMMENT 'system|data',
  UNIQUE KEY `primary_id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1407 DEFAULT CHARSET=utf8 COMMENT='设备各盘符(卷)所在磁盘';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `device_hints`
--

LOCK TABLES `device_hints` WRITE;
/*!40000 ALTER TABLE `device_hints` DISABLE KEYS */;
/*!40000 ALTER TABLE `device_hints` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `device_type`
--

DROP TABLE IF EXISTS `device_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `device_type` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `idc_id` varchar(255) NOT NULL DEFAULT '' COMMENT '机房id',
  `device_type_id` varchar(255) NOT NULL DEFAULT '' COMMENT '设备类型uuid',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '机型名称，如计算效能型,标准计算型',
  `device_type` varchar(255) NOT NULL DEFAULT '' COMMENT '机型规格, cps.c.normal',
  `device_series` varchar(255) NOT NULL DEFAULT '' COMMENT '机型类型，如计算型，存储型',
  `architecture` varchar(255) NOT NULL DEFAULT '' COMMENT '体系架构，如i386/x86_64/ ARM64 (aarch64)，默认 x86_64',
  `height` int(11) NOT NULL DEFAULT '0' COMMENT '【高度（U）】：显示机型高度',
  `description` varchar(1024) NOT NULL DEFAULT '' COMMENT '描述',
  `cpu_spec` varchar(256) NOT NULL DEFAULT '' COMMENT 'cpu规格是预置还是其它，预置：common,其它：user_defined',
  `cpu_amount` int(11) NOT NULL DEFAULT '0' COMMENT 'cpu数量',
  `cpu_cores` int(11) NOT NULL DEFAULT '0' COMMENT '单个cpu内核数',
  `cpu_manufacturer` varchar(16) NOT NULL DEFAULT '' COMMENT 'cpu厂商',
  `cpu_model` varchar(16) NOT NULL DEFAULT '' COMMENT 'cpu处理器型号',
  `cpu_frequency` varchar(8) NOT NULL DEFAULT '' COMMENT 'cpu频率(G)',
  `mem_spec` varchar(256) NOT NULL DEFAULT '' COMMENT 'mem规格是预置还是其它，预置：common,其它：user_defined',
  `mem_type` varchar(4) NOT NULL DEFAULT '' COMMENT '内存接口（如DDR3，DDR4）',
  `mem_size` int(11) NOT NULL DEFAULT '0' COMMENT '单个内存大小(GB)',
  `mem_amount` int(11) NOT NULL DEFAULT '0' COMMENT '内存数量',
  `mem_frequency` int(11) NOT NULL DEFAULT '0' COMMENT '内存主频（MHz)',
  `nic_amount` int(11) NOT NULL DEFAULT '0' COMMENT '网卡数量',
  `nic_rate` int(11) NOT NULL DEFAULT '0' COMMENT '网卡传输速率(GE)',
  `interface_mode` varchar(255) NOT NULL DEFAULT '' COMMENT '【网口模式】【网络设置】: bond单网口,dual双网口',
  `system_volume_type` varchar(16) NOT NULL DEFAULT '' COMMENT '系统盘类型（SSD，HDD）',
  `system_volume_interface_type` varchar(16) NOT NULL DEFAULT '' COMMENT '系统盘接口类型（SATA,SAS,NVME）',
  `system_volume_size` int(11) NOT NULL DEFAULT '0' COMMENT '系统盘单盘大小',
  `system_volume_unit` varchar(16) NOT NULL DEFAULT '0' COMMENT '系统盘单盘单位',
  `system_volume_amount` int(11) NOT NULL DEFAULT '0' COMMENT '系统盘数量',
  `gpu_amount` int(11) NOT NULL DEFAULT '0' COMMENT 'gpu数量',
  `gpu_manufacturer` varchar(16) NOT NULL DEFAULT '' COMMENT 'gpu厂商',
  `gpu_model` varchar(16) NOT NULL DEFAULT '' COMMENT 'gpu处理器型号',
  `data_volume_type` varchar(16) NOT NULL DEFAULT '' COMMENT '数据盘类型（SSD，HDD）',
  `data_volume_interface_type` varchar(16) NOT NULL DEFAULT '' COMMENT '数据盘接口类型（SATA,SAS,NVME）',
  `data_volume_size` int(11) NOT NULL DEFAULT '0' COMMENT '数据盘单盘大小',
  `data_volume_unit` varchar(16) NOT NULL DEFAULT '0' COMMENT '数据盘单盘单位',
  `data_volume_amount` int(11) NOT NULL DEFAULT '0' COMMENT '数据盘数量',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新者',
  `created_time` int(255) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  `boot_mode` varchar(32) NOT NULL DEFAULT '' COMMENT 'boot_type:bios/uefi',
  `raid_can` varchar(16) NOT NULL DEFAULT '' COMMENT 'raid_can (RAID)/(NO RAID)',
  `is_need_raid` varchar(16) NOT NULL DEFAULT '' COMMENT 'need_raid|no_need_raid',
  `nics` text CHARACTER SET utf8mb4,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `i_device_type_type` (`device_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=256 DEFAULT CHARSET=utf8 COMMENT='设备型号详情，此表并未使用';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `device_type`
--

LOCK TABLES `device_type` WRITE;
/*!40000 ALTER TABLE `device_type` DISABLE KEYS */;
/*!40000 ALTER TABLE `device_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `disk`
--

DROP TABLE IF EXISTS `disk`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `disk` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `disk_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'disk uuid',
  `device_id` varchar(256) NOT NULL COMMENT '设备uuid',
  `enclosure` varchar(256) NOT NULL DEFAULT '' COMMENT 'enclosure',
  `slot` int(11) NOT NULL DEFAULT '0' COMMENT '卡槽槽位',
  `disk_type` varchar(256) NOT NULL COMMENT '磁盘类型：system,data',
  `types` varchar(256) NOT NULL DEFAULT '' COMMENT 'control/mvme/panfu',
  `size` varchar(256) NOT NULL DEFAULT '' COMMENT '硬盘大小，不确定精度(非nvme盘)',
  `size_unit` varchar(256) NOT NULL DEFAULT '' COMMENT '硬盘大小单位 MB GB TB ,1024进制',
  `pd_type` varchar(256) NOT NULL DEFAULT '' COMMENT '盘类型:接口、nvme、盘符',
  `adapter_id` int(11) NOT NULL DEFAULT '0' COMMENT '适配ID',
  `media_type` varchar(256) NOT NULL DEFAULT '' COMMENT 'SDD',
  `device_path` varchar(256) NOT NULL DEFAULT '' COMMENT 'effective when pd_type=NVME',
  `index` int(11) NOT NULL DEFAULT '0' COMMENT 'NVME索引',
  `serial_number` varchar(256) NOT NULL DEFAULT '' COMMENT 'effective when pd_type=NVME',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新者',
  `created_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间戳',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT 'disk name/ nvme name / panfu name',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `i_disk_id` (`disk_id`) USING BTREE,
  KEY `i_device_id` (`device_id`) USING BTREE,
  KEY `i_pd_type` (`pd_type`) USING BTREE,
  KEY `i_is_del` (`is_del`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2511 DEFAULT CHARSET=utf8 COMMENT='设备磁盘卡槽信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `disk`
--

LOCK TABLES `disk` WRITE;
/*!40000 ALTER TABLE `disk` DISABLE KEYS */;
/*!40000 ALTER TABLE `disk` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `idc`
--

DROP TABLE IF EXISTS `idc`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `idc` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `idc_id` varchar(50) NOT NULL DEFAULT '' COMMENT '机房uuid',
  `name` varchar(200) NOT NULL DEFAULT '' COMMENT '机房名称',
  `name_en` varchar(200) NOT NULL DEFAULT '' COMMENT '机房英文名称',
  `shortname` varchar(200) NOT NULL DEFAULT '' COMMENT '机房名称缩写',
  `level` varchar(200) NOT NULL DEFAULT '' COMMENT '机房等级',
  `address` varchar(512) NOT NULL COMMENT '机房地址',
  `ilo_user` varchar(255) NOT NULL DEFAULT '' COMMENT '带外登录账号',
  `ilo_password` varchar(255) NOT NULL DEFAULT '' COMMENT '带外登录密码',
  `switch_user1` varchar(255) NOT NULL DEFAULT '' COMMENT '交换机1登录账号',
  `switch_password1` varchar(255) NOT NULL DEFAULT '' COMMENT '交换机1登录密码',
  `switch_user2` varchar(255) NOT NULL DEFAULT '' COMMENT '交换机2登录账号',
  `switch_password2` varchar(255) NOT NULL DEFAULT '' COMMENT '交换机2登录密码',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新者',
  `created_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间戳',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `i_os_uuid` (`idc_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='os';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `idc`
--

LOCK TABLES `idc` WRITE;
/*!40000 ALTER TABLE `idc` DISABLE KEYS */;
INSERT INTO `idc` VALUES (1,'idc-vm4xsulx1k2d9z4xkctrttig02zl','默认机房','default_idc','default_idc','T4','北京市经济开发区京东总部2号楼','jdcps','cPs@Cp2019U','','','','','system','admin',1666099361,1726126235,0,0);
/*!40000 ALTER TABLE `idc` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `image`
--

DROP TABLE IF EXISTS `image`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `image` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `image_id` varchar(255) NOT NULL DEFAULT '' COMMENT '镜像uuid',
  `image_name` varchar(255) NOT NULL DEFAULT '' COMMENT '镜像名称',
  `os_id` varchar(36) NOT NULL DEFAULT '' COMMENT 'OSID',
  `format` varchar(16) NOT NULL DEFAULT '' COMMENT '镜像格式（qcow2、tar）',
  `filename` varchar(64) NOT NULL DEFAULT '' COMMENT '镜像文件名称',
  `url` varchar(256) NOT NULL DEFAULT '' COMMENT '镜像源路径',
  `hash` varchar(128) NOT NULL DEFAULT '' COMMENT '镜像校验码',
  `source` varchar(20) NOT NULL DEFAULT '' COMMENT '镜像来源(common通用、customize定制、user_defined自定义)',
  `description` varchar(512) NOT NULL DEFAULT '' COMMENT '描述',
  `system_partition` varchar(4096) NOT NULL DEFAULT '' COMMENT '系统分区信息（系统盘：“/ ” ：50GiB，xfs；swap：10GiB，swap）',
  `data_partition` varchar(4096) NOT NULL DEFAULT '' COMMENT '数据分区信息',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新者',
  `created_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间戳',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  `boot_mode` varchar(32) NOT NULL DEFAULT '' COMMENT 'boot_type:bios/uefi',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `i_image_uuid` (`image_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8 COMMENT='镜像';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `image`
--

LOCK TABLES `image` WRITE;
/*!40000 ALTER TABLE `image` DISABLE KEYS */;
INSERT INTO `image` VALUES (1,'i-haxcsxs570fhqew36zuuilks5w6v','CentOS 7.9 64-bit','o-3iqgx4yf6o5gfxerglpwkh3ej8h1','tar','v1.7.0-centos-7.9-2022070716.tar.xz','{host}/v1.7.0-centos-7.9-2022070716.tar.xz','760711ba6ae6e386d46314caf775ca14','common','centos7.9','[{\"format\":\"xfs\",\"point\":\"/\",\"size\":51200},{\"format\":\"swap\",\"point\":\"swap\",\"size\":10240}]','','admin','admin',1675325216,1675325216,0,0,'Legacy/BIOS,UEFI'),(2,'i-1bt811r0feavi4sfln5nyyz8qgvg','Ubuntu 18.04-64bit','o-04uay5mh6549w703schf1zsk3z29','tar','v1.7.0-ubuntu-18.04-2022062709.tar.xz','{host}/v1.7.0-ubuntu-18.04-2022062709.tar.xz','fcd77d7f8b22a3ff98671bf9890af46d','common','18.04','[{\"format\":\"xfs\",\"point\":\"/\",\"size\":51200},{\"format\":\"swap\",\"point\":\"swap\",\"size\":10240}]','','admin','admin',1675325530,1675325530,0,0,'Legacy/BIOS,UEFI');
/*!40000 ALTER TABLE `image` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `instance`
--

DROP TABLE IF EXISTS `instance`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `instance` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `instance_id` varchar(255) NOT NULL DEFAULT '' COMMENT '实例ID（uuid）',
  `idc_id` varchar(255) NOT NULL DEFAULT '' COMMENT '机房uuid',
  `project_id` varchar(255) NOT NULL DEFAULT '' COMMENT '项目id',
  `user_id` varchar(255) NOT NULL DEFAULT '' COMMENT '用户uuid',
  `instance_name` varchar(255) NOT NULL DEFAULT '' COMMENT '实例名称',
  `device_id` varchar(255) NOT NULL DEFAULT '' COMMENT '设备uuid',
  `device_type_id` varchar(255) NOT NULL DEFAULT '' COMMENT '机型uuid',
  `hostname` varchar(255) NOT NULL DEFAULT '' COMMENT '主机名',
  `status` varchar(255) NOT NULL DEFAULT '' COMMENT '运行状态',
  `reason` varchar(255) NOT NULL DEFAULT '' COMMENT '失败原因',
  `image_id` varchar(255) NOT NULL DEFAULT '' COMMENT '镜像uuid',
  `system_volume_raid_id` varchar(255) NOT NULL DEFAULT '' COMMENT '系统盘raid',
  `locked` char(255) NOT NULL DEFAULT '' COMMENT '是否锁定解锁锁定:locked,解锁unlocked',
  `data_volume_raid_id` varchar(255) NOT NULL DEFAULT '' COMMENT '数据盘raid',
  `description` varchar(2048) NOT NULL DEFAULT '' COMMENT '描述',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新者',
  `created_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间戳',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  `boot_mode` varchar(32) NOT NULL DEFAULT '' COMMENT 'boot_type:bios/uefi',
  `is_installed_agent` varchar(255) NOT NULL DEFAULT '' COMMENT '0-> uninstalled, 1-> installed',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=105 DEFAULT CHARSET=utf8 COMMENT='实例表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `instance`
--

LOCK TABLES `instance` WRITE;
/*!40000 ALTER TABLE `instance` DISABLE KEYS */;
/*!40000 ALTER TABLE `instance` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `instance_partition`
--

DROP TABLE IF EXISTS `instance_partition`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `instance_partition` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `instance_id` varchar(36) NOT NULL DEFAULT '' COMMENT '实例ID',
  `image_id` varchar(36) NOT NULL DEFAULT '' COMMENT '镜像ID',
  `device_type_id` varchar(100) NOT NULL DEFAULT '' COMMENT '设备类型',
  `boot_mode` varchar(32) NOT NULL DEFAULT '' COMMENT 'boot类型：bios、uefi',
  `partition_type` varchar(32) NOT NULL DEFAULT '' COMMENT '分区类型：root、boot、system、data',
  `partition_size` int(11) NOT NULL DEFAULT '0' COMMENT '分区大小，单位MB',
  `partition_fs_type` varchar(32) NOT NULL DEFAULT '' COMMENT '文件系统类型：xfs',
  `partition_mount_point` varchar(32) NOT NULL DEFAULT '' COMMENT '分区目录',
  `partition_label` varchar(32) NOT NULL DEFAULT '' COMMENT '分区标签：l_分区目录',
  `system_disk_label` varchar(30) NOT NULL DEFAULT '' COMMENT '系统盘分区格式：gpt、msdos（做完RAID系统盘大于4T必用gpt）',
  `data_disk_label` varchar(32) NOT NULL DEFAULT '' COMMENT '数据盘分区格式：gpt、msdos',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新者',
  `created_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间戳',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除(0-未删, 1-已删)',
  `instance_partition_id` varchar(36) NOT NULL DEFAULT '' COMMENT 'instance_partition_uuid',
  `number` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'instance_partition_number',
  `volume_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'volume uuid',
  `volume_type` varchar(255) NOT NULL DEFAULT '' COMMENT 'system|data',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=624 DEFAULT CHARSET=utf8 COMMENT='实例分区关系表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `instance_partition`
--

LOCK TABLES `instance_partition` WRITE;
/*!40000 ALTER TABLE `instance_partition` DISABLE KEYS */;
/*!40000 ALTER TABLE `instance_partition` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `interface`
--

DROP TABLE IF EXISTS `interface`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `interface` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `interface_name` varchar(16) DEFAULT NULL,
  `interface_type` varchar(16) DEFAULT NULL COMMENT 'lan/wan',
  `sn` varchar(64) NOT NULL COMMENT 'SN',
  `mac` varchar(32) NOT NULL COMMENT 'MAC',
  `switch_ip` varchar(32) NOT NULL COMMENT 'IP',
  `switch_port` varchar(32) DEFAULT NULL COMMENT 'Port',
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '(0-1-)',
  `ipv6_address` varchar(64) DEFAULT NULL,
  `nic_number` tinyint(3) unsigned DEFAULT NULL COMMENT '1234....',
  `ipv4_address` varchar(32) CHARACTER SET utf8mb4 DEFAULT NULL,
  `bus_info` varchar(64) CHARACTER SET utf8mb4 DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `i_interface_ssi` (`switch_ip`,`switch_port`,`is_del`),
  KEY `i_interface_sn` (`sn`)
) ENGINE=InnoDB AUTO_INCREMENT=3302 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `interface`
--

LOCK TABLES `interface` WRITE;
/*!40000 ALTER TABLE `interface` DISABLE KEYS */;
/*!40000 ALTER TABLE `interface` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mail_message`
--

DROP TABLE IF EXISTS `mail_message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mail_message` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `server_addr` varchar(255) NOT NULL DEFAULT '' COMMENT '服务器地址',
  `server_port` varchar(255) NOT NULL DEFAULT '' COMMENT '服务器端口',
  `admin_addr` varchar(255) NOT NULL DEFAULT '' COMMENT '管理员邮箱',
  `admin_pass` varchar(255) NOT NULL DEFAULT '' COMMENT '管理员邮箱密码',
  `is_push` varchar(16) NOT NULL DEFAULT '0' COMMENT 'is push mail',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  `to` varchar(255) NOT NULL DEFAULT '',
  `is_pass` varchar(16) NOT NULL DEFAULT '0' COMMENT 'is checked send ok',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=555 DEFAULT CHARSET=utf8 COMMENT='mail_message';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mail_message`
--

LOCK TABLES `mail_message` WRITE;
/*!40000 ALTER TABLE `mail_message` DISABLE KEYS */;
INSERT INTO `mail_message` VALUES (554,'','','','','1',0,'','');
/*!40000 ALTER TABLE `mail_message` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `monitor_alerts`
--

DROP TABLE IF EXISTS `monitor_alerts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `monitor_alerts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `alert_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'alert uuid',
  `alert_time` int(11) NOT NULL DEFAULT '0' COMMENT '告警时间戳',
  `resource` varchar(255) NOT NULL DEFAULT '' COMMENT '资源类型 [只支持instance一种]',
  `resource_id` varchar(255) NOT NULL DEFAULT '' COMMENT '资源id,目前就是实例id',
  `resource_name` varchar(255) NOT NULL DEFAULT '' COMMENT '资源名称,目前就是实例名称',
  `trigger` varchar(255) NOT NULL DEFAULT '' COMMENT '触发条件,接口需要翻译',
  `alert_value` varchar(255) NOT NULL DEFAULT '' COMMENT '报警值',
  `alert_level` tinyint(4) NOT NULL DEFAULT '0' COMMENT '1表示一般，2表示严重，3表示紧急',
  `alert_period` int(11) NOT NULL DEFAULT '0' COMMENT '告警持续时间(分钟为单位)',
  `created_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  `rule_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'rule uuid',
  `user_id` varchar(255) NOT NULL DEFAULT '' COMMENT '通知对象 userid',
  `user_name` varchar(255) NOT NULL DEFAULT '' COMMENT 'notice username',
  `rule_name` varchar(255) NOT NULL DEFAULT '' COMMENT 'rule name',
  `calculation_unit` varchar(255) NOT NULL DEFAULT '' COMMENT '%,GB,bps',
  `is_recover` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0->alert, 1->recover',
  `project_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'project uuid',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `i_rule_id` (`alert_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1773 DEFAULT CHARSET=utf8 COMMENT='告警历史表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `monitor_alerts`
--

LOCK TABLES `monitor_alerts` WRITE;
/*!40000 ALTER TABLE `monitor_alerts` DISABLE KEYS */;
/*!40000 ALTER TABLE `monitor_alerts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `monitor_rules`
--

DROP TABLE IF EXISTS `monitor_rules`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `monitor_rules` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `rule_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'rule uuid',
  `enabled` tinyint(4) NOT NULL DEFAULT '0' COMMENT '1表示启用 2表示不启用',
  `rule_name` varchar(255) NOT NULL DEFAULT '' COMMENT '规则名称',
  `dimension` varchar(255) NOT NULL DEFAULT '' COMMENT '维度 [instance disk mountpoint nic]',
  `resource` varchar(255) NOT NULL DEFAULT '' COMMENT '资源类型 [只支持instance一种]',
  `trigger_option` varchar(8192) NOT NULL DEFAULT '' COMMENT '触发条件, json',
  `notice_option` varchar(255) NOT NULL DEFAULT '' COMMENT '通知策略, json',
  `created_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  `user_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'user uuid',
  `tags` varchar(255) NOT NULL DEFAULT '' COMMENT 'disk/mountpoint/nic tag',
  `user_name` varchar(255) NOT NULL DEFAULT '' COMMENT 'user name',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '1->enable,2->disable,3->alerted',
  `project_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'project uuid',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `i_rule_id` (`rule_id`)
) ENGINE=InnoDB AUTO_INCREMENT=119 DEFAULT CHARSET=utf8 COMMENT='带内监控规则表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `monitor_rules`
--

LOCK TABLES `monitor_rules` WRITE;
/*!40000 ALTER TABLE `monitor_rules` DISABLE KEYS */;
/*!40000 ALTER TABLE `monitor_rules` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `os`
--

DROP TABLE IF EXISTS `os`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `os` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `os_id` varchar(255) NOT NULL DEFAULT '' COMMENT '操作系统uuid',
  `os_name` varchar(255) NOT NULL DEFAULT '' COMMENT '操作系统名称',
  `os_type` varchar(16) NOT NULL DEFAULT '' COMMENT '操作系统分类:linux/windows',
  `architecture` varchar(16) NOT NULL DEFAULT '' COMMENT '架构:x86/x64/i386/',
  `bits` int(11) NOT NULL DEFAULT '0' COMMENT '指令宽度:64/32位',
  `os_version` varchar(256) NOT NULL DEFAULT '' COMMENT '操作系统版本',
  `sys_user` varchar(16) NOT NULL DEFAULT '' COMMENT '管理员账户',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新者',
  `created_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间戳',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `i_os_uuid` (`os_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8 COMMENT='os';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `os`
--

LOCK TABLES `os` WRITE;
/*!40000 ALTER TABLE `os` DISABLE KEYS */;
INSERT INTO `os` VALUES (1,'o-3iqgx4yf6o5gfxerglpwkh3ej8h1','CentOS 7.9 64-bit','CentOS','x86_64',64,'7.9','root','admin','admin',1675325216,1675325216,0,0),(2,'o-04uay5mh6549w703schf1zsk3z29','Ubuntu 18.04-64bit','Ubuntu','x86_64',64,'18.04','root','admin','admin',1675325530,1675325530,0,0);
/*!40000 ALTER TABLE `os` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `project`
--

DROP TABLE IF EXISTS `project`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `project` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `project_id` varchar(255) NOT NULL DEFAULT '' COMMENT '项目uuid',
  `project_name` varchar(255) NOT NULL DEFAULT '' COMMENT '项目名称',
  `is_default` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否默认项目0否 1是',
  `is_system` tinyint(4) NOT NULL COMMENT '是否系统项目 0否 1是',
  `user_id` varchar(255) NOT NULL DEFAULT '' COMMENT '项目拥有者用户id',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新者',
  `created_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间戳',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  `description` varchar(255) NOT NULL DEFAULT '' COMMENT 'project description',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8 COMMENT='项目';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `project`
--

LOCK TABLES `project` WRITE;
/*!40000 ALTER TABLE `project` DISABLE KEYS */;
INSERT INTO `project` VALUES (1,'project-vtdz6jyqux7gxgglzk5fsdiyqy8k','default project',1,1,'user-ta5c5tsos2wkm8d2qtmvx3vufr2h','admin','admin',1678342000,1722335468,1722335468,0,'');
/*!40000 ALTER TABLE `project` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `r_device_type_image`
--

DROP TABLE IF EXISTS `r_device_type_image`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `r_device_type_image` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `image_id` varchar(255) NOT NULL DEFAULT '' COMMENT '镜像ID',
  `device_type_id` varchar(255) NOT NULL DEFAULT '' COMMENT '设备类型id',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新者',
  `created_time` int(255) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `i_r_device_type` (`device_type_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=299 DEFAULT CHARSET=utf8 COMMENT='device_type/image关系';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `r_device_type_image`
--

LOCK TABLES `r_device_type_image` WRITE;
/*!40000 ALTER TABLE `r_device_type_image` DISABLE KEYS */;
/*!40000 ALTER TABLE `r_device_type_image` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `r_device_type_image_partition`
--

DROP TABLE IF EXISTS `r_device_type_image_partition`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `r_device_type_image_partition` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `image_id` varchar(36) NOT NULL DEFAULT '' COMMENT '镜像ID',
  `device_type_id` varchar(100) NOT NULL DEFAULT '' COMMENT '设备类型',
  `boot_mode` varchar(32) NOT NULL DEFAULT '' COMMENT 'boot类型：bios、uefi',
  `partition_type` varchar(32) NOT NULL DEFAULT '' COMMENT '分区类型：root、boot、system、data',
  `partition_size` int(11) NOT NULL DEFAULT '0' COMMENT '分区大小，单位MB',
  `partition_fs_type` varchar(32) NOT NULL DEFAULT '' COMMENT '文件系统类型：xfs',
  `partition_mount_point` varchar(32) NOT NULL DEFAULT '' COMMENT '分区目录',
  `partition_label` varchar(32) NOT NULL DEFAULT '' COMMENT '分区标签：l_分区目录',
  `system_disk_label` varchar(30) NOT NULL DEFAULT '' COMMENT '系统盘分区格式：gpt、msdos（做完RAID系统盘大于4T必用gpt）',
  `data_disk_label` varchar(32) NOT NULL DEFAULT '' COMMENT '数据盘分区格式：gpt、msdos',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新者',
  `created_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间戳',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除(0-未删, 1-已删)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=506 DEFAULT CHARSET=utf8 COMMENT='device_type/image分区关系表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `r_device_type_image_partition`
--

LOCK TABLES `r_device_type_image_partition` WRITE;
/*!40000 ALTER TABLE `r_device_type_image_partition` DISABLE KEYS */;
/*!40000 ALTER TABLE `r_device_type_image_partition` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `r_device_volume_disks`
--

DROP TABLE IF EXISTS `r_device_volume_disks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `r_device_volume_disks` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `device_id` varchar(255) NOT NULL DEFAULT '' COMMENT '设备uuid',
  `volume_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'volume uuid',
  `disk_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'disk uuid',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新者',
  `created_time` int(255) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  `volume_type` varchar(255) NOT NULL DEFAULT '' COMMENT 'data|system',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `i_device_id` (`device_id`) USING BTREE,
  KEY `i_volume_id` (`volume_id`) USING BTREE,
  KEY `i_disk_id` (`disk_id`) USING BTREE,
  KEY `i_is_del` (`is_del`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1083 DEFAULT CHARSET=utf8 COMMENT='设备-卷-磁盘 关系表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `r_device_volume_disks`
--

LOCK TABLES `r_device_volume_disks` WRITE;
/*!40000 ALTER TABLE `r_device_volume_disks` DISABLE KEYS */;
INSERT INTO `r_device_volume_disks` VALUES (1018,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-f2lfzmkllmbr5761067uico754xy','disk-y14tuk1sp29e27mq3c7hbpkrh66e','admin','admin',1726034901,1726034901,0,1,'system'),(1019,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-f2lfzmkllmbr5761067uico754xy','disk-38wl5et15pbaav3vy3tfy8iokk45','admin','admin',1726034901,1726034901,0,1,'system'),(1020,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-f2lfzmkllmbr5761067uico754xy','disk-ajauiohh19msr3k7hb3avdnuvggu','admin','admin',1726034901,1726034901,0,1,'system'),(1021,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-tril2lumeqijcthvrdulmuiw3vps','disk-zlap1sbe3j6wphx35rubc1yfq0wg','admin','admin',1726034901,1726034901,0,1,'data'),(1022,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-tril2lumeqijcthvrdulmuiw3vps','disk-u7sbhfdopeb7v3rnizijm1k7hgng','admin','admin',1726034901,1726034901,0,1,'data'),(1023,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-tril2lumeqijcthvrdulmuiw3vps','disk-5ycq3r8zy2ib5f3ej1hnlgbtoeoz','admin','admin',1726034901,1726034901,0,1,'data'),(1024,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-10s5phbaizv3e68ic94r3vy0yrty','disk-y14tuk1sp29e27mq3c7hbpkrh66e','admin','admin',1726040988,1726040988,0,1,'system'),(1025,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-10s5phbaizv3e68ic94r3vy0yrty','disk-38wl5et15pbaav3vy3tfy8iokk45','admin','admin',1726040988,1726040988,0,1,'system'),(1026,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-10s5phbaizv3e68ic94r3vy0yrty','disk-ajauiohh19msr3k7hb3avdnuvggu','admin','admin',1726040988,1726040988,0,1,'system'),(1027,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-md8cup28q0uxoj6h2798xu899kp5','disk-zlap1sbe3j6wphx35rubc1yfq0wg','admin','admin',1726040988,1726040988,0,1,'data'),(1028,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-md8cup28q0uxoj6h2798xu899kp5','disk-u7sbhfdopeb7v3rnizijm1k7hgng','admin','admin',1726040988,1726040988,0,1,'data'),(1029,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-md8cup28q0uxoj6h2798xu899kp5','disk-5ycq3r8zy2ib5f3ej1hnlgbtoeoz','admin','admin',1726040988,1726040988,0,1,'data'),(1030,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-9hxfqwx7guzyxmhuvtosuz0iu9iw','disk-y14tuk1sp29e27mq3c7hbpkrh66e','admin','admin',1726042102,1726042102,0,1,'system'),(1031,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-9hxfqwx7guzyxmhuvtosuz0iu9iw','disk-38wl5et15pbaav3vy3tfy8iokk45','admin','admin',1726042102,1726042102,0,1,'system'),(1032,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-f4gkqpm76zj1bq8pibg1rzxmusci','disk-ajauiohh19msr3k7hb3avdnuvggu','admin','admin',1726042102,1726042102,0,1,'data'),(1033,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-f4gkqpm76zj1bq8pibg1rzxmusci','disk-zlap1sbe3j6wphx35rubc1yfq0wg','admin','admin',1726042102,1726042102,0,1,'data'),(1034,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-9hxfqwx7guzyxmhuvtosuz0iu9iw','disk-y14tuk1sp29e27mq3c7hbpkrh66e','admin','admin',1726050714,1726050714,0,1,'system'),(1035,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-9hxfqwx7guzyxmhuvtosuz0iu9iw','disk-38wl5et15pbaav3vy3tfy8iokk45','admin','admin',1726050714,1726050714,0,1,'system'),(1036,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-f4gkqpm76zj1bq8pibg1rzxmusci','disk-ajauiohh19msr3k7hb3avdnuvggu','admin','admin',1726050714,1726050714,0,1,'data'),(1037,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-f4gkqpm76zj1bq8pibg1rzxmusci','disk-zlap1sbe3j6wphx35rubc1yfq0wg','admin','admin',1726050714,1726050714,0,1,'data'),(1038,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-9hxfqwx7guzyxmhuvtosuz0iu9iw','disk-y14tuk1sp29e27mq3c7hbpkrh66e','admin','admin',1726050965,1726050965,0,1,'system'),(1039,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-9hxfqwx7guzyxmhuvtosuz0iu9iw','disk-38wl5et15pbaav3vy3tfy8iokk45','admin','admin',1726050965,1726050965,0,1,'system'),(1040,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-f4gkqpm76zj1bq8pibg1rzxmusci','disk-zlap1sbe3j6wphx35rubc1yfq0wg','admin','admin',1726050965,1726050965,0,1,'data'),(1041,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-f4gkqpm76zj1bq8pibg1rzxmusci','disk-ajauiohh19msr3k7hb3avdnuvggu','admin','admin',1726050965,1726050965,0,1,'data'),(1042,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-9hxfqwx7guzyxmhuvtosuz0iu9iw','disk-y14tuk1sp29e27mq3c7hbpkrh66e','admin','admin',1726051070,1726051070,0,1,'system'),(1043,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-9hxfqwx7guzyxmhuvtosuz0iu9iw','disk-ajauiohh19msr3k7hb3avdnuvggu','admin','admin',1726051070,1726051070,0,1,'system'),(1044,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-f4gkqpm76zj1bq8pibg1rzxmusci','disk-38wl5et15pbaav3vy3tfy8iokk45','admin','admin',1726051070,1726051070,0,1,'data'),(1045,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-f4gkqpm76zj1bq8pibg1rzxmusci','disk-zlap1sbe3j6wphx35rubc1yfq0wg','admin','admin',1726051070,1726051070,0,1,'data'),(1046,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-9hxfqwx7guzyxmhuvtosuz0iu9iw','disk-y14tuk1sp29e27mq3c7hbpkrh66e','admin','admin',1726051124,1726051124,0,1,'system'),(1047,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-9hxfqwx7guzyxmhuvtosuz0iu9iw','disk-ajauiohh19msr3k7hb3avdnuvggu','admin','admin',1726051124,1726051124,0,1,'system'),(1048,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-f4gkqpm76zj1bq8pibg1rzxmusci','disk-38wl5et15pbaav3vy3tfy8iokk45','admin','admin',1726051124,1726051124,0,1,'data'),(1049,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-f4gkqpm76zj1bq8pibg1rzxmusci','disk-u7sbhfdopeb7v3rnizijm1k7hgng','admin','admin',1726051124,1726051124,0,1,'data'),(1050,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-9hxfqwx7guzyxmhuvtosuz0iu9iw','disk-y14tuk1sp29e27mq3c7hbpkrh66e','admin','admin',1726107009,1726107009,0,1,'system'),(1051,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-9hxfqwx7guzyxmhuvtosuz0iu9iw','disk-38wl5et15pbaav3vy3tfy8iokk45','admin','admin',1726107009,1726107009,0,1,'system'),(1052,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-f4gkqpm76zj1bq8pibg1rzxmusci','disk-ajauiohh19msr3k7hb3avdnuvggu','admin','admin',1726107009,1726107009,0,1,'data'),(1053,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-f4gkqpm76zj1bq8pibg1rzxmusci','disk-zlap1sbe3j6wphx35rubc1yfq0wg','admin','admin',1726107009,1726107009,0,1,'data'),(1054,'d-4omuvfyprcalrbh7euvotavi0bqq','vl-e14318rf9peds0lqqzpsg18gz3ob','disk-oldkxkr9d620oui7jruvmduo13vx','admin','admin',1726126872,1726126872,0,1,'system'),(1055,'d-4omuvfyprcalrbh7euvotavi0bqq','vl-e14318rf9peds0lqqzpsg18gz3ob','disk-dirka822l9fxziwqc12toenfzhy4','admin','admin',1726126872,1726126872,0,1,'system'),(1056,'d-4omuvfyprcalrbh7euvotavi0bqq','vl-e14318rf9peds0lqqzpsg18gz3ob','disk-eyof6e2fgalbpu2e42i950kqdfi6','admin','admin',1726126872,1726126872,0,1,'system'),(1057,'d-4omuvfyprcalrbh7euvotavi0bqq','vl-e14318rf9peds0lqqzpsg18gz3ob','disk-6ha0dr3xbigczzi8el7v5rg199r4','admin','admin',1726126872,1726126872,0,1,'system'),(1058,'d-4omuvfyprcalrbh7euvotavi0bqq','vl-xwwibm0c9nsuac3hizh3t9t181ca','disk-hcwk1qjt5yqm8slnh4kxl6p2dz6b','admin','admin',1726126872,1726126872,0,1,'data'),(1059,'d-4omuvfyprcalrbh7euvotavi0bqq','vl-xwwibm0c9nsuac3hizh3t9t181ca','disk-fwkx8ibkrtmkfel0z03njm13sh17','admin','admin',1726126872,1726126872,0,1,'data'),(1060,'d-4omuvfyprcalrbh7euvotavi0bqq','vl-xwwibm0c9nsuac3hizh3t9t181ca','disk-yqc6121eja29u7ckgci2zwzood63','admin','admin',1726126872,1726126872,0,1,'data'),(1061,'d-4omuvfyprcalrbh7euvotavi0bqq','vl-xwwibm0c9nsuac3hizh3t9t181ca','disk-f503j9hn8l2d6ihumceryp8tgkmn','admin','admin',1726126872,1726126872,0,1,'data'),(1062,'d-brj02kttnzcatjs6iynhw2icxzbg','vl-e14318rf9peds0lqqzpsg18gz3ob','disk-72dkds4vvqrlh847y0mmrencez11','admin','admin',1726127260,1726127260,0,0,'system'),(1063,'d-brj02kttnzcatjs6iynhw2icxzbg','vl-e14318rf9peds0lqqzpsg18gz3ob','disk-6d0h79kogkeyxjabezca42ij4cc9','admin','admin',1726127260,1726127260,0,0,'system'),(1064,'d-brj02kttnzcatjs6iynhw2icxzbg','vl-xwwibm0c9nsuac3hizh3t9t181ca','disk-ssa0m02jb2beujguea7kjwhnephi','admin','admin',1726127260,1726127260,0,0,'data'),(1065,'d-brj02kttnzcatjs6iynhw2icxzbg','vl-xwwibm0c9nsuac3hizh3t9t181ca','disk-e44c992gpjmcd29m3wdw102o04tt','admin','admin',1726127260,1726127260,0,0,'data'),(1066,'d-kxqwytjua091giupgjk8hqiky7ml','vl-e5cdm3q8h8kovbqstp9rh9vfh5u6','disk-5nfaqrm9kl6mg7u8lc5w84avu2vl','admin','admin',1726129391,1726129391,0,1,'system'),(1067,'d-kxqwytjua091giupgjk8hqiky7ml','vl-e5cdm3q8h8kovbqstp9rh9vfh5u6','disk-5nfaqrm9kl6mg7u8lc5w84avu2vl','admin','admin',1726130391,1726130391,0,1,'system'),(1068,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-7a6m0y03b5144nfcwxgukytwgaam','disk-kp1soakbm4u58oi2rkh6jj9ksny2','admin','admin',1726142501,1726142501,0,1,'system'),(1069,'d-4omuvfyprcalrbh7euvotavi0bqq','vl-wlrgrcmpk5rkfg0gs7lbkz9cqopa','disk-q349uunnl4dz4z0lk851yjbm46pc','admin','admin',1726142700,1726142700,0,0,'system'),(1070,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-191o946aefsmgi7lmeo0e62xrct2','disk-tvm6jqyly0i7khpjswzpzw9spip8','admin','admin',1726143815,1726143815,0,0,'system'),(1071,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-191o946aefsmgi7lmeo0e62xrct2','disk-g4s92psenzdrsysf3pxejuze7i34','admin','admin',1726143815,1726143815,0,0,'system'),(1072,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-191o946aefsmgi7lmeo0e62xrct2','disk-c3sgf073fbg9xk0umyjh7v6i6b0c','admin','admin',1726143815,1726143815,0,0,'system'),(1073,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-191o946aefsmgi7lmeo0e62xrct2','disk-ebnvdjqhatm58h2e2as5o6csdane','admin','admin',1726143815,1726143815,0,0,'system'),(1074,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-191o946aefsmgi7lmeo0e62xrct2','disk-bx10c0oxobyms6l5xafcfix90zb6','admin','admin',1726143815,1726143815,0,0,'system'),(1075,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-191o946aefsmgi7lmeo0e62xrct2','disk-pnem7w4crdyv9l253it877nbiaye','admin','admin',1726143815,1726143815,0,0,'system'),(1076,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-191o946aefsmgi7lmeo0e62xrct2','disk-l8gbjleigazlr1m8obop82z96opy','admin','admin',1726143815,1726143815,0,0,'system'),(1077,'d-nca3rvyoc2s67hoqcu597nl2pd6k','vl-191o946aefsmgi7lmeo0e62xrct2','disk-ux763wfd8697hoopj0ibqserp3xm','admin','admin',1726143815,1726143815,0,0,'system'),(1078,'d-kxqwytjua091giupgjk8hqiky7ml','vl-wlrgrcmpk5rkfg0gs7lbkz9cqopa','disk-kydj8bf60rfrduo94kmx4ntjb2f5','admin','admin',1726143831,1726143831,0,1,'system'),(1079,'d-kxqwytjua091giupgjk8hqiky7ml','vl-wlrgrcmpk5rkfg0gs7lbkz9cqopa','disk-ymkygj5n2gpkfvd0i8tuzl1txfaa','admin','admin',1726145157,1726145157,0,1,'system'),(1080,'d-kxqwytjua091giupgjk8hqiky7ml','vl-wlrgrcmpk5rkfg0gs7lbkz9cqopa','disk-kydj8bf60rfrduo94kmx4ntjb2f5','admin','admin',1726211007,1726211007,0,1,'system'),(1081,'d-kxqwytjua091giupgjk8hqiky7ml','vl-wlrgrcmpk5rkfg0gs7lbkz9cqopa','disk-r5vepgotjkru7y34y1xjrl2sutu5','admin','admin',1726221183,1726221183,0,1,'system'),(1082,'d-kxqwytjua091giupgjk8hqiky7ml','vl-wlrgrcmpk5rkfg0gs7lbkz9cqopa','disk-3y24cy5fr6f5zijucyc4io996qy6','admin','admin',1726221196,1726221196,0,0,'system');
/*!40000 ALTER TABLE `r_device_volume_disks` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `r_instance_sshkey`
--

DROP TABLE IF EXISTS `r_instance_sshkey`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `r_instance_sshkey` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `instance_id` varchar(255) NOT NULL DEFAULT '' COMMENT '实例ID',
  `sshkey_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'sshkeyid',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新者',
  `created_time` int(255) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='实例ssh秘钥关系表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `r_instance_sshkey`
--

LOCK TABLES `r_instance_sshkey` WRITE;
/*!40000 ALTER TABLE `r_instance_sshkey` DISABLE KEYS */;
/*!40000 ALTER TABLE `r_instance_sshkey` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `r_instance_volume_raid`
--

DROP TABLE IF EXISTS `r_instance_volume_raid`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `r_instance_volume_raid` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `instance_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'instance uuid',
  `volume_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'volume uuid',
  `volume_type` varchar(255) NOT NULL DEFAULT '' COMMENT 'data|system',
  `raid_can` varchar(16) NOT NULL DEFAULT '' COMMENT 'RAID配置： (RAID,NO RAID)',
  `raid_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'RAID模式：raidid,一个一条',
  `raid_name` varchar(255) NOT NULL DEFAULT '' COMMENT 'RAID名称',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新者',
  `created_time` int(255) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `i_instance_id` (`instance_id`) USING BTREE,
  KEY `i_volume_type` (`volume_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET=utf8 COMMENT='实例每个数据卷选择的raid';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `r_instance_volume_raid`
--

LOCK TABLES `r_instance_volume_raid` WRITE;
/*!40000 ALTER TABLE `r_instance_volume_raid` DISABLE KEYS */;
/*!40000 ALTER TABLE `r_instance_volume_raid` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `r_monitor_rules_instance`
--

DROP TABLE IF EXISTS `r_monitor_rules_instance`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `r_monitor_rules_instance` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `rule_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'rule uuid',
  `rule_name` varchar(255) NOT NULL DEFAULT '' COMMENT '规则名称',
  `instance_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'instance uuid',
  `instance_name` varchar(255) NOT NULL DEFAULT '' COMMENT 'instance name',
  `created_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  `tags` varchar(255) NOT NULL DEFAULT '' COMMENT 'disk/mountpoint/nic tag',
  `project_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'project uuid',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `i_rule_id` (`rule_id`),
  KEY `i_instance_id` (`instance_id`)
) ENGINE=InnoDB AUTO_INCREMENT=215 DEFAULT CHARSET=utf8 COMMENT='带内监控规则-实例关联表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `r_monitor_rules_instance`
--

LOCK TABLES `r_monitor_rules_instance` WRITE;
/*!40000 ALTER TABLE `r_monitor_rules_instance` DISABLE KEYS */;
/*!40000 ALTER TABLE `r_monitor_rules_instance` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `r_volume_raid`
--

DROP TABLE IF EXISTS `r_volume_raid`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `r_volume_raid` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `volume_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'volume uuid',
  `device_type_id` varchar(255) NOT NULL DEFAULT '' COMMENT '设备类型uuid',
  `raid_can` varchar(16) NOT NULL DEFAULT '' COMMENT 'RAID配置： (RAID,NO RAID)',
  `raid_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'RAID模式：raidid,一个一条',
  `raid_name` varchar(255) NOT NULL DEFAULT '' COMMENT 'RAID名称',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新者',
  `created_time` int(255) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  `volume_type` varchar(255) NOT NULL DEFAULT '' COMMENT 'data|system',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=822 DEFAULT CHARSET=utf8 COMMENT='卷管理表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `r_volume_raid`
--

LOCK TABLES `r_volume_raid` WRITE;
/*!40000 ALTER TABLE `r_volume_raid` DISABLE KEYS */;
/*!40000 ALTER TABLE `r_volume_raid` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `raid`
--

DROP TABLE IF EXISTS `raid`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `raid` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `raid_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'raid uuid',
  `name` varchar(16) NOT NULL DEFAULT '' COMMENT 'raid类型 noraid，raid0，raid1，raid10',
  `description_en` varchar(512) NOT NULL DEFAULT '' COMMENT 'RAID英文描述',
  `description_zh` varchar(256) NOT NULL DEFAULT '' COMMENT 'RAID中文描述',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新者',
  `created_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间戳',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `i_raid_uuid` (`raid_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='raid';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `raid`
--

LOCK TABLES `raid` WRITE;
/*!40000 ALTER TABLE `raid` DISABLE KEYS */;
INSERT INTO `raid` VALUES (1,'r-qytwf9r5h0yn9x4evjkyr0n1cwyb','NO RAID','NO RAID means that each physical disk is displayed independently as a logical disk. There is no block stripping and no data redundancy.','NO RAID是每个物理磁盘独立显示为一个逻辑盘，没有数据块分条（no block stripping），不提供数据冗余。','','',0,0,0,0),(2,'r-sfykwibcvcyztuez3mk1klyha9nn','RAID0','RAID0 means that the entire logical disk data is striped among multiple physical disks, it allows read/write in parallel and provides the fastest speed with no redundancy.','RAID0是整个逻辑盘的数据是被分条（stripped）分布在多个物理磁盘上，可以并行读/写，提供最快的速度，但没有冗余能力。','','',0,0,0,0),(3,'r-wtzluqacgzzxgunnabdkpnpjew3d','RAID1','RAID1 also known as mirroring method, provides data redundancy. In the entire mirroring process, only half of the disk capacity is valid, the other half is used to store the same data. RAID1 has taken security into consideration with half capacity and full speed.','RAID1又称镜像方式，提供数据的冗余。在整个镜像过程中，只有一半的磁盘容量是有效的（另一半磁盘容量用来存放同样的数据），RAID1考虑了安全性，容量减半、速度不变。','','',0,0,0,0),(4,'r-l6pounvfife0njlinhk6ztf2xyio','RAID10','RAID10 is used for achieving the high speed and security purposes. RAID 10 can be simply interpreted as the RAID 0 array composed of several disks, which then provides the image.','RAID10是为了达到既高速又安全，可以把RAID 10简单地理解成由多个磁盘组成的RAID 0阵列再进行镜像。','','',0,0,0,0),(5,'r-lowdd8nf8y7v6jk5ocr5paqhvhre','RAID51H','RAID 5+1 hotspare offers you a resolution featuring high performance, data security and reasonable cost in storage, with added hot backup disks for better data security.','RAID 5+1 hotspare是一种存储性能、数据安全和存储成本兼顾的存储解决方案，同时增加热备盘来提升数据安全。','','',0,0,0,1),(6,'r-493obv9i0fyubxfgnuatrffj5qy0','RAID5','raid5 is...','raid5 是一种。。。','','',0,0,0,0);
/*!40000 ALTER TABLE `raid` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role`
--

DROP TABLE IF EXISTS `role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` varchar(255) NOT NULL DEFAULT '' COMMENT '角色uuid',
  `role_name_en` varchar(255) NOT NULL DEFAULT '' COMMENT '角色名称，唯一',
  `role_name_cn` varchar(255) NOT NULL DEFAULT '' COMMENT '角色名称，唯一',
  `description_en` varchar(255) NOT NULL DEFAULT '' COMMENT '权限描述',
  `description_cn` varchar(255) NOT NULL DEFAULT '' COMMENT '权限描述',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新者',
  `created_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间戳',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='角色';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role`
--

LOCK TABLES `role` WRITE;
/*!40000 ALTER TABLE `role` DISABLE KEYS */;
INSERT INTO `role` VALUES (1,'role-admin-uuid-01','Owner','平台拥有者','Possess access to the console & operation platform, and all operation rights.','拥有访问控制台&运营平台访问权限，全部操作权限。','admin','admin',1666275445,1666275445,0,0),(2,'role-operator-uuid-01','Operator','运营人员','Possess access rights to the operation platform, all operation rights (except user management and role management).','拥有运营平台访问权限，所有操作权限（除用户管理和角色管理）。','admin','admin',1666278266,1666278266,0,0),(3,'role-user-uuid-01','User','普通用户','Have access to the console, all operations, no access to the operating platform.','拥有控制台访问权限，所有操作权限，无运营平台访问权限。','admin','admin',1666275474,1666275474,0,0);
/*!40000 ALTER TABLE `role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sharing_project`
--

DROP TABLE IF EXISTS `sharing_project`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sharing_project` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `project_id` varchar(255) NOT NULL DEFAULT '' COMMENT '项目uuid',
  `owner_user_id` varchar(255) NOT NULL DEFAULT '' COMMENT '项目拥有者用户id',
  `shared_user_id` varchar(255) NOT NULL DEFAULT '' COMMENT '项目共享者用户id',
  `owner_user_name` varchar(255) NOT NULL DEFAULT '' COMMENT '项目拥有者用户名',
  `shared_user_name` varchar(255) NOT NULL DEFAULT '' COMMENT '项目拥有者用户名',
  `is_default` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否默认项目0否 1是',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新者',
  `created_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间戳',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  `project_name` varchar(255) NOT NULL DEFAULT '',
  `premission` varchar(64) NOT NULL DEFAULT '' COMMENT 'read or write premission',
  `shared_instance_ids` text NOT NULL COMMENT 'sharding instances',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COMMENT='共享项目';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sharing_project`
--

LOCK TABLES `sharing_project` WRITE;
/*!40000 ALTER TABLE `sharing_project` DISABLE KEYS */;
/*!40000 ALTER TABLE `sharing_project` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sms_Logs`
--

DROP TABLE IF EXISTS `sms_Logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sms_Logs` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `service_provider` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `phone` varchar(16) DEFAULT NULL,
  `send_time` int(11) NOT NULL,
  `is_success` varchar(255) NOT NULL,
  `err_msg` text NOT NULL,
  `is_del` varchar(11) NOT NULL,
  `phone_prefix` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=116 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sms_Logs`
--

LOCK TABLES `sms_Logs` WRITE;
/*!40000 ALTER TABLE `sms_Logs` DISABLE KEYS */;
/*!40000 ALTER TABLE `sms_Logs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sms_template`
--

DROP TABLE IF EXISTS `sms_template`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sms_template` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `service_provider` varchar(255) NOT NULL,
  `secret_id` varchar(255) NOT NULL,
  `secret_key` varchar(255) NOT NULL,
  `sign_code` varchar(255) NOT NULL,
  `template_id` varchar(255) NOT NULL,
  `tencent_app_id` varchar(255) NOT NULL,
  `hw_from` varchar(255) NOT NULL,
  `hw_callback_url` varchar(255) NOT NULL,
  `is_push` varchar(255) NOT NULL,
  `is_del` varchar(255) NOT NULL,
  `is_pass` varchar(255) NOT NULL,
  `created_time` int(11) NOT NULL,
  `updated_time` int(11) NOT NULL,
  `deleted_time` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sms_template`
--

LOCK TABLES `sms_template` WRITE;
/*!40000 ALTER TABLE `sms_template` DISABLE KEYS */;
/*!40000 ALTER TABLE `sms_template` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ssh_key`
--

DROP TABLE IF EXISTS `ssh_key`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ssh_key` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sshkey_id` varchar(255) NOT NULL DEFAULT '' COMMENT '秘钥uuid',
  `user_id` varchar(255) NOT NULL DEFAULT '' COMMENT '用户uuid',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '秘钥名称',
  `key` varchar(1024) NOT NULL DEFAULT '0' COMMENT '公钥，格式：ssh-rsa AAA',
  `finger_print` varchar(255) NOT NULL DEFAULT '' COMMENT '公钥指纹',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新者',
  `created_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间戳',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='秘钥';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ssh_key`
--

LOCK TABLES `ssh_key` WRITE;
/*!40000 ALTER TABLE `ssh_key` DISABLE KEYS */;
/*!40000 ALTER TABLE `ssh_key` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) NOT NULL DEFAULT '' COMMENT '用户uuid',
  `role_id` varchar(255) NOT NULL DEFAULT '' COMMENT '角色uuid',
  `user_name` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名，唯一',
  `email` varchar(255) NOT NULL DEFAULT '' COMMENT '邮箱',
  `phone_prefix` varchar(255) NOT NULL DEFAULT '' COMMENT '国家地区码，如86',
  `phone_number` varchar(255) NOT NULL DEFAULT '' COMMENT '手机号',
  `default_project_id` varchar(255) NOT NULL DEFAULT '' COMMENT '所属项目uuid',
  `language` varchar(255) DEFAULT NULL COMMENT '语言（中文/English）',
  `timezone` varchar(64) NOT NULL DEFAULT 'Asia/Shanghai' COMMENT 'timezone',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码，sha256非对称加密后存储',
  `description` varchar(1024) NOT NULL DEFAULT '' COMMENT '描述',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新者',
  `created_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间戳',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8 COMMENT='用户';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'user-ta5c5tsos2wkm8d2qtmvx3vufr2h','role-admin-uuid-01','admin','','','','project-vtdz6jyqux7gxgglzk5fsdiyqy8k','zh_CN','Asia/Shanghai','$2a$10$vl6tEXezDh32egRS5/xkFePmKbS.6D70BaGoR9s180iKJbsdfoaqm','默认管理员','admin','admin',1674959410,1725521379,1675664498,0);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `volume`
--

DROP TABLE IF EXISTS `volume`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `volume` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `volume_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'volume uuid',
  `device_type_id` varchar(255) NOT NULL DEFAULT '' COMMENT '设备类型uuid',
  `volume_name` varchar(100) NOT NULL DEFAULT '' COMMENT '卷名称',
  `volume_type` varchar(255) NOT NULL DEFAULT '' COMMENT '卷类型：系统卷，数据卷',
  `disk_type` varchar(16) NOT NULL DEFAULT '' COMMENT '硬盘类型（SSD,HDD）',
  `interface_type` varchar(16) NOT NULL DEFAULT '' COMMENT '接口类型（SATA,SAS,NVME,不限制）',
  `volume_size` varchar(50) NOT NULL DEFAULT '0' COMMENT '单盘大小（最小容量）',
  `volume_unit` varchar(16) NOT NULL DEFAULT '0' COMMENT '硬盘单位（GB,TB）',
  `volume_amount` int(11) NOT NULL DEFAULT '0' COMMENT '硬盘数量（最低块数）',
  `created_by` varchar(255) NOT NULL DEFAULT '' COMMENT '创建者',
  `updated_by` varchar(255) NOT NULL DEFAULT '' COMMENT '更新者',
  `created_time` int(255) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `deleted_time` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `i_device_type_type` (`volume_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=722 DEFAULT CHARSET=utf8 COMMENT='卷管理表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `volume`
--

LOCK TABLES `volume` WRITE;
/*!40000 ALTER TABLE `volume` DISABLE KEYS */;
/*!40000 ALTER TABLE `volume` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `web_message`
--

DROP TABLE IF EXISTS `web_message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `web_message` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `message_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'message_id',
  `user_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'user_id',
  `resource_type` varchar(255) NOT NULL DEFAULT '' COMMENT '资源类型',
  `result` varchar(255) NOT NULL DEFAULT '' COMMENT '操作结果',
  `finish_time` int(11) NOT NULL DEFAULT '0' COMMENT '完成时间戳',
  `has_read` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0未读, 1已读',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除0未删除 1已删除',
  `message_type` varchar(255) NOT NULL DEFAULT '',
  `message_sub_type` varchar(255) NOT NULL DEFAULT '',
  `instance_id` varchar(255) NOT NULL DEFAULT '',
  `sn` varchar(255) NOT NULL DEFAULT '',
  `detail` varchar(4096) NOT NULL DEFAULT '',
  `fault_type` varchar(255) NOT NULL DEFAULT '' COMMENT 'oob-alert spec field',
  `content` varchar(4096) NOT NULL DEFAULT '' COMMENT 'oob-alert spec field',
  `logid` varchar(255) NOT NULL DEFAULT '' COMMENT 'oob-alert spec field',
  `alert_time` int(11) NOT NULL DEFAULT '0' COMMENT 'oob-alert spec field',
  `alert_count` int(11) NOT NULL DEFAULT '0' COMMENT 'oob-alert spec field',
  `license_name` varchar(255) NOT NULL DEFAULT '' COMMENT 'license spec field',
  `license_number` varchar(255) NOT NULL DEFAULT '' COMMENT 'license spec field',
  `license_start_time` int(11) NOT NULL DEFAULT '0' COMMENT 'license spec field',
  `license_end_time` int(11) NOT NULL DEFAULT '0' COMMENT 'license spec field',
  `instance_name` varchar(255) NOT NULL DEFAULT '',
  `user_name` varchar(255) NOT NULL DEFAULT '' COMMENT 'user_name',
  `rule_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'inbond_rule_id',
  `rule_name` varchar(255) NOT NULL DEFAULT '' COMMENT 'inbond_rule_name',
  `is_recover` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0->alert, 1->recover',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `i_message_userid` (`user_id`),
  KEY `i_message_read` (`has_read`),
  KEY `message_type_idx` (`message_type`),
  KEY `message_sub_type_idx` (`message_sub_type`),
  KEY `has_read_idx` (`has_read`),
  KEY `userid_idx` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7648 DEFAULT CHARSET=utf8 COMMENT='web_message';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `web_message`
--

LOCK TABLES `web_message` WRITE;
/*!40000 ALTER TABLE `web_message` DISABLE KEYS */;
/*!40000 ALTER TABLE `web_message` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-09-14  9:10:50
