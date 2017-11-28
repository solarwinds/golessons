package web

import "net/http"

// GetHello returns an http.Handler, which is a function with a special signature. Any
// function with this signature can be *cast* to an http.Handler and therefore used by
// the http package's default multiplexer to handle requests.
func GetHello(verbose bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		if verbose {
			w.Write([]byte("Here's some awesome body content!\n"))
			// good to put these here to ensure that this is the end of execution
			return
		}else{
			w.WriteHeader(http.StatusAccepted)
		}
	})
}