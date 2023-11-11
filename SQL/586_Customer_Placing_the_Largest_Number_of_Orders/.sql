CREATE TABLE IF NOT EXISTS orders (order_number INT, customer_number INT);
TRUNCATE TABLE orders;
INSERT INTO orders (order_number, customer_number) VALUES ('1', '1');
INSERT INTO orders (order_number, customer_number) VALUES ('2', '2');
INSERT INTO orders (order_number, customer_number) VALUES ('3', '3');
INSERT INTO orders (order_number, customer_number) VALUES ('4', '3');

-- 586. Customer Placing the Largest Number of Orders

SELECT customer_number
FROM Orders
GROUP BY customer_number
ORDER BY COUNT(customer_number) DESC
LIMIT 1


-- Input
-- orders =
-- | order_number | customer_number |
-- | ------------ | --------------- |
-- | 1            | 1               |
-- | 2            | 2               |
-- | 3            | 3               |
-- | 4            | 3               |
-- Output
-- | customer_number |
-- | --------------- |
-- | 3               |
