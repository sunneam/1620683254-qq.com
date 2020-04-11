CREATE TABLE `hb_product` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `product_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品名称',
  `manual` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '说明书地址',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '详情页地址',
  `content` varchar(3000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '详情',
  `state` tinyint(3) NOT NULL DEFAULT '0' COMMENT '状态，1为正常，0为禁用',
  `adminuser_id` int(11) NOT NULL DEFAULT '0' COMMENT '添加信息的管理员ID',
  `sorts` smallint(3) NOT NULL DEFAULT '0' COMMENT '排序',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `site_id` int(11) NOT NULL DEFAULT '0' COMMENT '网站ID',
  `areas` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '现在只能存一个国家,用于区分不同站点下的国家',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `time` (`created`,`updated`) USING BTREE,
  KEY `area` (`site_id`,`areas`) USING BTREE,
  KEY `state` (`state`,`adminuser_id`,`sorts`) USING BTREE,
  KEY `product` (`product_name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='商品表'