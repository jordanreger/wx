module github.com/fjalldev/wx/bots

go 1.22.0

replace github.com/fjalldev/wx/products/warnings => ../products/warnings

replace github.com/fjalldev/wx/social => ../social

require github.com/fjalldev/wx/products/warnings v0.0.0-20240306015657-5f82b4ae67b6

require (
	github.com/fjalldev/wx/social v0.0.0-20240306015657-5f82b4ae67b6
	github.com/joho/godotenv v1.5.1
	github.com/jordanreger/htmlsky/util v0.0.0-20240306013103-2321888b9137 // indirect
)
