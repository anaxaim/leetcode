CREATE TABLE IF NOT EXISTS Person (id INT, email VARCHAR(255));
TRUNCATE TABLE Person;
INSERT INTO Person (id, email) VALUES ('1', 'a@b.com');
INSERT INTO Person (id, email) VALUES ('2', 'c@d.com');
INSERT INTO Person (id, email) VALUES ('3', 'a@b.com');

-- 182. Duplicate Emails

SELECT Email
FROM Person
GROUP BY email
HAVING COUNT(Email) > 1

-- Input
-- Person =
-- | id | email   |
-- | -- | ------- |
-- | 1  | a@b.com |
-- | 2  | c@d.com |
-- | 3  | a@b.com |
--
-- Output
-- | Email   |
-- | ------- |
-- | a@b.com |
