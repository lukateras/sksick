package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
)

var buf = bytes.NewBuffer(nil)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write(buf.Bytes())
}

func main() {
	io.Copy(buf, os.Stdin)
	http.HandleFunc("/pks/lookup", handler)
	log.Fatal(http.ListenAndServe(":11371", nil))
}
