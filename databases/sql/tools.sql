-- mysql
-- CREATE DATABASE IF NOT EXISTS tools DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;
--
-- USE tools;
--
-- CREATE TABLE `user_code_detail` (
--   `id` BIGINT NOT NULL AUTO_INCREMENT,
--   `md5sum` VARCHAR(128) NOT NULL,
--   `code_type` VARCHAR(32) NOT NULL,
--   `cmd` TEXT NOT NULL,
--   `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
--   `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
--   PRIMARY KEY (`id`),
--   UNIQUE KEY `KEY_md5sum` (`md5sum`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- sqlite3
-- 用户代码详细表
CREATE TABLE user_code_detail (
   id               INTEGER PRIMARY KEY   AUTOINCREMENT  NOT NULL,
   code_id          CHAR(50)              UNIQUE NOT NULL,
   md5sum           CHAR(50)              NOT NULL,
   code_type        CHAR(10)              NOT NULL,
   insert_type      CHAR(10)              NOT NULL,
   cmd              TEXT                  NOT NULL,
   run_result       INT                   DEFAULT 0,
   run_result_info  TEXT                  NOT NULL  DEFAULT "",
   client_ip        CHAR(20)              NOT NULL,
   user_agent       CHAR(200)             NOT NULL,
   create_time      DATETIME              DEFAULT (datetime('now', 'localtime')),
   update_time      DATETIME              DEFAULT (datetime('now', 'localtime')),
   ext_a            TEXT                  NOT NULL  DEFAULT "",
   UNIQUE (md5sum, code_type)
);

-- 用户分享表
CREATE TABLE user_code_share_detail (
  id                INTEGER PRIMARY KEY   AUTOINCREMENT  NOT NULL,
  detail_id         INTEGER               NOT NULL,
  share_id          CHAR(50)              UNIQUE NOT NULL,
  view_count        INTEGER               DEFAULT 0,
  create_time       DATETIME              DEFAULT (datetime('now', 'localtime')),
  ext_a             TEXT                  NOT NULL  DEFAULT ""
);