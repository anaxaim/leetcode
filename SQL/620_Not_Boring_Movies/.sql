CREATE TABLE IF NOT EXISTS cinema (id INT, movie VARCHAR(255), description VARCHAR(255), rating FLOAT(2));
TRUNCATE TABLE cinema;
INSERT INTO cinema (id, movie, description, rating) VALUES ('1', 'War', 'great 3D', '8.9');
INSERT INTO cinema (id, movie, description, rating) VALUES ('2', 'Science', 'fiction', '8.5');
INSERT INTO cinema (id, movie, description, rating) VALUES ('3', 'irish', 'boring', '6.2');
INSERT INTO cinema (id, movie, description, rating) VALUES ('4', 'Ice song', 'Fantacy', '8.6');
INSERT INTO cinema (id, movie, description, rating) VALUES ('5', 'House card', 'Interesting', '9.1');

-- 620. Not Boring Movies

SELECT * FROM Cinema c
WHERE id % 2 = 1 AND description != 'boring'
ORDER BY rating DESC

-- Input
-- cinema =
-- | id | movie      | description | rating |
-- | -- | ---------- | ----------- | ------ |
-- | 1  | War        | great 3D    | 8.9    |
-- | 2  | Science    | fiction     | 8.5    |
-- | 3  | irish      | boring      | 6.2    |
-- | 4  | Ice song   | Fantacy     | 8.6    |
-- | 5  | House card | Interesting | 9.1    |
-- Output
-- | id | movie      | description | rating |
-- | -- | ---------- | ----------- | ------ |
-- | 5  | House card | Interesting | 9.1    |
-- | 1  | War        | great 3D    | 8.9    |
