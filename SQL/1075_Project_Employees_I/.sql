CREATE TABLE IF NOT EXISTS Project (project_id INT, employee_id INT);
CREATE TABLE IF NOT EXISTS Employee (employee_id INT, name VARCHAR(10), experience_years INT);
TRUNCATE TABLE Project;
INSERT INTO Project (project_id, employee_id) VALUES ('1', '1');
INSERT INTO Project (project_id, employee_id) VALUES ('1', '2');
INSERT INTO Project (project_id, employee_id) VALUES ('1', '3');
INSERT INTO Project (project_id, employee_id) VALUES ('2', '1');
INSERT INTO Project (project_id, employee_id) VALUES ('2', '4');
TRUNCATE TABLE Employee;
INSERT INTO Employee (employee_id, name, experience_years) VALUES ('1', 'Khaled', '3');
INSERT INTO Employee (employee_id, name, experience_years) VALUES ('2', 'Ali', '2');
INSERT INTO Employee (employee_id, name, experience_years) VALUES ('3', 'John', '1');
INSERT INTO Employee (employee_id, name, experience_years) VALUES ('4', 'Doe', '2');

-- 1075. Project Employees I

SELECT p.project_id, ROUND(AVG(e.experience_years), 2) AS average_years
FROM Project p
         JOIN Employee e ON p.employee_id = e.employee_id
GROUP BY p.project_id

-- Input
-- Project =
-- | project_id | employee_id |
-- | ---------- | ----------- |
-- | 1          | 1           |
-- | 1          | 2           |
-- | 1          | 3           |
-- | 2          | 1           |
-- | 2          | 4           |
-- Employee =
-- | employee_id | name   | experience_years |
-- | ----------- | ------ | ---------------- |
-- | 1           | Khaled | 3                |
-- | 2           | Ali    | 2                |
-- | 3           | John   | 1                |
-- | 4           | Doe    | 2                |
-- Output
-- | project_id | average_years |
-- | ---------- | ------------- |
-- | 2          | 2.5           |
-- | 1          | 2             |
