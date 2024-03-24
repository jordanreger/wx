package social_test

import (
	"testing"

	"github.com/jordanreger/wx/bots/social"
	"github.com/jordanreger/wx/products/warnings"
)

func TestPost_Real(t *testing.T) {
	latest := warnings.Latest("Tornado warning")
	post := social.GetPost(latest)
	t.Log(post)
}

func TestPost_Fake(t *testing.T) {
	latest := warnings.Latest("Test")
	post := social.GetPost(latest)
	t.Log(post)
}
