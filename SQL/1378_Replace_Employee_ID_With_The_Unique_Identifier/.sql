CREATE TABLE IF NOT EXISTS Employees (id INT, name VARCHAR(20));
CREATE TABLE IF NOT EXISTS EmployeeUNI (id INT, unique_id INT);
TRUNCATE TABLE Employees;
INSERT INTO Employees (id, name) VALUES ('1', 'Alice');
INSERT INTO Employees (id, name) VALUES ('7', 'Bob');
INSERT INTO Employees (id, name) VALUES ('11', 'Meir');
INSERT INTO Employees (id, name) VALUES ('90', 'Winston');
INSERT INTO Employees (id, name) VALUES ('3', 'Jonathan');
TRUNCATE TABLE EmployeeUNI;
INSERT INTO EmployeeUNI (id, unique_id) VALUES ('3', '1');
INSERT INTO EmployeeUNI (id, unique_id) VALUES ('11', '2');
INSERT INTO EmployeeUNI (id, unique_id) VALUES ('90', '3');

-- 1378. Replace Employee ID With The Unique Identifier

SELECT u.unique_id, e.name
FROM Employees e
LEFT OUTER JOIN EmployeeUNI u ON e.id = u.id

-- Input
-- Employees =
-- | id | name     |
-- | -- | -------- |
-- | 1  | Alice    |
-- | 7  | Bob      |
-- | 11 | Meir     |
-- | 90 | Winston  |
-- | 3  | Jonathan |
-- EmployeeUNI =
-- | id | unique_id |
-- | -- | --------- |
-- | 3  | 1         |
-- | 11 | 2         |
-- | 90 | 3         |
-- Output
-- | unique_id | name     |
-- | --------- | -------- |
-- | null      | Alice    |
-- | 1         | Jonathan |
-- | null      | Bob      |
-- | 2         | Meir     |
-- | 3         | Winston  |
