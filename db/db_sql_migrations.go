package db

const sqlCreateSchema = `
	CREATE TABLE IF NOT EXISTS schema_migrations (
		version integer
	);
`

const sqlSelectSchema = `
	SELECT * FROM schema_migrations;
`

const sqlInsertSchema = `
	INSERT INTO schema_migrations (version) VALUES (0);
`
