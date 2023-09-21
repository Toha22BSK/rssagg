-- migrate:up
ALTER TABLE feeds ADD COLUMN last_fetched_at TIMESTAMP;

-- migrate:down
ALTER TABLE feeds DROP COLUMN last_fetched_at;
