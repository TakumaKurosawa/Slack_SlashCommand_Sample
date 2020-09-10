package handler

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()

	fmt.Fprintf(w, "Hello!, %s", values.Get("text"))
}
