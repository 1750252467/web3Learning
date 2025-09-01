-- Active: 1756683664904@@127.0.0.1@3306
-- Drop the user table from zero_demo database
USE zero_demo;

-- This will delete all data and structure of the user table
-- Warning: This operation cannot be undone!
DROP TABLE IF EXISTS user;

CREATE TABLE user (
    id bigint AUTO_INCREMENT PRIMARY KEY,
    nickname varchar(255) NULL DEFAULT '' COMMENT 'The nickname',
    age INT
);