CREATE TABLE `t_instance_asset`
(`id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
 `inst_id` varchar(32) NOT NULL DEFAULT '' COMMENT '实例id',
 `inst_name` varchar(32) NOT NULL DEFAULT '' COMMENT '实例别名',
 `dms_id` int NOT NULL DEFAULT '0' COMMENT 'DMS中的实例id',
 `inst_type` tinyint NOT NULL DEFAULT '1' COMMENT '实例类型：1->mysql,2->redis,3->mongo,4->postgreSQL,5->SQLserver',
 `version` varchar(10) NOT NULL DEFAULT '' COMMENT '数据库实例版本',
 `role_type` tinyint NOT NULL DEFAULT '1' COMMENT '实例类型：1->Primary,2->Readonly',
 `inst_status` tinyint NOT NULL DEFAULT '1' COMMENT '实例类型：1->Running,2->Closed',
 `port` smallint NOT NULL DEFAULT '0' COMMENT '实例端口',
 `host` varchar(64) NOT NULL DEFAULT '' COMMENT '实例地址',
 `region` tinyint NOT NULL DEFAULT '1' COMMENT '实例类型：1->上海,2->美国',
 `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '生成时间',
 `sys_auto_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '当前时间',
 PRIMARY KEY (`id`),
 UNIQUE KEY `uniq_instid` (`inst_id`) USING BTREE)
 ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='数据库实例资产表';