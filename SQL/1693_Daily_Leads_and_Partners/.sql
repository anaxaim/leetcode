CREATE TABLE IF NOT EXISTS DailySales(date_id DATE, make_name VARCHAR(20), lead_id INT, partner_id INT);
TRUNCATE TABLE DailySales;
INSERT INTO DailySales (date_id, make_name, lead_id, partner_id) VALUES ('2020-12-8', 'toyota', '0', '1');
INSERT INTO DailySales (date_id, make_name, lead_id, partner_id) VALUES ('2020-12-8', 'toyota', '1', '0');
INSERT INTO DailySales (date_id, make_name, lead_id, partner_id) VALUES ('2020-12-8', 'toyota', '1', '2');
INSERT INTO DailySales (date_id, make_name, lead_id, partner_id) VALUES ('2020-12-7', 'toyota', '0', '2');
INSERT INTO DailySales (date_id, make_name, lead_id, partner_id) VALUES ('2020-12-7', 'toyota', '0', '1');
INSERT INTO DailySales (date_id, make_name, lead_id, partner_id) VALUES ('2020-12-8', 'honda', '1', '2');
INSERT INTO DailySales (date_id, make_name, lead_id, partner_id) VALUES ('2020-12-8', 'honda', '2', '1');
INSERT INTO DailySales (date_id, make_name, lead_id, partner_id) VALUES ('2020-12-7', 'honda', '0', '1');
INSERT INTO DailySales (date_id, make_name, lead_id, partner_id) VALUES ('2020-12-7', 'honda', '1', '2');
INSERT INTO DailySales (date_id, make_name, lead_id, partner_id) VALUES ('2020-12-7', 'honda', '2', '1');

-- 1693. Daily Leads and Partners

SELECT date_id, make_name, COUNT(DISTINCT lead_id) AS unique_leads, COUNT(DISTINCT partner_id) AS unique_partners
FROM DailySales
GROUP BY date_id, make_name

-- Input
-- DailySales =
-- | date_id   | make_name | lead_id | partner_id |
-- | --------- | --------- | ------- | ---------- |
-- | 2020-12-8 | toyota    | 0       | 1          |
-- | 2020-12-8 | toyota    | 1       | 0          |
-- | 2020-12-8 | toyota    | 1       | 2          |
-- | 2020-12-7 | toyota    | 0       | 2          |
-- | 2020-12-7 | toyota    | 0       | 1          |
-- | 2020-12-8 | honda     | 1       | 2          |...
-- Output
-- | date_id    | make_name | unique_leads | unique_partners |
-- | ---------- | --------- | ------------ | --------------- |
-- | 2020-12-07 | honda     | 3            | 2               |
-- | 2020-12-07 | toyota    | 1            | 2               |
-- | 2020-12-08 | honda     | 2            | 2               |
-- | 2020-12-08 | toyota    | 2            | 3               |
