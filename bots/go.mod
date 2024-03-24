module github.com/jordanreger/wx/bots

go 1.22.0

//replace github.com/jordanreger/wx/bots/social => ./social

//replace github.com/jordanreger/wx/products/warnings => ../products/warnings

require (
	github.com/joho/godotenv v1.5.1
)

require (
	github.com/jordanreger/htmlsky/bsky v0.0.0-20240306155219-da790a5855ce // indirect
	github.com/jordanreger/htmlsky/util v0.0.0-20240306155219-da790a5855ce // indirect
)
