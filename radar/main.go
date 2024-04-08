package main

import (
	"bytes"
	"compress/gzip"
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res, err := http.Get("https://mrms.ncep.noaa.gov/data/RIDGEII/L3/KSHV/SR_BREF/KSHV_L3_SR_BREF_20240408_230425.tif.gz")
		if err != nil {
			panic(err)
		}

		data, err := io.ReadAll(res.Body)
		b := bytes.NewBuffer([]byte(data))
		if err != nil {
			panic(err)
		}
		gz, err := gzip.NewReader(b)
		if err != nil {
			panic(err)
		}

		page, _ := io.ReadAll(gz)

		w.Header().Add("Content-type", "image/tiff")
		w.Write(page)
	})

	log.Fatal(http.ListenAndServe(":80", mux))
}
