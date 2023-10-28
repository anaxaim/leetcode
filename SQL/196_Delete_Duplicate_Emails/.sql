CREATE TABLE IF NOT EXISTS Person (Id INT, Email VARCHAR(255));
TRUNCATE TABLE Person;
INSERT INTO Person (id, email) VALUES ('1', 'john@example.com');
INSERT INTO Person (id, email) VALUES ('2', 'bob@example.com');
INSERT INTO Person (id, email) VALUES ('3', 'john@example.com');

-- 183. Customers Who Never Order

DELETE FROM Person p1
USING Person p2
WHERE p1.Email = p2.Email AND p1.id > p2.id

-- Input
-- Person =
-- | id | email            |
-- | -- | ---------------- |
-- | 1  | john@example.com |
-- | 2  | bob@example.com  |
-- | 3  | john@example.com |
-- Output
-- | id | email            |
-- | -- | ---------------- |
-- | 1  | john@example.com |
-- | 2  | bob@example.com  |
