package repository

func (r *Repository) CreateScheduler() error {
	// query := `CREATE TABLE IF NOT EXISTS scheduler (
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	date CHAR(8) NOT NULL DEFAULT '',
	// 	title VARCHAR(256) NOT NULL DEFAULT '',
	// 	comment TEXT NOT NULL DEFAULT '',
	// 	repeat VARCHAR(128) NOT NULL DEFAULT '');

	// 	CREATE INDEX IF NOT EXISTS date_scheduler ON scheduler (date);`

	query := `CREATE TABLE IF NOT EXISTS scheduler (
			id SERIAL PRIMARY KEY,
			date CHAR(8) NOT NULL DEFAULT '',
			title VARCHAR(256) NOT NULL DEFAULT '',
			comment TEXT NOT NULL DEFAULT '',
			repeat VARCHAR(128) NOT NULL DEFAULT '');
			
			CREATE INDEX IF NOT EXISTS date_scheduler ON scheduler (date);`

	if _, err := r.db.Exec(query); err != nil {
		return err
	}
	return nil
}
