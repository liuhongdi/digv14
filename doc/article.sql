-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2020-11-20 15:35:27
-- 服务器版本： 8.0.22-0ubuntu0.20.10.2
-- PHP 版本： 7.4.9

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `dig`
--

-- --------------------------------------------------------

--
-- 表的结构 `article`
--

CREATE TABLE `article` (
  `articleId` bigint UNSIGNED NOT NULL COMMENT 'id',
  `type` tinyint UNSIGNED NOT NULL DEFAULT '0' COMMENT '类型',
  `subject` varchar(500) NOT NULL DEFAULT '' COMMENT '标题',
  `addTime` datetime NOT NULL DEFAULT '2020-11-19 00:00:00' COMMENT '添加时间',
  `isHot` tinyint UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否热榜,0:非热，1，热',
  `url` varchar(500) NOT NULL DEFAULT '' COMMENT '链接地址',
  `domain` varchar(100) NOT NULL DEFAULT '' COMMENT '域',
  `userId` bigint UNSIGNED NOT NULL DEFAULT '0' COMMENT '推荐用户',
  `approvalStaffId` int UNSIGNED NOT NULL DEFAULT '0' COMMENT '批准员工',
  `digSum` int UNSIGNED NOT NULL DEFAULT '0' COMMENT '被顶的次数',
  `commentSum` int UNSIGNED NOT NULL DEFAULT '0' COMMENT '被评论的次数',
  `isPublish` tinyint NOT NULL DEFAULT '0' COMMENT '0,未发布，1，已发布'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文章表';

--
-- 转存表中的数据 `article`
--

INSERT INTO `article` (`articleId`, `type`, `subject`, `addTime`, `isHot`, `url`, `domain`, `userId`, `approvalStaffId`, `digSum`, `commentSum`, `isPublish`) VALUES
(1, 0, '【最近，南非发现一座大油田】石油和天然气两种重要资源是南非储量的最短板，“贫油”的帽子也一直扣在南非的头上摘不下来。可就在最近，在南非海域进行油气勘探已久的道达尔透露了其新的勘探成果：在距离南非南部海岸约175公里的奥特尼夸盆地的11B/12B地区，再次发现了重要的凝析气，可能蕴藏着大量天然气及原油。', '2020-11-19 00:00:00', 1, 'https://mp.weixin.qq.com/s/1btbmouH-2KuIHUMoucq2w', '', 1, 1, 0, 0, 1),
(2, 1, '让喵喵来开启周五的早晨吧！', '2020-11-19 00:00:00', 0, 'https://m.weibo.cn/status/4573112073720433?', 'm.weibo.cn', 0, 0, 0, 0, 1),
(3, 0, '汤姆·赫兰德、黛茜·雷德利、麦斯·米科尔森、尼克·乔纳斯主演的《混沌漫步》公开预告。影片由《明日边缘》导演道格·里曼执导，暂时定档明年1月22日上映。', '2020-11-19 00:00:00', 1, 'http://news.mtime.com/2020/11/19/1604795.html', 'news.mtime.com', 0, 0, 0, 0, 1),
(4, 1, '扫地机器人这个东西确实方便，大多数时候把房间扫的比较干净，但还有很多边边角角清扫不上，希望厂家能够在app上显示出机器人的路线图，明确告知哪些路段没有照顾到，要不然就再开发一个项目经理机器人，跟在扫地机器人屁股后面监督。//@大窑儿：可以弄个步数排行榜，让机器人们互相点赞，揣摩，攀比，内卷，无意义竞争。', '2020-11-19 00:00:00', 0, '', '', 0, 0, 0, 0, 1),
(5, 0, '【世卫组织建议医生不要使用瑞德西韦治疗新冠】世卫组织指导小组表示，证据显示，瑞德西韦对提高新冠肺炎患者的存活率没有显著影响。这项建议发表在上周五的《英国医学杂志》中。', '2020-11-19 00:00:00', 0, 'https://www.thepaper.cn/newsDetail_forward_10067542', 'thepaper.cn', 0, 0, 0, 0, 1),
(6, 0, '11月19日0—24时，31个省（自治区、直辖市）和新疆生产建设兵团报告新增确诊病例17例，均为境外输入病例（福建6例，上海4例，陕西3例，广东2例，北京1例，四川1例）；无新增死亡病例；新增疑似病例1例，为本土病例（在天津）。', '2020-11-19 00:00:00', 0, 'http://m.news.cctv.com/2020/11/20/ARTIIR9o72TuDF80s6hY2IvD201120.shtml', 'm.news.cctv.com', 0, 0, 0, 0, 1);

--
-- 转储表的索引
--

--
-- 表的索引 `article`
--
ALTER TABLE `article`
  ADD PRIMARY KEY (`articleId`),
  ADD KEY `isPublish` (`isPublish`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `article`
--
ALTER TABLE `article`
  MODIFY `articleId` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id', AUTO_INCREMENT=7;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
