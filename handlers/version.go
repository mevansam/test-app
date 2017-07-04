package handlers

import "net/http"

type Version struct {
}

func (_ *Version) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("v0.0.2"))
}
