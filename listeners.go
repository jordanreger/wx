package main

import (
	"time"

	"github.com/fjallwx/products/warnings"
)

func t_l(warning chan string) {
	for {
		latest := warnings.Latest("Tornado warning")
		warning <- latest
		time.Sleep(15 * time.Second)
	}
}

func tstm_l(warning chan string) {
	for {
		latest := warnings.Latest("Severe thunderstorm warning")
		warning <- latest
		time.Sleep(15 * time.Second)
	}
}

func ff_l(warning chan string) {
	for {
		latest := warnings.Latest("Flash flood warning")
		warning <- latest
		time.Sleep(15 * time.Second)
	}
}

// func test_l(warning chan string) {
// 	for {
// 		latest := warnings.Latest("Small craft advisory")
// 		warning <- latest
// 		time.Sleep(15 * time.Second)
// 	}
// }
