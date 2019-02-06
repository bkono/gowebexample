package handler

import (
	"fmt"
	"net/http"
)

func ImageUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from upload, the url is %q\n", r.URL)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from index, the url is %q\n", r.URL)
}
