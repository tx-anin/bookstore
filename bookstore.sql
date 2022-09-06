/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50730
 Source Host           : localhost:3306
 Source Schema         : bookstore

 Target Server Type    : MySQL
 Target Server Version : 50730
 File Encoding         : 65001

 Date: 06/09/2022 15:08:56
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for books
-- ----------------------------
DROP TABLE IF EXISTS `books`;
CREATE TABLE `books` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  `author` varchar(100) NOT NULL,
  `price` float(100,0) NOT NULL,
  `sales` int(100) NOT NULL,
  `stock` int(100) NOT NULL,
  `img_path` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of books
-- ----------------------------
BEGIN;
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (1, '解忧杂货店', '东野圭吾', 27, 108, 92, 'static/img/解忧杂货店.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (2, '边城', '沈从文', 23, 111, 89, 'static/img/边城.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (3, '中国哲学史', '冯友兰', 44, 109, 91, 'static/img/中国哲学史.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (4, '忽然七日', ' 劳伦', 19, 111, 89, 'static/img/忽然七日.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (5, '苏东坡传', '林语堂', 19, 100, 100, 'static/img/苏东坡传.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (6, '百年孤独', '马尔克斯', 29, 100, 100, 'static/img/百年孤独.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (7, '扶桑', '严歌苓', 20, 102, 98, 'static/img/扶桑.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (8, '给孩子的诗', '北岛', 22, 102, 98, 'static/img/给孩子的诗.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (9, '为奴十二年', '所罗门', 16, 101, 99, 'static/img/为奴十二年.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (10, '平凡的世界', '路遥', 55, 101, 99, 'static/img/平凡的世界.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (11, '悟空传', '今何在', 14, 103, 97, 'static/img/悟空传.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (12, '硬派健身', '斌卡', 31, 101, 99, 'static/img/硬派健身.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (13, '从晚清到民国', '唐德刚', 40, 100, 100, 'static/img/从晚清到民国.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (14, '三体', '刘慈欣', 56, 100, 100, 'static/img/三体.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (15, '看见', '柴静', 19, 102, 98, 'static/img/看见.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (16, '活着', '余华', 11, 100, 100, 'static/img/活着.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (17, '小王子', '安托万', 19, 100, 100, 'static/img/小王子.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (18, '我们仨', '杨绛', 11, 100, 100, 'static/img/我们仨.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (19, '生命不息,折腾不止', '罗永浩', 25, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (20, '皮囊', '蔡崇达', 24, 100, 100, 'static/img/皮囊.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (21, '恰到好处的幸福', '毕淑敏', 16, 100, 100, 'static/img/恰到好处的幸福.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (22, '大数据预测', '埃里克', 37, 101, 99, 'static/img/大数据预测.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (23, '人月神话', '布鲁克斯', 56, 100, 100, 'static/img/人月神话.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (24, 'C语言入门经典', '霍尔顿', 45, 101, 99, 'static/img/C语言入门经典.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (25, '数学之美', '吴军', 30, 100, 100, 'static/img/数学之美.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (26, 'Java编程思想', '埃史尔', 70, 100, 100, 'static/img/Java编程思想.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (27, '设计模式之禅', '秦小波', 20, 100, 100, 'static/img/设计模式之禅.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (28, '图解机器学习', '杉山将', 34, 100, 100, 'static/img/图解机器学习.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (29, '艾伦图灵传', '安德鲁', 47, 100, 100, 'static/img/艾伦图灵传.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (30, '教父', '马里奥普佐', 29, 100, 100, 'static/img/教父.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (40, 'Go语言学习笔记', '雨痕', 51, 100, 33, '/static/img/default.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (43, 'go语言基础', 'l1ng14', 30, 10, 111, '/static/img/default.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (44, 'go语言基础', 'l1ng14', 30, 10, 120, '/static/img/default.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (45, '三国演义第二季', '罗小中', 9, 900, 100, '/static/img/default.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (49, '三个女人和一百零五个男人的故事', '罗贯中', 67, 500, 500, '/static/img/default.jpg');
INSERT INTO `books` (`id`, `title`, `author`, `price`, `sales`, `stock`, `img_path`) VALUES (52, '西游记后传', '吴承恩', 40, 100, 10, '/static/img/default.jpg');
COMMIT;

-- ----------------------------
-- Table structure for cart_items
-- ----------------------------
DROP TABLE IF EXISTS `cart_items`;
CREATE TABLE `cart_items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `COUNT` int(11) NOT NULL,
  `amount` double(11,2) NOT NULL,
  `book_id` int(11) NOT NULL,
  `cart_id` varchar(100) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `book_id` (`book_id`) USING BTREE,
  KEY `cart_id` (`cart_id`) USING BTREE,
  CONSTRAINT `cart_items_ibfk_1` FOREIGN KEY (`book_id`) REFERENCES `books` (`id`),
  CONSTRAINT `cart_items_ibfk_2` FOREIGN KEY (`cart_id`) REFERENCES `carts` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of cart_items
-- ----------------------------
BEGIN;
INSERT INTO `cart_items` (`id`, `COUNT`, `amount`, `book_id`, `cart_id`) VALUES (20, 9, 207.00, 2, 'ad282fc9-6a38-4bb2-5b68-0955edad023b');
INSERT INTO `cart_items` (`id`, `COUNT`, `amount`, `book_id`, `cart_id`) VALUES (21, 9, 396.00, 3, 'ad282fc9-6a38-4bb2-5b68-0955edad023b');
INSERT INTO `cart_items` (`id`, `COUNT`, `amount`, `book_id`, `cart_id`) VALUES (22, 9, 171.00, 4, 'ad282fc9-6a38-4bb2-5b68-0955edad023b');
COMMIT;

-- ----------------------------
-- Table structure for carts
-- ----------------------------
DROP TABLE IF EXISTS `carts`;
CREATE TABLE `carts` (
  `id` varchar(100) NOT NULL,
  `total_count` int(11) NOT NULL,
  `total_amount` double(11,2) NOT NULL,
  `user_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  CONSTRAINT `carts_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of carts
-- ----------------------------
BEGIN;
INSERT INTO `carts` (`id`, `total_count`, `total_amount`, `user_id`) VALUES ('ad282fc9-6a38-4bb2-5b68-0955edad023b', 27, 774.00, 24);
COMMIT;

-- ----------------------------
-- Table structure for order_items
-- ----------------------------
DROP TABLE IF EXISTS `order_items`;
CREATE TABLE `order_items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `COUNT` int(11) NOT NULL,
  `amount` double(11,2) NOT NULL,
  `title` varchar(100) NOT NULL,
  `author` varchar(100) NOT NULL,
  `price` double(11,2) NOT NULL,
  `img_path` varchar(100) NOT NULL,
  `order_id` varchar(100) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `order_id` (`order_id`) USING BTREE,
  CONSTRAINT `order_items_ibfk_1` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=58 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of order_items
-- ----------------------------
BEGIN;
INSERT INTO `order_items` (`id`, `COUNT`, `amount`, `title`, `author`, `price`, `img_path`, `order_id`) VALUES (40, 1, 27.00, '解忧杂货店', '东野圭吾', 27.00, 'static/img/解忧杂货店.jpg', '6a34bfb8-5fad-4795-5bf8-56e63bd123ff');
INSERT INTO `order_items` (`id`, `COUNT`, `amount`, `title`, `author`, `price`, `img_path`, `order_id`) VALUES (41, 1, 23.00, '边城', '沈从文', 23.00, 'static/img/边城.jpg', '6a34bfb8-5fad-4795-5bf8-56e63bd123ff');
INSERT INTO `order_items` (`id`, `COUNT`, `amount`, `title`, `author`, `price`, `img_path`, `order_id`) VALUES (42, 1, 44.00, '中国哲学史', '冯友兰', 44.00, 'static/img/中国哲学史.jpg', '6a34bfb8-5fad-4795-5bf8-56e63bd123ff');
INSERT INTO `order_items` (`id`, `COUNT`, `amount`, `title`, `author`, `price`, `img_path`, `order_id`) VALUES (43, 1, 19.00, '忽然七日', ' 劳伦', 19.00, 'static/img/忽然七日.jpg', '6a34bfb8-5fad-4795-5bf8-56e63bd123ff');
INSERT INTO `order_items` (`id`, `COUNT`, `amount`, `title`, `author`, `price`, `img_path`, `order_id`) VALUES (44, 1, 27.00, '解忧杂货店', '东野圭吾', 27.00, 'static/img/解忧杂货店.jpg', 'df90a33f-f3d5-4366-41aa-d8306243ae9c');
INSERT INTO `order_items` (`id`, `COUNT`, `amount`, `title`, `author`, `price`, `img_path`, `order_id`) VALUES (45, 1, 23.00, '边城', '沈从文', 23.00, 'static/img/边城.jpg', 'df90a33f-f3d5-4366-41aa-d8306243ae9c');
INSERT INTO `order_items` (`id`, `COUNT`, `amount`, `title`, `author`, `price`, `img_path`, `order_id`) VALUES (46, 1, 44.00, '中国哲学史', '冯友兰', 44.00, 'static/img/中国哲学史.jpg', 'df90a33f-f3d5-4366-41aa-d8306243ae9c');
INSERT INTO `order_items` (`id`, `COUNT`, `amount`, `title`, `author`, `price`, `img_path`, `order_id`) VALUES (47, 1, 19.00, '忽然七日', ' 劳伦', 19.00, 'static/img/忽然七日.jpg', 'df90a33f-f3d5-4366-41aa-d8306243ae9c');
INSERT INTO `order_items` (`id`, `COUNT`, `amount`, `title`, `author`, `price`, `img_path`, `order_id`) VALUES (48, 1, 300.00, '三国', '罗三胖', 300.00, '/static/img/default.jpg', '6170a6ad-99bb-4d3d-4b5f-2272a4f72c53');
INSERT INTO `order_items` (`id`, `COUNT`, `amount`, `title`, `author`, `price`, `img_path`, `order_id`) VALUES (49, 1, 100.00, '西游', '吴二傻', 100.00, '/static/img/default.jpg', '6170a6ad-99bb-4d3d-4b5f-2272a4f72c53');
INSERT INTO `order_items` (`id`, `COUNT`, `amount`, `title`, `author`, `price`, `img_path`, `order_id`) VALUES (50, 2, 46.00, '边城', '沈从文', 23.00, 'static/img/边城.jpg', 'b3f09f77-975b-4922-65ad-488942e71f8d');
INSERT INTO `order_items` (`id`, `COUNT`, `amount`, `title`, `author`, `price`, `img_path`, `order_id`) VALUES (51, 1, 23.00, '边城', '沈从文', 23.00, 'static/img/边城.jpg', '78a56b4c-676c-463a-7ef3-7b079647224d');
INSERT INTO `order_items` (`id`, `COUNT`, `amount`, `title`, `author`, `price`, `img_path`, `order_id`) VALUES (52, 1, 44.00, '中国哲学史', '冯友兰', 44.00, 'static/img/中国哲学史.jpg', '78a56b4c-676c-463a-7ef3-7b079647224d');
INSERT INTO `order_items` (`id`, `COUNT`, `amount`, `title`, `author`, `price`, `img_path`, `order_id`) VALUES (53, 1, 19.00, '忽然七日', ' 劳伦', 19.00, 'static/img/忽然七日.jpg', '78a56b4c-676c-463a-7ef3-7b079647224d');
INSERT INTO `order_items` (`id`, `COUNT`, `amount`, `title`, `author`, `price`, `img_path`, `order_id`) VALUES (54, 1, 27.00, '解忧杂货店', '东野圭吾', 27.00, 'static/img/解忧杂货店.jpg', 'bf7bfccf-8767-4062-6f53-b23f683f3092');
INSERT INTO `order_items` (`id`, `COUNT`, `amount`, `title`, `author`, `price`, `img_path`, `order_id`) VALUES (55, 1, 23.00, '边城', '沈从文', 23.00, 'static/img/边城.jpg', 'bf7bfccf-8767-4062-6f53-b23f683f3092');
INSERT INTO `order_items` (`id`, `COUNT`, `amount`, `title`, `author`, `price`, `img_path`, `order_id`) VALUES (56, 1, 44.00, '中国哲学史', '冯友兰', 44.00, 'static/img/中国哲学史.jpg', 'bf7bfccf-8767-4062-6f53-b23f683f3092');
INSERT INTO `order_items` (`id`, `COUNT`, `amount`, `title`, `author`, `price`, `img_path`, `order_id`) VALUES (57, 1, 19.00, '忽然七日', ' 劳伦', 19.00, 'static/img/忽然七日.jpg', 'bf7bfccf-8767-4062-6f53-b23f683f3092');
COMMIT;

-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `id` varchar(100) NOT NULL,
  `create_time` varchar(100) NOT NULL,
  `total_count` int(11) NOT NULL,
  `total_amount` double(11,2) NOT NULL,
  `state` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of orders
-- ----------------------------
BEGIN;
INSERT INTO `orders` (`id`, `create_time`, `total_count`, `total_amount`, `state`, `user_id`) VALUES ('6170a6ad-99bb-4d3d-4b5f-2272a4f72c53', '2022-04-11 20:48:44', 2, 400.00, 2, 26);
INSERT INTO `orders` (`id`, `create_time`, `total_count`, `total_amount`, `state`, `user_id`) VALUES ('6a34bfb8-5fad-4795-5bf8-56e63bd123ff', '2022-04-11 20:42:33', 4, 113.00, 2, 26);
INSERT INTO `orders` (`id`, `create_time`, `total_count`, `total_amount`, `state`, `user_id`) VALUES ('78a56b4c-676c-463a-7ef3-7b079647224d', '2022-09-06 09:37:39', 3, 86.00, 0, 26);
INSERT INTO `orders` (`id`, `create_time`, `total_count`, `total_amount`, `state`, `user_id`) VALUES ('b3f09f77-975b-4922-65ad-488942e71f8d', '2022-09-06 09:37:18', 2, 46.00, 0, 26);
INSERT INTO `orders` (`id`, `create_time`, `total_count`, `total_amount`, `state`, `user_id`) VALUES ('bf7bfccf-8767-4062-6f53-b23f683f3092', '2022-09-06 09:39:23', 4, 113.00, 0, 26);
INSERT INTO `orders` (`id`, `create_time`, `total_count`, `total_amount`, `state`, `user_id`) VALUES ('df90a33f-f3d5-4366-41aa-d8306243ae9c', '2022-04-11 20:43:06', 4, 113.00, 1, 18);
COMMIT;

-- ----------------------------
-- Table structure for sessions
-- ----------------------------
DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions` (
  `session_id` varchar(100) NOT NULL,
  `username` varchar(100) NOT NULL,
  `user_id` int(11) NOT NULL,
  PRIMARY KEY (`session_id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  CONSTRAINT `sessions_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sessions
-- ----------------------------
BEGIN;
INSERT INTO `sessions` (`session_id`, `username`, `user_id`) VALUES ('1b767cde-e804-4624-5c2a-20cf207aee97', 'admin', 26);
INSERT INTO `sessions` (`session_id`, `username`, `user_id`) VALUES ('3190b06768295c752f5ee1a2eb7f401c', 'lili', 1);
INSERT INTO `sessions` (`session_id`, `username`, `user_id`) VALUES ('4cdb2ce2-ac18-469b-444c-eeb6bd6a736a', 'admin', 26);
INSERT INTO `sessions` (`session_id`, `username`, `user_id`) VALUES ('ad3dc6fe8d7e714180d8186bae2a0eb2', 'lilian', 14);
INSERT INTO `sessions` (`session_id`, `username`, `user_id`) VALUES ('ba117a576b68b683be736ca3a7bac600', 'lili', 1);
INSERT INTO `sessions` (`session_id`, `username`, `user_id`) VALUES ('c50e680dc5d4185a707a293263a2bc33', 'lili', 1);
INSERT INTO `sessions` (`session_id`, `username`, `user_id`) VALUES ('eba9dc646c47ee3e05464c95a062c449', 'lili', 1);
INSERT INTO `sessions` (`session_id`, `username`, `user_id`) VALUES ('f71e425b-87aa-4c6d-7b0b-4c70d0c8d71a', 'admin6', 24);
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `PASSWORD` varchar(100) NOT NULL,
  `email` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`,`username`) USING BTREE,
  UNIQUE KEY `username` (`username`) USING BTREE,
  KEY `id` (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` (`id`, `username`, `PASSWORD`, `email`) VALUES (1, 'lili', 'e10adc3949ba59abbe56e057f20f883e', '10010@qq.com');
INSERT INTO `users` (`id`, `username`, `PASSWORD`, `email`) VALUES (8, 'chan', '21218cca77804d2ba1922c33e0151105', '111@qq.com');
INSERT INTO `users` (`id`, `username`, `PASSWORD`, `email`) VALUES (9, 'kiki', '69f8ea31de0c00502b2ae571fbab1f95', '10010@136.com');
INSERT INTO `users` (`id`, `username`, `PASSWORD`, `email`) VALUES (11, 'Chan Mark', '8ddcff3a80f4189ca1c9d4d902c3c909', 'asfsdfsfsfs11@foxmail.com');
INSERT INTO `users` (`id`, `username`, `PASSWORD`, `email`) VALUES (12, 'li xiao long', '25d55ad283aa400af464c76d713c07ad', 'l34635464213114@foxmail.com');
INSERT INTO `users` (`id`, `username`, `PASSWORD`, `email`) VALUES (13, 'li1ng14', '091d2f632b9187475ed9ab0adac83168', 'l1ng14@foxmail.com');
INSERT INTO `users` (`id`, `username`, `PASSWORD`, `email`) VALUES (14, 'lilian', '25d55ad283aa400af464c76d713c07ad', '11111111@qq.com');
INSERT INTO `users` (`id`, `username`, `PASSWORD`, `email`) VALUES (16, 'jack', '25d55ad283aa400af464c76d713c07ad', '55555@136.com');
INSERT INTO `users` (`id`, `username`, `PASSWORD`, `email`) VALUES (18, 'admin2', '123456', 'admin@qq.com');
INSERT INTO `users` (`id`, `username`, `PASSWORD`, `email`) VALUES (20, 'admin3', '123456', 'admin3@qq.com');
INSERT INTO `users` (`id`, `username`, `PASSWORD`, `email`) VALUES (21, 'admin4', '123456', 'admin@qq.com');
INSERT INTO `users` (`id`, `username`, `PASSWORD`, `email`) VALUES (23, 'admin5', '123456', 'admin5@qq.com');
INSERT INTO `users` (`id`, `username`, `PASSWORD`, `email`) VALUES (24, 'admin6', '123456', 'admin6@qq.com');
INSERT INTO `users` (`id`, `username`, `PASSWORD`, `email`) VALUES (26, 'admin', '123456', 'admin@qq.com');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
