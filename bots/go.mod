module github.com/fjalldev/wx/bots

go 1.22.0

replace github.com/fjalldev/wx/products/warnings => ../products/warnings
replace github.com/fjalldev/wx/social => ../social

require (
	github.com/fjalldev/wx/products/warnings v0.0.1
	github.com/fjalldev/wx/social v0.0.0-20240221182737-f25f3710c363
)

require github.com/joho/godotenv v1.5.1
