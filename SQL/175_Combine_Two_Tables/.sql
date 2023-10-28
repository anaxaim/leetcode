CREATE TABLE If NOT EXISTS Person (personId INT, firstName VARCHAR(255), lastName VARCHAR(255));
CREATE TABLE If NOT EXISTS Address (addressId INT, personId INT, city VARCHAR(255), state VARCHAR(255));
TRUNCATE TABLE Person;
INSERT INTO Person (personId, lastName, firstName) VALUES ('1', 'Wang', 'Allen');
INSERT INTO Person (personId, lastName, firstName) VALUES ('2', 'Alice', 'Bob');
TRUNCATE TABLE Address;
INSERT INTO Address (addressId, personId, city, state) VALUES ('1', '2', 'New York City', 'New York');
INSERT INTO Address (addressId, personId, city, state) VALUES ('2', '3', 'Leetcode', 'California');

-- 175. Combine Two Tables

SELECT Person.firstName, Person.lastName, Address.city, Address.state
FROM Person
LEFT JOIN Address ON Person.personId = Address.personId


-- Input
-- Person =
-- | personId | lastName | firstName |
-- | -------- | -------- | --------- |
-- | 1        | Wang     | Allen     |
-- | 2        | Alice    | Bob       |
--
-- Address =
-- | addressId | personId | city          | state      |
-- | --------- | -------- | ------------- | ---------- |
-- | 1         | 2        | New York City | New York   |
-- | 2         | 3        | Leetcode      | California |
--
-- Output
-- | firstName | lastName | city          | state    |
-- | --------- | -------- | ------------- | -------- |
-- | Allen     | Wang     | null          | null     |
-- | Bob       | Alice    | New York City | New York |

