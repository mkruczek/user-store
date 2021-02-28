ALTER TABLE users ADD COLUMN update_date BIGINT;

WITH icd AS (SELECT id, create_date FROM users WHERE update_date is null)
UPDATE users SET update_date= icd.create_date FROM icd WHERE users.id = icd.id;

ALTER TABLE users ALTER COLUMN update_date SET NOT NULL;