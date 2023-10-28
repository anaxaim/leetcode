CREATE TABLE IF NOT EXISTS Salary (id INT, name VARCHAR(100), sex CHAR(1), salary INT);
TRUNCATE TABLE Salary;
INSERT INTO Salary (id, name, sex, salary) VALUES ('1', 'A', 'm', '2500');
INSERT INTO Salary (id, name, sex, salary) VALUES ('2', 'B', 'f', '1500');
INSERT INTO Salary (id, name, sex, salary) VALUES ('3', 'C', 'm', '5500');
INSERT INTO Salary (id, name, sex, salary) VALUES ('4', 'D', 'f', '500');

-- 627. Swap Salary

UPDATE Salary
SET sex =
CASE sex
    WHEN 'm' THEN 'f'
    ELSE 'm'
END;


-- Input
-- Salary =
-- | id | name | sex | salary |
-- | -- | ---- | --- | ------ |
-- | 1  | A    | m   | 2500   |
-- | 2  | B    | f   | 1500   |
-- | 3  | C    | m   | 5500   |
-- | 4  | D    | f   | 500    |
-- Output
-- | id | name | sex | salary |
-- | -- | ---- | --- | ------ |
-- | 1  | A    | f   | 2500   |
-- | 2  | B    | m   | 1500   |
-- | 3  | C    | f   | 5500   |
-- | 4  | D    | m   | 500    |
