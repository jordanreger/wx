module github.com/fjalldev/wx/social

go 1.22.0

require (
	github.com/joho/godotenv v1.5.1
	github.com/jordanreger/htmlsky/bsky v0.0.0-00010101000000-000000000000
)

require github.com/jordanreger/htmlsky/util v0.0.0-20240306053622-b168496a8417 // indirect

replace github.com/jordanreger/htmlsky/bsky => ../../htmlsky/bsky
