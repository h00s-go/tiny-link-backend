package migrations

const CreateLinks = `
	BEGIN;

	CREATE TABLE IF NOT EXISTS links (
		id bigserial PRIMARY KEY,
		short_uri text UNIQUE,
		url text UNIQUE NOT NULL,
		created_at timestamp NOT NULL
		last_accessed_at timestamp NULL,
		access_count bigint NOT NULL DEFAULT 0
	);
	
	UPDATE schema_migrations SET version = 1 WHERE version = 0;
	
	COMMIT;
`
