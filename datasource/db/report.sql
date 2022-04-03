CREATE TABLE statusreport
(
    `reportid`            INT NOT NULL AUTO_INCREMENT,
    `max_temperature`     DOUBLE NOT NULL,
    `min_temperature`     DOUBLE NOT NULL,
    `average_temperature` DOUBLE NOT NULL,
    `max_heartrate`       DOUBLE NOT NULL,
    `min_heartrate`       DOUBLE NOT NULL,
    `average_heartrate`   DOUBLE NOT NULL,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`reportid`)
);