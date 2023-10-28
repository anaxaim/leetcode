CREATE TABLE IF NOT EXISTS Activity (player_id INT, device_id INT, event_date DATE, games_played INT);
TRUNCATE TABLE Activity;
INSERT INTO Activity (player_id, device_id, event_date, games_played) VALUES ('1', '2', '2016-03-01', '5');
INSERT INTO Activity (player_id, device_id, event_date, games_played) VALUES ('1', '2', '2016-05-02', '6');
INSERT INTO Activity (player_id, device_id, event_date, games_played) VALUES ('2', '3', '2017-06-25', '1');
INSERT INTO Activity (player_id, device_id, event_date, games_played) VALUES ('3', '1', '2016-03-02', '0');
INSERT INTO Activity (player_id, device_id, event_date, games_played) VALUES ('3', '4', '2018-07-03', '5');

-- 511. Game Play Analysis I

SELECT player_id, MIN(event_date) AS first_login
FROM Activity
GROUP BY player_id

-- Input
-- Activity =
-- | player_id | device_id | event_date | games_played |
-- | --------- | --------- | ---------- | ------------ |
-- | 1         | 2         | 2016-03-01 | 5            |
-- | 1         | 2         | 2016-05-02 | 6            |
-- | 2         | 3         | 2017-06-25 | 1            |
-- | 3         | 1         | 2016-03-02 | 0            |
-- | 3         | 4         | 2018-07-03 | 5            |
--
-- Output
-- | player_id | first_login |
-- | --------- | ----------- |
-- | 3         | 2016-03-02  |
-- | 2         | 2017-06-25  |
-- | 1         | 2016-03-01  |
