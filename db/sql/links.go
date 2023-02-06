package sql

const GetLinkByID = `
	SELECT id, short_uri, url, created_at, last_accessed_at, access_count
	FROM links
	WHERE id = $1
`

const GetLinkByShortURI = `
	SELECT id, short_uri, url, created_at, last_accessed_at, access_count
	FROM links
	WHERE short_uri = $1
`

const GetLinkByURL = `
	SELECT id, short_uri, url, created_at, last_accessed_at, access_count
	FROM links
	WHERE url = $1
`

const CreateLink = `
	INSERT INTO links (
		url, created_at
	)
	VALUES (
		$1, NOW()
	)
	RETURNING id
`

const UpdateLinkShortURI = `
	UPDATE links
	SET short_uri = $1
	WHERE id = $2
`
