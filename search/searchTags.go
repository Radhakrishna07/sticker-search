package search

import (
	"fmt"
	"net/http"
)

func GetStickers(w http.ResponseWriter, r *http.Request) {
	orderIdString := r.URL.Query().Get("q") // Returns a url.Values, which is a map[string][]string
	if orderIdString != "" {
		fmt.Fprintln(w, "Hello, you have entered a query : "+orderIdString)
	} else {
		fmt.Fprintln(w, "Hello, you haven't entered a query!")
	}

}
