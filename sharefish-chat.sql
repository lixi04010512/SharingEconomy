/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80023
 Source Host           : localhost:3306
 Source Schema         : sharefish

 Target Server Type    : MySQL
 Target Server Version : 80023
 File Encoding         : 65001

 Date: 22/04/2022 23:54:39
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for chat
-- ----------------------------
DROP TABLE IF EXISTS `chat`;
CREATE TABLE `chat`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `addr_from` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `addr_to` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `img_from` varchar(35) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `img_to` varchar(35) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `time` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of chat
-- ----------------------------
INSERT INTO `chat` VALUES (9, '0x95d87f7FE4E710A04D2149B374AD65324ED132fb', '0x8bccBf348cd6d7E9c7aa7F9382DD68066ef69aD4', 'hhhh', 'share/static/picture/user-1.jpg', 'share/static/picture/user-2.jpg', '15:01:01');
INSERT INTO `chat` VALUES (10, '0x95d87f7FE4E710A04D2149B374AD65324ED132fb', '0x8bccBf348cd6d7E9c7aa7F9382DD68066ef69aD4', '哈哈', 'share/static/picture/user-1.jpg', 'share/static/picture/user-2.jpg', '17:28:28');
INSERT INTO `chat` VALUES (11, '0x95d87f7FE4E710A04D2149B374AD65324ED132fb', '0x8bccBf348cd6d7E9c7aa7F9382DD68066ef69aD4', '你好', 'share/static/picture/user-1.jpg', 'share/static/picture/user-2.jpg', '17:41:41');

-- ----------------------------
-- Table structure for chat_list
-- ----------------------------
DROP TABLE IF EXISTS `chat_list`;
CREATE TABLE `chat_list`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `owner` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `addr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `time` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of chat_list
-- ----------------------------
INSERT INTO `chat_list` VALUES (2, '0x8bccBf348cd6d7E9c7aa7F9382DD68066ef69aD4', '0x95d87f7FE4E710A04D2149B374AD65324ED132fb', '张三', 'share/static/picture/user-2.jpg', '21.36');
INSERT INTO `chat_list` VALUES (4, '0x95d87f7FE4E710A04D2149B374AD65324ED132fb', '0x8bccBf348cd6d7E9c7aa7F9382DD68066ef69aD4\r\n', '李四', 'share/static/picture/user-1.jpg', '21:11');

SET FOREIGN_KEY_CHECKS = 1;
