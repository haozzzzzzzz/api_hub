ALTER TABLE `ah_doc` ADD COLUMN `doc_type` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '0:swagger; 1:markdown' AFTER `title`;