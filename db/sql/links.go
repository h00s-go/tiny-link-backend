package sql

const GetLinkByID = `
	SELECT id, url, created_by, created_at
	FROM links
	WHERE id = $1
`

const GetLinkByURL = `
	SELECT id, url, created_by, created_at
	FROM links
	WHERE url = $1
`

const CreateLink = `
	INSERT INTO links (
		url, created_by, created_at
	)
	VALUES (
		$1, $2, NOW()
	)
	RETURNING id
`
