use servercreate;
INSERT INTO statusreport(`max_temperature`,
                         `min_temperature`,
                         `average_temperature`,
                         `max_heartrate`,
                         `min_heartrate`,
                         `average_heartrate`) SELECT MAX(), MIN(),AVG(),MAX(),MIN(),AVG() FROM user_status;