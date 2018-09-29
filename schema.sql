DROP DATABASE IF EXISTS demodb;
-- CREATE DATABASE [IF NOT EXISTS] databaseName -- somehow I was getting an error using this command, so now I'm dropping and creating a new one
CREATE DATABASE demodb;

USE demodb;

CREATE TABLE ibovespa
(
    paperName VARCHAR(255) NOT NULL,
    companyName VARCHAR(255) NOT NULL,
    dailyRate VARCHAR(20) NOT NULL,
    marketValue FLOAT NOT NULL,
);