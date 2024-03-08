package main

import (
	"bytes"
	"io"
	"net/http"
	"os"

	"github.com/fjalldev/wx/bots/social"
	"github.com/joho/godotenv"
)

func last_warning(warning_type string) string {
	res, _ := http.Get("https://kv.wx.fjall.dev/get/" + warning_type)

	w, _ := io.ReadAll(res.Body)
	return string(w)
}

func update_last_warning(warning_type string, warning_text string) {
	godotenv.Load(".env")

	client := &http.Client{}

	warning := []byte(warning_text)
	req, _ := http.NewRequest("POST", "https://kv.wx.fjall.dev/put", bytes.NewBuffer(warning))
	req.Header.Add("Authentication", os.Getenv("bsky_nws"+warning_type+"bskysocial"))
	req.Header.Add("User-Agent", "nws"+warning_type+".bsky.social")
	client.Do(req)
	defer req.Body.Close()
}

func router() {
	for {
		select {
		case w := <-t_chan:
			lw := last_warning("tornado")
			if w != lw {
				update_last_warning("tornado", w)
				if w != "No current tornado warnings" {
					social.PostToBluesky("nwstornado.bsky.social", w)
				}
			}
		case w := <-tstm_chan:
			lw := last_warning("severetstorm")
			if w != lw {
				update_last_warning("severetstorm", w)
				if w != "No current severe thunderstorm warnings" {
					social.PostToBluesky("nwsseveretstorm.bsky.social", w)
				}
			}
		case w := <-ff_chan:
			lw := last_warning("flashflood")
			if w != lw {
				update_last_warning("flashflood", w)
				if w != "No current flash flood warnings" {
					social.PostToBluesky("nwsflashflood.bsky.social", w)
				}
			}
		}
	}
}
