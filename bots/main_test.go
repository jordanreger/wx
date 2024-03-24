package main_test

import (
	"github.com/jordanreger/wx/bots/social"
	"testing"
)

func TestPost(t *testing.T) {
	wt := "This is a test. Please disregard."

	social.PostToBluesky("nwstornado.bsky.social", wt)
	social.PostToBluesky("nwsseveretstorm.bsky.social", wt)
	social.PostToBluesky("nwsflashflood.bsky.social", wt)
}
