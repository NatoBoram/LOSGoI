DROP TABLE IF EXISTS `builds`;
CREATE TABLE IF NOT EXISTS `builds`(
    `device` VARCHAR(16) NOT NULL,
    `date` DATE NOT NULL,
    `datetime` DATETIME NOT NULL,
    `filename` VARCHAR(64) NOT NULL,
    `filepath` VARCHAR(128) NOT NULL,
    `sha1` VARCHAR(40) NOT NULL,
    `sha256` VARCHAR(64) NOT NULL,
    `size` INTEGER NOT NULL,
    `type` VARCHAR(16) NOT NULL,
    `version` VARCHAR(8) NOT NULL,
    `ipfs` VARCHAR(128) NOT NULL PRIMARY KEY
);