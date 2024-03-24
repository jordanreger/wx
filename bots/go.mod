module github.com/jordanreger/wx/bots

go 1.22.0

//replace github.com/jordanreger/wx/bots/social => ./social

//replace github.com/jordanreger/wx/products/warnings => ../products/warnings

require (
	github.com/joho/godotenv v1.5.1
	github.com/jordanreger/wx/bots/social v0.0.0-20240324140212-0f04ca07067f
	github.com/jordanreger/wx/products/warnings v0.0.0-20240324140212-0f04ca07067f
)

require (
	github.com/jordanreger/htmlsky/bsky v0.0.0-20240306155219-da790a5855ce // indirect
	github.com/jordanreger/htmlsky/util v0.0.0-20240306155219-da790a5855ce // indirect
)
