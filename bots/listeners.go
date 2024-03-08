package main

import (
	"time"

	"github.com/fjalldev/wx/bots/social"
	"github.com/fjalldev/wx/products/warnings"
)

func t_l(warning chan string) {
	for {
		latest := warnings.Latest("Tornado warning")
		post := social.GetPost(latest)
		warning <- post
		time.Sleep(15 * time.Second)
	}
}

func tstm_l(warning chan string) {
	for {
		latest := warnings.Latest("Severe thunderstorm warning")
		post := social.GetPost(latest)
		warning <- post
		time.Sleep(15 * time.Second)
	}
}

func ff_l(warning chan string) {
	for {
		latest := warnings.Latest("Flash flood warning")
		post := social.GetPost(latest)
		warning <- post
		time.Sleep(15 * time.Second)
	}
}
