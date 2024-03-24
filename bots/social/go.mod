module github.com/jordanreger/wx/bots/social

go 1.22.0

require (
	github.com/joho/godotenv v1.5.1
	github.com/jordanreger/htmlsky/bsky v0.0.0-20240306155219-da790a5855ce
	github.com/jordanreger/wx/products/warnings v0.0.0-00010101000000-000000000000
)

require (
	github.com/fjalldev/wx/products/warnings v0.0.0-20240308015331-748edfb98f43 // indirect
	github.com/jordanreger/htmlsky/util v0.0.0-20240306155219-da790a5855ce // indirect
)

replace github.com/jordanreger/wx/products/warnings => ../../products/warnings
