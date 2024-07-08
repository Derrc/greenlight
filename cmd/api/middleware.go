package main

import (
	"fmt"
	"net/http"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// create deferred function which always runs in the event of a panic
		// only applies to current goroutine
		defer func() {
			// recover() stops panicking sequence in event of a panic
			if err := recover(); err != nil {
				// close underlying HTTP connection
				w.Header().Set("Connection", "close")
				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
