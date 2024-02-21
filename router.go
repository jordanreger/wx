package main

import (
	"os"
	"strings"

	"github.com/fjalldev/social"
)

func last_warning(warning_type string) string {
	fp := path + "/warnings/" + warning_type + "_warning"
	m := os.FileMode(0644)

	stored_warning, err := os.ReadFile(fp)
	if err != nil {
		warning_type_friendly := strings.Replace(warning_type, "_", " ", -1)
		os.WriteFile(fp, []byte("No current "+warning_type_friendly+" warnings"), m)
	}

	return string(stored_warning)
}

func update_last_warning(warning_type string, warning_text string) {
	fp := path + "/warnings/" + warning_type + "_warning"
	m := os.FileMode(0644)

	os.WriteFile(fp, []byte(warning_text), m)
}

func router() {
	for {
		select {
		// tornado warning
		case w := <-t_chan:
			lw := last_warning("tornado")
			if w != lw {
				update_last_warning("tornado", w)
				if w != "No current tornado warnings" {
					social.PostToBluesky("nwstornado.bsky.social", w)
				}
			}

		// severe thunderstorm warning
		case w := <-tstm_chan:
			lw := last_warning("severe_thunderstorm")
			if w != lw {
				update_last_warning("severe_thunderstorm", w)
				if w != "No current severe thunderstorm warnings" {
					social.PostToBluesky("nwsseveretstorm.bsky.social", w)
				}
			}

		// flash flood warning
		case w := <-ff_chan:
			lw := last_warning("flash_flood")
			if w != lw {
				update_last_warning("flash_flood", w)
				if w != "No current flash flood warnings" {
					social.PostToBluesky("nwsflashflood.bsky.social", w)
				}
			}

			// test warning
			// case w := <-test_chan:
			// 	lw := last_warning("test")
			// 	if w != lw {
			// 		update_last_warning("test", w)
			// 		if w != "No current small craft advisorys" {
			// 			social.PostToBluesky("fjall.net", w)
			// 		}
			// 	}
		}
	}
}
