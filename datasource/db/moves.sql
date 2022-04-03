USE servercreate;

CREATE TABLE moves
(
    phone   varchar(255) NOT NULL,
    smoke   varchar(255) NOT NULL,
    /*ha*/
    drink   varchar(255) NOT NULL,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
)