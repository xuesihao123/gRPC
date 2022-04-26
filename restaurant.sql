/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80021
 Source Host           : localhost:3307
 Source Schema         : restaurant

 Target Server Type    : MySQL
 Target Server Version : 80021
 File Encoding         : 65001

 Date: 26/04/2022 20:48:13
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for fh
-- ----------------------------
DROP TABLE IF EXISTS `fh`;
CREATE TABLE `fh`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `having_id` int NULL DEFAULT NULL,
  `food_id` int NULL DEFAULT NULL,
  `food_weight` float NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of fh
-- ----------------------------
INSERT INTO `fh` VALUES (1, 2, 2, 2);

-- ----------------------------
-- Table structure for food
-- ----------------------------
DROP TABLE IF EXISTS `food`;
CREATE TABLE `food`  (
  `food_id` int NOT NULL AUTO_INCREMENT,
  `food_cw` float NULL DEFAULT NULL,
  `food_oil` float NULL DEFAULT NULL,
  `food_energy` float NULL DEFAULT NULL,
  `food_protein` float NULL DEFAULT NULL,
  `food_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `food_type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`food_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of food
-- ----------------------------
INSERT INTO `food` VALUES (2, 1.2, 2.1, 1.2, 131, '食用油', '材料');

-- ----------------------------
-- Table structure for having
-- ----------------------------
DROP TABLE IF EXISTS `having`;
CREATE TABLE `having`  (
  `having_id` int NOT NULL AUTO_INCREMENT,
  `having_price` float NULL DEFAULT NULL,
  `having_name` varchar(25) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `having_cw` float NULL DEFAULT NULL,
  `having_oil` float NULL DEFAULT NULL,
  `having_energy` float NULL DEFAULT NULL,
  `having_protein` float NULL DEFAULT NULL,
  `having_update_time` datetime(0) NULL DEFAULT NULL,
  `having_content` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`having_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of having
-- ----------------------------
INSERT INTO `having` VALUES (1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `having` VALUES (2, 0.1, '油', 0.1, 0.1, 0.1, 0.1, '2022-04-26 14:17:50', '就是油');

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`  (
  `menu_id` int NOT NULL AUTO_INCREMENT,
  `create_date` date NULL DEFAULT NULL,
  `menu_status` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`menu_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES (1, NULL, 'DayDate', '2022-04-26 00:00:00');

-- ----------------------------
-- Table structure for mh
-- ----------------------------
DROP TABLE IF EXISTS `mh`;
CREATE TABLE `mh`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `menu_id` int NULL DEFAULT NULL,
  `having_id` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of mh
-- ----------------------------
INSERT INTO `mh` VALUES (1, 0, 0);
INSERT INTO `mh` VALUES (2, 1, 1);

-- ----------------------------
-- Table structure for oh
-- ----------------------------
DROP TABLE IF EXISTS `oh`;
CREATE TABLE `oh`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_id` int NULL DEFAULT NULL,
  `having_id` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of oh
-- ----------------------------
INSERT INTO `oh` VALUES (3, 0, 0);

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order`  (
  `order_id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NULL DEFAULT NULL,
  `order_tag` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `order_price` float NULL DEFAULT NULL,
  `order_create_time` datetime(0) NULL DEFAULT NULL,
  `order_update_time` datetime(0) NULL DEFAULT NULL,
  `order_status` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `order_having_id` int NULL DEFAULT NULL,
  `order_remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`order_id`) USING BTREE,
  INDEX `having`(`order_having_id`) USING BTREE,
  INDEX `userId`(`user_id`) USING BTREE,
  CONSTRAINT `having` FOREIGN KEY (`order_having_id`) REFERENCES `having` (`having_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `userId` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of order
-- ----------------------------
INSERT INTO `order` VALUES (23, 6, '123', 2, '2022-04-25 22:31:28', '2022-04-25 22:31:28', 'create', NULL, '');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `user_id` int NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `user_password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `user_email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `user_status` bigint NULL DEFAULT NULL,
  `user_pid` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  UNIQUE INDEX `user_email`(`user_email`) USING BTREE,
  UNIQUE INDEX `uix_user_user_email`(`user_email`) USING BTREE,
  INDEX `id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (6, 'xuesihao', '123', '754769243@qq.com', 1, NULL);

SET FOREIGN_KEY_CHECKS = 1;
