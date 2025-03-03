-- +goose Up
CREATE TABLE IF NOT EXISTS scheduler (
		-- id INTEGER PRIMARY KEY AUTOINCREMENT,
		id SERIAL PRIMARY KEY,
		date CHAR(8) NOT NULL DEFAULT '',
		title VARCHAR(256) NOT NULL DEFAULT '',
		comment TEXT NOT NULL DEFAULT '',
		repeat VARCHAR(128) NOT NULL DEFAULT '');
		
CREATE INDEX IF NOT EXISTS date_scheduler ON scheduler (date);

-- +goose Down
DROP TABLE scheduler;