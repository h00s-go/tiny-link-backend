package sql

const GetLinkByID = `
	SELECT id, short_uri, url, created_at
	FROM links
	WHERE id = $1
`

const GetLinkByShortURI = `
	SELECT id, short_uri, url, created_at
	FROM links
	WHERE short_uri = $1
`
