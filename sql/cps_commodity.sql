CREATE TABLE IF NOT EXISTS `cps_commodity` (
  `id`                   INT          NOT NULL AUTO_INCREMENT,
  `name`                 VARCHAR(20)  NOT NULL,
  `type`                 TINYINT      NOT NULL,
  `agency_type`          TINYINT      NOT NULL,
  `link_url`             VARCHAR(200) NOT NULL,
  `promotion_image`      VARCHAR(200) NOT NULL,
  `price`                FLOAT        NOT NULL,
  `commission_rate`      FLOAT        NOT NULL,
  `creation_time`        INT          NOT NULL,
  `promotion_time_type`  TINYINT      NOT NULL,
  `promotion_time_start` INT          NOT NULL,
  `promotion_time_end`   INT          NOT NULL,
  `biz_id`               SMALLINT     NOT NULL,
  `status`               TINYINT      NOT NULL,
  PRIMARY KEY (`id`)
);

ALTER TABLE `cps_commodity`
  ADD INDEX idx_name(`name`);
ALTER TABLE `cps_commodity`
  ADD INDEX idx_creation_time(`creation_time`);
ALTER TABLE `cps_commodity`
  ADD INDEX idx_biz_id(`biz_id`);