package main

import (
	// "fmt"
	"os"
	"path/filepath"
)

var path string

var t_chan = make(chan string)
var tstm_chan = make(chan string)
var ff_chan = make(chan string)

// var test_chan = make(chan string)

func main() {
	ex, _ := os.Executable()
	path = filepath.Dir(ex)
	os.MkdirAll("./warnings", 0764)

	go t_l(t_chan)
	go tstm_l(tstm_chan)
	go ff_l(ff_chan)
	// go test_l(test_chan)

	router()
}
