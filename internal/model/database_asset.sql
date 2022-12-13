CREATE TABLE `t_database_asset`
(`id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
 `db_id` varchar(32) NOT NULL DEFAULT '' COMMENT '数据库id',
 `db_name` varchar(16) NOT NULL DEFAULT '' COMMENT '数据库名',
 `alias_name` varchar(32) NOT NULL DEFAULT '' COMMENT '数据库别名',
 `env` tinyint NOT NULL DEFAULT '1' COMMENT '环境：1->生产,2->预生产,3->开发,4->测试,5->灾备',
 `encoding` tinyint NOT NULL DEFAULT '1' COMMENT '数据库编码：1->utf8mb4,2->utf8',
 `db_type` tinyint NOT NULL DEFAULT '1' COMMENT '数据库类型：1->mysql,2->postgreSQL,3->mongodb,4->SQLserver',
 `inst_id` varchar(32) NOT NULL DEFAULT '' COMMENT '数据库实例id，外键参照t_instance_asset的inst_id',
 `host` varchar(64) NOT NULL DEFAULT '' COMMENT '数据库地址',
 `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '生成时间',
 `sys_auto_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '当前时间',
 PRIMARY KEY (`id`),
 UNIQUE KEY `uniq_env_dbname` (`db_name`,`env`) USING BTREE,
 KEY `idx_dbtype` (`db_type`) USING BTREE,
 KEY `idx_instid` (`inst_id`) USING BTREE)
 ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='数据库资产表';