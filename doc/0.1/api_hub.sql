CREATE DATABASE api_hub DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
USE api_hub;

CREATE TABLE ah_doc
(
    doc_id int(10) unsigned NOT NULL AUTO_INCREMENT,
    title varchar(255) NOT NULL COMMENT 'doc title',
    spec_url varchar(500) NOT NULL DEFAULT '' COMMENT 'api specification url',
    spec_content longtext NOT NULL COMMENT 'api specification content',
    category_id int(10) unsigned NOT NULL DEFAULT '1' COMMENT 'category id',
    author_id int(10) unsigned NOT NULL COMMENT 'author id',
    post_status tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT 'post status',
    update_time int(10) unsigned NOT NULL COMMENT 'update time',
    create_time int(10) unsigned NOT NULL COMMENT 'create time',
    PRIMARY KEY (doc_id) USING BTREE,
    KEY idx_category_id (category_id),
    KEY idx_author_id (author_id)
) ENGINE=innodb AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE ah_account
(
    acc_id int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'account id',
    name varchar(32) NOT NULL COMMENT 'account name',
    update_time int(10) unsigned NOT NULL COMMENT 'update time',
    create_time int(10) unsigned NOT NULL COMMENT 'create time',
    PRIMARY KEY (acc_id) USING BTREE
) ENGINE=innodb AUTO_INCREMENT=1 DEFAULT  CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE ah_category
(
    cat_id int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'category id',
    name varchar(32) NOT NULL COMMENT 'category name',
    doc_num int(10) unsigned NOT NULL COMMENT 'document number',
    update_time int(10) unsigned NOT NULL COMMENT 'update time',
    create_time int(10) unsigned NOT NULL COMMENT 'create time',
    PRIMARY  KEY (cat_id) USING BTREE
) ENGINE=innodb AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- CREATE TABLE ah_tag
-- (
--     tag_id int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'tag id',
--     name varchar(32) NOT NULL COMMENT 'tag name',
--     doc_num int(10) unsigned NOT NULL COMMENT 'document number',
--     update_time int(10) unsigned COMMENT 'update time',
--     create_time int(10) unsigned COMMENT 'create time',
--     PRIMARY KEY (tag_id) USING BTREE,
--     UNIQUE KEY `un_name` (`name`)
-- ) ENGINE=innodb AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
--
-- CREATE TABLE ah_doc_tag
-- (
--     doc_id int(10) unsigned NOT NULL COMMENT 'doc id',
--     tag_id int(10) unsigned NOT NULL COMMENT 'tag id',
--     create_time int(10) unsigned NOT NULL COMMENT 'create time',
--     PRIMARY KEY (doc_id, tag_id) USING BTREE,
--     KEY idx_tag_id (tag_id)
-- ) ENGINE=innodb DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- default values
INSERT INTO ah_account (acc_id, name, update_time, create_time) VALUES (1, "admin", 1566891150, 1566891150);
INSERT INTO ah_category (cat_id, name, doc_num, update_time, create_time) VALUES (1, "默认目录", 0, 1566891150, 1566891150);
