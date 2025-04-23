package constants

const (
	BasePath                = "https://eliferden.com"
	DOMAIN                  = "localhost"
	NONCE_BYTE_LENGTH       = 32
	NONCE_HASHING_ALGORITHM = "SHA256"
	EXPECTED_USERNAME       = "expected_username"
	EXPECTED_PASSWORD       = "expected_password"
	REALM                   = "Access to the site"
	EXPECTED_BEARER_TOKEN 	= "my-super-secret-token-123"
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
	ROBOTSCONTENT 			= `User-agent: *
							Disallow: /private/
							Disallow: /admin/
							`

	FORMCONTENT				= `
							<!DOCTYPE html>
							<html lang="en">
							<head>
							<meta charset="UTF-8">
							<meta name="viewport" content="width=device-width, initial-scale=1.0">
							<title>POST Form Example</title>
							</head>
							<body>
							<h1>Submit Your Data</h1>
							<form action="/post" method="POST">
								<label for="message">Message:</label><br>
								<input type="text" id="message" name="message" placeholder="Enter your message"><br><br>
								<input type="submit" value="Submit">
							</form>
							</body>
							</html>
							`
)
