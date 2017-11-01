package network

import (
	"net/http"
)

var req http.Request
var err error

// SetCAPIHeaders sets capi headers for the http request 'req'
func SetCAPIHeaders(req *http.Request) *http.Request {
	req.Header["Country_Name"] = []string{"IND"}
	req.Header["Channel_Name"] = []string{"MOBILE_APP"}
	req.Header["auth_key"] = []string{"487a342c-92f1-41ae-81fa-aaa5120f6bb3"}
	req.Header["BusinessUnit"] = []string{"BUS"}
	req.Header["Content-Type"] = []string{"application/json"}
	req.Header["os"] = []string{"Android"}
	req.Header["AppVersionCode"] = []string{"30026"}
	return req
}

// SetGOIHeaders sets capi headers for the http request 'req'
func SetGOIHeaders(req *http.Request) *http.Request {
	if err != nil {
		// Add GoIbIbo headers here

	}
	return req
}
