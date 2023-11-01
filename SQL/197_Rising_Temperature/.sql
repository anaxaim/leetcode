CREATE TABLE IF NOT EXISTS Weather (id INT, recordDate DATE, temperature INT);
TRUNCATE TABLE Weather;
INSERT INTO Weather (id, recordDate, temperature) values ('1', '2015-01-01', '10');
INSERT INTO Weather (id, recordDate, temperature) values ('2', '2015-01-02', '25');
INSERT INTO Weather (id, recordDate, temperature) values ('3', '2015-01-03', '20');
INSERT INTO Weather (id, recordDate, temperature) values ('4', '2015-01-04', '30');

-- 197. Rising Temperature

SELECT w2.id
FROM Weather w1, Weather w2
WHERE (w2.recordDate - w1.recordDate) = 1 AND w2.temperature > w1.temperature

-- Input
-- Weather =
-- | id | recordDate | temperature |
-- | -- | ---------- | ----------- |
-- | 1  | 2015-01-01 | 10          |
-- | 2  | 2015-01-02 | 25          |
-- | 3  | 2015-01-03 | 20          |
-- | 4  | 2015-01-04 | 30          |
-- Output
-- | id |
-- | -- |
-- | 2  |
-- | 4  |
