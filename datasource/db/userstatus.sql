USE servercreate;
CREATE TABLE user_status
(
    `id` INT NOT NULL AUTO_INCREMENT,
    `mqttid` INT NOT NULL ,
    `heart_rate` DOUBLE,
    temperature DOUBLE,
    `alcohol_strength` DOUBLE,
    `dht` DOUBLE,
    `gy` DOUBLE,
    `updated_time` DATETIME NOT NULL,
    PRIMARY KEY (`id`, `mqttid`)
);