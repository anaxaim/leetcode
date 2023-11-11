CREATE TABLE IF NOT EXISTS World (name VARCHAR(255), continent VARCHAR(255), area INT, population INT, gdp BIGINT);
TRUNCATE TABLE World;
INSERT INTO World (name, continent, area, population, gdp) VALUES ('Afghanistan', 'Asia', '652230', '25500100', '20343000000');
INSERT INTO World (name, continent, area, population, gdp) VALUES ('Albania', 'Europe', '28748', '2831741', '12960000000');
INSERT INTO World (name, continent, area, population, gdp) VALUES ('Algeria', 'Africa', '2381741', '37100000', '188681000000');
INSERT INTO World (name, continent, area, population, gdp) VALUES ('Andorra', 'Europe', '468', '78115', '3712000000');
INSERT INTO World (name, continent, area, population, gdp) VALUES ('Angola', 'Africa', '1246700', '20609294', '100990000000');

-- 595. Big Countries

SELECT name, population, area
FROM World
WHERE area >= 3000000 OR population >= 25000000

-- Input
-- World =
-- | name        | continent | area    | population | gdp          |
-- | ----------- | --------- | ------- | ---------- | ------------ |
-- | Afghanistan | Asia      | 652230  | 25500100   | 20343000000  |
-- | Albania     | Europe    | 28748   | 2831741    | 12960000000  |
-- | Algeria     | Africa    | 2381741 | 37100000   | 188681000000 |
-- | Andorra     | Europe    | 468     | 78115      | 3712000000   |
-- | Angola      | Africa    | 1246700 | 20609294   | 100990000000 |
-- Output
-- | name        | population | area    |
-- | ----------- | ---------- | ------- |
-- | Afghanistan | 25500100   | 652230  |
-- | Algeria     | 37100000   | 2381741 |
