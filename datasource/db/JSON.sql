USE servercreate;
DELIMITER $$
CREATE PROCEDURE `json_del`()
BEGIN
    #Routine body goes here...
    DECLARE id2 INT;
    DECLARE time DATETIME;
    DECLARE payload2 JSON;
    DECLARE json_length INT;#遍历时存储临时长度的变量
    DECLARE count INT; #本次处理总数
    DECLARE i INT DEFAULT(0); #游标循环变量
    DECLARE j INT default(0); #json 循环变量

    DECLARE getData CURSOR FOR SELECT `id`, `payload`, `arrived` FROM t_mqtt_msg;
    SELECT count(*) INTO count FROM t_mqtt_msg;

    OPEN getData;

    REPEAT
        # 开始解析逻辑
        FETCH getData INTO id2,payload2,time;
        SET j = 0;
        SET json_length = JSON_LENGTH(payload2);

        REPEAT
            # 判断空

            IF id2 OR
               JSON_UNQUOTE(JSON_EXTRACT(payload2,CONCAT('$[',j,'].heart_rate'))) OR
               JSON_UNQUOTE(JSON_EXTRACT(payload2,CONCAT('$[',j,'].temperature'))) OR
               JSON_UNQUOTE(JSON_EXTRACT(payload2,CONCAT('$[',j,'].alcohol_strength'))) OR
               JSON_UNQUOTE(JSON_EXTRACT(payload2,CONCAT('$[',j,'].dht'))) OR
               JSON_UNQUOTE(JSON_EXTRACT(payload2,CONCAT('$[',j,'].gy'))) IS NOT NULL THEN

                INSERT INTO user_status(mqttid,gy,dht, heart_rate, temperature, alcohol_strength,updated_time)
                VALUES (id2,
                        JSON_UNQUOTE(JSON_EXTRACT(payload2,CONCAT('$[',j,'].gy'))),
                        JSON_UNQUOTE(JSON_EXTRACT(payload2,CONCAT('$[',j,'].dht'))),
                        JSON_UNQUOTE(JSON_EXTRACT(payload2,CONCAT('$[',j,'].heart_rate'))),
                        JSON_UNQUOTE(JSON_EXTRACT(payload2,CONCAT('$[',j,'].temperature'))),
                        JSON_UNQUOTE(JSON_EXTRACT(payload2,CONCAT('$[',j,'].alcohol_strength'))),


                        time);
            END IF;
            SET j:= j+1;
        UNTIL j >= json_length END REPEAT ;
        SET i := i+1;
    UNTIL i >= count END REPEAT ;
    #关闭游标
    CLOSE getData;
END $$

DELIMITER ;
CALL json_del();
DROP PROCEDURE json_del;

INSERT INTO statusreport(`max_temperature`,
                         `min_temperature`,
                         `average_temperature`,
                         `max_heartrate`,
                         `min_heartrate`,
                         `average_heartrate`) SELECT MAX(), MIN(),AVG(),MAX(),MIN(),AVG() FROM user_status;