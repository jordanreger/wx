package social

import (
	"bytes"
	"regexp"
	"strings"

	"github.com/jordanreger/htmlsky/util"
)

var handleRegex = regexp.MustCompile(`[$|\W](@([a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)`)

func ParseMentions(raw string) []*Facet {
	var mentions []*Facet

	rawBytes := []byte(raw)

	for _, m := range handleRegex.FindAllString(raw, -1) {
		did := util.GetDID(strings.Split(m, "@")[1])

		mentions = append(mentions,
			&Facet{
				Index: FacetIndex{
					ByteStart: bytes.Index(rawBytes, []byte(m)),
					ByteEnd:   bytes.Index(rawBytes, []byte(m)) + len(m),
				},
				Features: []*FacetFeature{
					{
						Type: "app.bsky.richtext.facet#mention",
						DID:  did,
					},
				},
			},
		)
	}

	return mentions
}

var urlRegex = regexp.MustCompile(`[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)
var emailRegex = regexp.MustCompile(`.*@.*`)

func ParseURLs(raw string) []*Facet {
	raw = strings.ReplaceAll(raw, "https://", "")
	raw = strings.ReplaceAll(raw, "http://", "")

	var urls []*Facet

	rawBytes := []byte(raw)

	for _, u := range urlRegex.FindAllString(raw, -1) {
		if !emailRegex.MatchString(u) {
			urls = append(urls,
				&Facet{
					Index: FacetIndex{
						ByteStart: bytes.Index(rawBytes, []byte(u)),
						ByteEnd:   bytes.Index(rawBytes, []byte(u)) + len(u),
					},
					Features: []*FacetFeature{
						{
							Type: "app.bsky.richtext.facet#link",
							URI:  "https://" + string(u),
						},
					},
				},
			)
		}
	}
	return urls
}

func ParseFacets(text string) []*Facet {
	var facets []*Facet

	facets = append(facets, ParseURLs(text)...)
	facets = append(facets, ParseMentions(text)...)

	return facets
}
