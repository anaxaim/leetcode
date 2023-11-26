CREATE TABLE IF NOT EXISTS MyNumbers (num INT);
TRUNCATE TABLE MyNumbers;
INSERT INTO MyNumbers (num) VALUES ('8');
INSERT INTO MyNumbers (num) VALUES ('8');
INSERT INTO MyNumbers (num) VALUES ('3');
INSERT INTO MyNumbers (num) VALUES ('3');
INSERT INTO MyNumbers (num) VALUES ('1');
INSERT INTO MyNumbers (num) VALUES ('4');
INSERT INTO MyNumbers (num) VALUES ('5');
INSERT INTO MyNumbers (num) VALUES ('6');

-- 619. Biggest Single Number

SELECT MAX(num) AS num
FROM (
         SELECT num
         FROM MyNumbers
         GROUP BY num
         HAVING COUNT(num) = 1
     )


-- Input
-- MyNumbers =
-- Open Raw
-- | num |
-- | --- |
-- | 8   |
-- | 8   |
-- | 3   |
-- | 3   |
-- | 1   |
-- | 4   |...
-- Output
-- | num |
-- | --- |
-- | 6   |
