// https://fjall.dev/wx/bots/social/bluesky.go

package social

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/jordanreger/htmlsky/bsky"
)

type BskyUser struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type PostRecord struct {
	Repo       string `json:"repo"`
	Collection string `json:"collection"`
	Record     Record `json:"record"`
}

type Record struct {
	Type      string       `json:"$type"`
	Text      string       `json:"text,omitempty"`
	CreatedAt time.Time    `json:"createdAt,omitempty"`
	Facets    []bsky.Facet `json:"facets,omitempty"`
	Langs     []string     `json:"langs,omitempty"`
}

type SessionBody struct {
	AccessJwt  string `json:"accessJwt"`
	RefreshJwt string `json:"refreshJwt"`
	Handle     string `json:"handle"`
	DID        string `json:"did"`
}

func PostToBluesky(identifier string, warning_text string) {
	//err := godotenv.Load(".env")
	//if err != nil {
	//	panic(err)
	//}

	client := &http.Client{}

	auth := BskyUser{Identifier: identifier, Password: os.Getenv("bsky_" + strings.ReplaceAll(identifier, ".", ""))}
	authJson, _ := json.Marshal(auth)
	getSession, err := http.NewRequest("POST", "https://bsky.social/xrpc/com.atproto.server.createSession", bytes.NewBuffer(authJson))
	if err != nil {
		fmt.Println(err)
	}
	getSession.Header.Add("Content-Type", "application/json")
	getSession.Header.Add("Accept", "application/json")
	getSession.Header.Add("User-Agent", identifier)
	res, _ := client.Do(getSession)
	defer getSession.Body.Close()

	var session SessionBody
	b, _ := io.ReadAll(res.Body)
	json.Unmarshal(b, &session)

	accessJwt := session.AccessJwt

	facets := bsky.ParseFacets(warning_text)

	now := time.Now()
	body := Record{Type: "app.bsky.feed.post", Text: warning_text, CreatedAt: now, Facets: facets, Langs: []string{"en-US"}}
	post := PostRecord{Repo: identifier, Collection: "app.bsky.feed.post", Record: body}
	postJson, _ := json.Marshal(post)

	sendPost, _ := http.NewRequest("POST", "https://bsky.social/xrpc/com.atproto.repo.createRecord?repo="+identifier, bytes.NewBuffer(postJson))
	sendPost.Header.Add("Authorization", "Bearer "+accessJwt)
	sendPost.Header.Add("Content-Type", "application/json")
	sendPost.Header.Add("User-Agent", identifier)
	client.Do(sendPost)
}
