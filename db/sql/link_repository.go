package sql

const GetLinkByID = `
	SELECT id, short_uri, url, created_at
	FROM links
	WHERE id = $1
`
