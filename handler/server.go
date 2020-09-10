package handler

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Hello!, %s", values.Get("text"))
}
