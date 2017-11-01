package network

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func RecoverWrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			r := recover()
			if r != nil {
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknown error")
				}
				log.Println("slr-recover", fmt.Sprint("recoverd from:", err.Error(), ", stacktrace:", string(debug.Stack())))
				http.Error(w, "Sorry.Something went wrong internal to our server", http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	})
}
