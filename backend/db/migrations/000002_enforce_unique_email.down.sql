BEGIN;

ALTER TABLE users
DROP CONSTRAINT IF EXISTS unique_email;

COMMIT;