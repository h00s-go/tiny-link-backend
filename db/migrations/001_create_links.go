package migrations

const CreateLinks = `
	BEGIN;

	CREATE TABLE IF NOT EXISTS links (
		id bigserial PRIMARY KEY,
		url text UNIQUE NOT NULL,
		created_at timestamp NOT NULL
	);
	
	UPDATE schema_migrations SET version = 1 WHERE version = 0;
	
	COMMIT;
`
