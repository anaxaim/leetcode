CREATE TABLE IF NOT EXISTS Employee (empId INT, name VARCHAR(255), supervisor INT, salary INT);
CREATE TABLE IF NOT EXISTS Bonus (empId INT, bonus INT);
TRUNCATE TABLE Employee;
INSERT INTO Employee (empId, name, supervisor, salary) VALUES ('3', 'Brad', 'None', '4000');
INSERT INTO Employee (empId, name, supervisor, salary) VALUES ('1', 'John', '3', '1000');
INSERT INTO Employee (empId, name, supervisor, salary) VALUES ('2', 'Dan', '3', '2000');
INSERT INTO Employee (empId, name, supervisor, salary) VALUES ('4', 'Thomas', '3', '4000');
TRUNCATE TABLE Bonus;
INSERT INTO Bonus (empId, bonus) values ('2', '500');
INSERT INTO Bonus (empId, bonus) values ('4', '2000');

-- 577. Employee Bonus

SELECT e.name, b.bonus
FROM Employee e
         LEFT JOIN Bonus b ON e.empId = b.empId
WHERE b.bonus < 1000 OR b.bonus IS NULL

-- Input
-- Employee =
-- | empId | name   | supervisor | salary |
-- | ----- | ------ | ---------- | ------ |
-- | 3     | Brad   | null       | 4000   |
-- | 1     | John   | 3          | 1000   |
-- | 2     | Dan    | 3          | 2000   |
-- | 4     | Thomas | 3          | 4000   |
-- Bonus =
-- | empId | bonus |
-- | ----- | ----- |
-- | 2     | 500   |
-- | 4     | 2000  |
-- Output
-- | name | bonus |
-- | ---- | ----- |
-- | Dan  | 500   |
-- | John | null  |
-- | Brad | null  |
