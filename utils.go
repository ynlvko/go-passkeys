package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func logRequest(r *http.Request) {
	buf, _ := io.ReadAll(r.Body)

	rdr1 := io.NopCloser(bytes.NewBuffer(buf))
	rdr2 := io.NopCloser(bytes.NewBuffer(buf))
	log.Printf("BODY: %q", rdr1)
	r.Body = rdr2
}
