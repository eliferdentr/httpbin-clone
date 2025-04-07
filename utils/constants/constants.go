package constants

const (
	BasePath                = "https://eliferden.com"
	DOMAIN                  = "localhost"
	NONCE_BYTE_LENGTH       = 32
	NONCE_HASHING_ALGORITHM = "SHA256"
	EXPECTED_USERNAME       = "expected_username"
	EXPECTED_PASSWORD       = "expected_password"
	REALM                   = "Access to the site"
	EXPECTED_NONCE          = "1234567890abcdef"  // Demo amaçlı sabit nonce
	HTMLCONTENT				= `
		<!DOCTYPE html>
		<html lang="en">
			<head>
				<title>Sample HTML</title>
				<meta charset="UTF-8">
			</head>
			<body>
				<h1>Hello, this is a sample HTML page!</h1>
				<p>This is some example HTML content for the /html endpoint.</p>
			</body>
		</html>
		`
	ROBOTSCONTENT = `User-agent: *
		Disallow: /private/
		Disallow: /admin/
		`
)
