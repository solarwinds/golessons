package web

import "net/http"

// GetHello returns an http.Handler, which is a function with a special signature. Any
// function with this signature can be *cast* to an http.Handler and therefore used by
// the http package's default multiplexer to handle requests.
func GetHello() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// uncomment the below to see the status header change from 200 to 202
		//w.WriteHeader(http.StatusAccepted)

		w.Write([]byte("awesome!"))
	})
}
