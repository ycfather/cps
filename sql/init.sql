CREATE TABLE IF NOT EXISTS `cps_admin_user` (
  `id`              SMALLINT    NOT NULL AUTO_INCREMENT,
  `email`           VARCHAR(50) NOT NULL,
  `status`          TINYINT     NOT NULL DEFAULT 0,
  `last_login_time` INT         NOT NULL,
  PRIMARY KEY (`id`)
);

ALTER TABLE cps_admin_user
  ADD UNIQUE (email);