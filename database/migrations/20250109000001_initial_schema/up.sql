-- 爬虫目标表：存储需要爬取的网站和规则
CREATE TABLE `spider_target` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `name` varchar(100) NOT NULL COMMENT '目标名称',
    `url` varchar(500) NOT NULL COMMENT '目标URL',
    `data_type` varchar(50) NOT NULL COMMENT '数据类型：admission_score, college_info, etc.',
    `crawl_frequency` int(11) NOT NULL DEFAULT '24' COMMENT '爬取频率(小时)',
    `parse_rules` json NOT NULL COMMENT '解析规则(JSON格式)',
    `last_crawl_time` datetime DEFAULT NULL COMMENT '最后爬取时间',
    `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态：1-启用，0-禁用',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_data_type` (`data_type`),
    KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='爬虫目标表';

-- 爬虫任务表：记录每次爬取任务
CREATE TABLE `spider_task` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `target_id` bigint(20) NOT NULL COMMENT '目标ID',
    `task_id` varchar(100) NOT NULL COMMENT 'Asynq任务ID',
    `status` varchar(20) NOT NULL DEFAULT 'pending' COMMENT '状态：pending, running, completed, failed, cancelled',
    `start_time` datetime DEFAULT NULL COMMENT '开始时间',
    `end_time` datetime DEFAULT NULL COMMENT '结束时间',
    `total_items` int(11) DEFAULT '0' COMMENT '总条目数',
    `success_count` int(11) DEFAULT '0' COMMENT '成功条目数',
    `error_message` text COMMENT '错误信息',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_target_id` (`target_id`),
    KEY `idx_task_id` (`task_id`),
    KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='爬虫任务表';

-- 录取分数线表：核心数据表
CREATE TABLE `admission_score` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `year` int(11) NOT NULL COMMENT '年份',
    `province` varchar(20) NOT NULL COMMENT '省份',
    `college_code` varchar(20) NOT NULL COMMENT '学校代码',
    `college_name` varchar(100) NOT NULL COMMENT '学校名称',
    `major_code` varchar(20) DEFAULT NULL COMMENT '专业代码',
    `major_name` varchar(100) DEFAULT NULL COMMENT '专业名称',
    `batch` varchar(50) NOT NULL COMMENT '录取批次',
    `score_type` tinyint(4) NOT NULL COMMENT '分数类型：1-文科，2-理科，3-综合改革',
    `min_score` int(11) DEFAULT NULL COMMENT '最低分',
    `min_rank` int(11) DEFAULT NULL COMMENT '最低位次',
    `avg_score` int(11) DEFAULT NULL COMMENT '平均分',
    `enrollment_count` int(11) DEFAULT NULL COMMENT '录取人数',
    `data_source` varchar(200) NOT NULL COMMENT '数据来源URL',
    `data_quality` tinyint(4) DEFAULT '1' COMMENT '数据质量评分：1-5',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_year_province_college_major_batch_score_type` (`year`, `province`, `college_code`, `major_code`, `batch`, `score_type`),
    KEY `idx_query` (`year`, `province`, `college_name`, `batch`, `score_type`),
    KEY `idx_college` (`college_code`),
    KEY `idx_major` (`major_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='录取分数线表';

-- 学校信息表
CREATE TABLE `college_info` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `college_code` varchar(20) NOT NULL COMMENT '学校代码',
    `college_name` varchar(100) NOT NULL COMMENT '学校名称',
    `province` varchar(20) NOT NULL COMMENT '所在省份',
    `city` varchar(20) DEFAULT NULL COMMENT '所在城市',
    `level` varchar(20) DEFAULT NULL COMMENT '学校等级：985, 211, 双一流, 普通本科',
    `type` varchar(20) DEFAULT NULL COMMENT '学校类型：综合, 理工, 师范, 农业, etc.',
    `nature` varchar(20) DEFAULT NULL COMMENT '办学性质：公办, 民办',
    `website` varchar(200) DEFAULT NULL COMMENT '学校官网',
    `description` text COMMENT '学校简介',
    `ranking` int(11) DEFAULT NULL COMMENT '学校排名',
    `employment_rate` decimal(5,2) DEFAULT NULL COMMENT '就业率',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_college_code` (`college_code`),
    KEY `idx_province` (`province`),
    KEY `idx_level` (`level`),
    KEY `idx_type` (`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='学校信息表';

-- 专业信息表
CREATE TABLE `major_info` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `major_code` varchar(20) NOT NULL COMMENT '专业代码',
    `major_name` varchar(100) NOT NULL COMMENT '专业名称',
    `category` varchar(50) DEFAULT NULL COMMENT '专业大类',
    `subcategory` varchar(50) DEFAULT NULL COMMENT '专业小类',
    `degree` varchar(20) DEFAULT NULL COMMENT '学位类型：学士, 硕士, 博士',
    `duration` int(11) DEFAULT NULL COMMENT '学制(年)',
    `description` text COMMENT '专业简介',
    `employment_direction` text COMMENT '就业方向',
    `salary_range` varchar(50) DEFAULT NULL COMMENT '薪资范围',
    `demand_level` varchar(20) DEFAULT NULL COMMENT '需求程度：高, 中, 低',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_major_code` (`major_code`),
    KEY `idx_category` (`category`),
    KEY `idx_demand_level` (`demand_level`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='专业信息表';

-- 用户表
CREATE TABLE `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `open_id` varchar(100) NOT NULL COMMENT '微信OpenID',
    `union_id` varchar(100) DEFAULT NULL COMMENT '微信UnionID',
    `nickname` varchar(50) DEFAULT NULL COMMENT '昵称',
    `avatar` varchar(500) DEFAULT NULL COMMENT '头像URL',
    `gender` tinyint(4) DEFAULT NULL COMMENT '性别：0-未知，1-男，2-女',
    `province` varchar(20) DEFAULT NULL COMMENT '省份',
    `city` varchar(20) DEFAULT NULL COMMENT '城市',
    `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态：0-未激活，1-活跃，2-封禁',
    `last_login_time` datetime DEFAULT NULL COMMENT '最后登录时间',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_open_id` (`open_id`),
    KEY `idx_union_id` (`union_id`),
    KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 用户档案表
CREATE TABLE `user_profile` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `user_id` bigint(20) NOT NULL COMMENT '用户ID',
    `graduation_year` int(11) DEFAULT NULL COMMENT '毕业年份',
    `province` varchar(20) DEFAULT NULL COMMENT '高考省份',
    `score_type` tinyint(4) DEFAULT NULL COMMENT '分数类型：1-文科，2-理科，3-综合改革',
    `subjects` varchar(100) DEFAULT NULL COMMENT '选考科目，逗号分隔',
    `total_score` int(11) DEFAULT NULL COMMENT '总分',
    `rank` int(11) DEFAULT NULL COMMENT '位次',
    `target_college` varchar(100) DEFAULT NULL COMMENT '目标学校',
    `target_major` varchar(100) DEFAULT NULL COMMENT '目标专业',
    `interest_tags` varchar(500) DEFAULT NULL COMMENT '兴趣标签，逗号分隔',
    `personality_type` varchar(20) DEFAULT NULL COMMENT '性格类型',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_id` (`user_id`),
    KEY `idx_province` (`province`),
    KEY `idx_score_type` (`score_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户档案表';

-- 学长分享表
CREATE TABLE `senior_share` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `user_id` bigint(20) NOT NULL COMMENT '分享者用户ID',
    `college_code` varchar(20) NOT NULL COMMENT '学校代码',
    `college_name` varchar(100) NOT NULL COMMENT '学校名称',
    `major_code` varchar(20) DEFAULT NULL COMMENT '专业代码',
    `major_name` varchar(100) DEFAULT NULL COMMENT '专业名称',
    `share_type` varchar(20) NOT NULL COMMENT '分享类型：experience, advice, interview',
    `title` varchar(200) NOT NULL COMMENT '标题',
    `content` text NOT NULL COMMENT '内容',
    `tags` varchar(500) DEFAULT NULL COMMENT '标签，逗号分隔',
    `is_anonymous` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否匿名：0-否，1-是',
    `view_count` int(11) NOT NULL DEFAULT '0' COMMENT '浏览数',
    `like_count` int(11) NOT NULL DEFAULT '0' COMMENT '点赞数',
    `comment_count` int(11) NOT NULL DEFAULT '0' COMMENT '评论数',
    `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态：0-待审核，1-已发布，2-已隐藏，3-已删除',
    `published_at` datetime DEFAULT NULL COMMENT '发布时间',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`),
    KEY `idx_college_code` (`college_code`),
    KEY `idx_major_code` (`major_code`),
    KEY `idx_share_type` (`share_type`),
    KEY `idx_status` (`status`),
    KEY `idx_published_at` (`published_at`),
    FULLTEXT KEY `ft_content` (`title`, `content`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='学长分享表';

-- 学长分享点赞表
CREATE TABLE `senior_share_like` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `user_id` bigint(20) NOT NULL COMMENT '用户ID',
    `share_id` bigint(20) NOT NULL COMMENT '分享ID',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_share` (`user_id`, `share_id`),
    KEY `idx_share_id` (`share_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='学长分享点赞表';

-- 职业测评表
CREATE TABLE `career_assessment` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `user_id` bigint(20) NOT NULL COMMENT '用户ID',
    `assessment_type` varchar(20) NOT NULL COMMENT '测评类型：mbti, holland, etc.',
    `result` json NOT NULL COMMENT '测评结果(JSON)',
    `score_details` json DEFAULT NULL COMMENT '分数详情(JSON)',
    `recommendations` text COMMENT '职业建议',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`),
    KEY `idx_assessment_type` (`assessment_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='职业测评表';

-- 选科规划表
CREATE TABLE `subject_plan` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `province` varchar(20) NOT NULL COMMENT '省份',
    `score_type` tinyint(4) NOT NULL COMMENT '分数类型：1-文科，2-理科，3-综合改革',
    `subject_combination` varchar(100) NOT NULL COMMENT '科目组合',
    `recommended_majors` text NOT NULL COMMENT '推荐专业，逗号分隔',
    `success_rate` decimal(5,2) DEFAULT NULL COMMENT '成功率',
    `avg_score_required` int(11) DEFAULT NULL COMMENT '平均所需分数',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_province_score_type_subjects` (`province`, `score_type`, `subject_combination`),
    KEY `idx_province` (`province`),
    KEY `idx_score_type` (`score_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='选科规划表';
