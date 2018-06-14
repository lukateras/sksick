package main

import (
	"bytes"
	"io"
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
	http.ListenAndServe(":11371", nil)
}
