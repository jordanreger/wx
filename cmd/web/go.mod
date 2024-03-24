module github.com/jordanreger/wx/cmd/web

go 1.22.0

replace github.com/jordanreger/wx/bots/social => ../../bots/social

replace github.com/jordanreger/wx/products/warnings => ../../products/warnings

require (
	github.com/jordanreger/wx/bots/social v0.0.0-00010101000000-000000000000
	github.com/jordanreger/wx/products/warnings v0.0.0-00010101000000-000000000000
)

require (
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/jordanreger/htmlsky/bsky v0.0.0-20240306155219-da790a5855ce // indirect
	github.com/jordanreger/htmlsky/util v0.0.0-20240306155219-da790a5855ce // indirect
)
