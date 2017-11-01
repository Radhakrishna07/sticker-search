package main

import (
	"net/http"

	"github.com/Radhakrishna07/sticker-search/network"
	"github.com/Radhakrishna07/sticker-search/search"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {

	// config.SetupConfig()

	router := mux.NewRouter().StrictSlash(true)

	router.Handle("/stickers",
		network.RecoverWrap(http.HandlerFunc(search.GetStickers)))

	router.Handle("/stickers1",
		network.RecoverWrap(http.HandlerFunc(search.GetStickers))).Queries("q", "")

	negroniMiddleware := negroni.New()
	negroniMiddleware.UseHandler(router)
	negroniMiddleware.Run(":8080")

	// log.Fatal(http.ListenAndServe(":8080", router))

}
