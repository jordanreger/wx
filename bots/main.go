package main

var t_chan = make(chan string)
var tstm_chan = make(chan string)
var ff_chan = make(chan string)

func main() {
	go t_l(t_chan)
	go tstm_l(tstm_chan)
	go ff_l(ff_chan)

	router()
}
