package sql

const CreateSchema = `
	CREATE TABLE IF NOT EXISTS schema_migrations (
		version integer
	);
`

const SelectSchema = `
	SELECT * FROM schema_migrations;
`

const InsertSchema = `
	INSERT INTO schema_migrations (version) VALUES (0);
`
