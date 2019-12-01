package server

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/koheyeng/go_clean_architecture_sample/cmd/server/core"
)

func Serve(c *core.Core) error {
	if err := c.Valid(); err != nil {
		return err
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
		InsecureSkipVerify: c.APIInsecureTLS,
	}
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 8

	//
	// Set handlers
	//

	rt := mux.NewRouter().StrictSlash(true)

	// Start serving
	log.Println("Sample API Server starting...")
	http.ListenAndServe(c.BindAddress, rt)

	return nil
}
