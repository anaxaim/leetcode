CREATE TABLE IF NOT EXISTS Triangle (x INT, y INT, z INT);
TRUNCATE TABLE Triangle;
INSERT INTO Triangle (x, y, z) VALUES ('13', '15', '30');
INSERT INTO Triangle (x, y, z) VALUES ('10', '20', '15');

-- 610. Triangle Judgement

SELECT x, y, z,
       CASE WHEN x+y>z AND y+z>x AND z+x>y THEN 'Yes' ELSE 'No'
           END triangle
FROM Triangle

-- Input
-- Triangle =
-- | x  | y  | z  |
-- | -- | -- | -- |
-- | 13 | 15 | 30 |
-- | 10 | 20 | 15 |
-- Output
-- | x  | y  | z  | triangle |
-- | -- | -- | -- | -------- |
-- | 13 | 15 | 30 | No       |
-- | 10 | 20 | 15 | Yes      |
