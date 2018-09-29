DROP DATABASE IF EXISTS demodb;
-- CREATE DATABASE [IF NOT EXISTS] databaseName -- somehow I was getting an error using this command, so now I'm dropping and creating a new one
CREATE DATABASE demodb;

ALTER DATABASE demodb
-- CHARACTER
-- SET = utf8
-- COLLATE = utf8

SET NAMES
'utf8';
SET CHARACTER
SET utf8;

USE demodb;

CREATE TABLE ibovespa
(
    paperName VARCHAR(255) NOT NULL,
    companyName VARCHAR(255) NOT NULL,
    dailyRate VARCHAR(20) NOT NULL,
    marketValue FLOAT NOT NULL,
);