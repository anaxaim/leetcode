CREATE TABLE IF NOT EXISTS SalesPerson (sales_id INT, name VARCHAR(255), salary INT, commission_rate INT, hire_date DATE);
CREATE TABLE IF NOT EXISTS Company (com_id INT, name VARCHAR(255), city VARCHAR(255));
CREATE TABLE IF NOT EXISTS Orders (order_id INT, order_date DATE, com_id INT, sales_id INT, amount INT);
TRUNCATE TABLE SalesPerson;
INSERT INTO SalesPerson (sales_id, name, salary, commission_rate, hire_date) VALUES ('1', 'John', '100000', '6', '4/1/2006');
INSERT INTO SalesPerson (sales_id, name, salary, commission_rate, hire_date) VALUES ('2', 'Amy', '12000', '5', '5/1/2010');
INSERT INTO SalesPerson (sales_id, name, salary, commission_rate, hire_date) VALUES ('3', 'Mark', '65000', '12', '12/25/2008');
INSERT INTO SalesPerson (sales_id, name, salary, commission_rate, hire_date) VALUES ('4', 'Pam', '25000', '25', '1/1/2005');
INSERT INTO SalesPerson (sales_id, name, salary, commission_rate, hire_date) VALUES ('5', 'Alex', '5000', '10', '2/3/2007');
TRUNCATE TABLE Company;
INSERT INTO Company (com_id, name, city) VALUES ('1', 'RED', 'Boston');
INSERT INTO Company (com_id, name, city) VALUES ('2', 'ORANGE', 'New York');
INSERT INTO Company (com_id, name, city) VALUES ('3', 'YELLOW', 'Boston');
INSERT INTO Company (com_id, name, city) VALUES ('4', 'GREEN', 'Austin');
TRUNCATE TABLE Orders;
INSERT INTO Orders (order_id, order_date, com_id, sales_id, amount) VALUES ('1', '1/1/2014', '3', '4', '10000');
INSERT INTO Orders (order_id, order_date, com_id, sales_id, amount) VALUES ('2', '2/1/2014', '4', '5', '5000');
INSERT INTO Orders (order_id, order_date, com_id, sales_id, amount) VALUES ('3', '3/1/2014', '1', '1', '50000');
INSERT INTO Orders (order_id, order_date, com_id, sales_id, amount) VALUES ('4', '4/1/2014', '1', '4', '25000');

-- 607. Sales Person

SELECT s.name
FROM SalesPerson s
WHERE s.sales_id NOT IN (
    SELECT o.sales_id
    FROM Orders o
             JOIN Company c ON o.com_id = c.com_id
    WHERE c.name = 'RED'
)

-- Input
-- SalesPerson =
-- | sales_id | name | salary | commission_rate | hire_date  |
-- | -------- | ---- | ------ | --------------- | ---------- |
-- | 1        | John | 100000 | 6               | 4/1/2006   |
-- | 2        | Amy  | 12000  | 5               | 5/1/2010   |
-- | 3        | Mark | 65000  | 12              | 12/25/2008 |
-- | 4        | Pam  | 25000  | 25              | 1/1/2005   |
-- | 5        | Alex | 5000   | 10              | 2/3/2007   |
-- Company =
-- | com_id | name   | city     |
-- | ------ | ------ | -------- |
-- | 1      | RED    | Boston   |
-- | 2      | ORANGE | New York |
-- | 3      | YELLOW | Boston   |
-- | 4      | GREEN  | Austin   |
-- Orders =
-- | order_id | order_date | com_id | sales_id | amount |
-- | -------- | ---------- | ------ | -------- | ------ |
-- | 1        | 1/1/2014   | 3      | 4        | 10000  |
-- | 2        | 2/1/2014   | 4      | 5        | 5000   |
-- | 3        | 3/1/2014   | 1      | 1        | 50000  |
-- | 4        | 4/1/2014   | 1      | 4        | 25000  |
-- Output
-- | name |
-- | ---- |
-- | Amy  |
-- | Mark |
-- | Alex |
