package main

import (
	"net/http"
	"time"
)

func clearCookie(w *http.ResponseWriter, r *http.Request, name string) {
	if _, err := r.Cookie(name); err == nil {
		c := http.Cookie{
			Name:    name,
			Value:   "",
			Expires: time.Unix(0, 0),
		}
		http.SetCookie(*w, &c)
	}
}
