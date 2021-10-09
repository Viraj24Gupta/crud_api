package server

import (
	"net/http"
	"time"
)

func StartAPIServer(port string) error {
	srv := &http.Server{
		Handler:      Router(),
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}
