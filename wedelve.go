package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	db Database
)

func main() {
	var (
		success      bool
		port         string
		address      string
		mongoAddress string
		err          error
	)

	if port, success = os.LookupEnv("WEDELVE_PORT"); !success {
		log.Fatal("No port number specified.")
	}

	if mongoAddress, success = os.LookupEnv("WEDELVE_DB_ADDR"); !success {
		log.Fatal("No database address specified.")
	}

	if db, err = OpenDatabase(mongoAddress); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	address = fmt.Sprintf(":%s", port)

	server := Server(address)
	server.ListenAndServe()
}

func Server(address string) http.Server {
	return http.Server{
		Addr:    address,
		Handler: router(),
		TLSConfig: &tls.Config{
			MinVersion:               tls.VersionTLS13,
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			},
		},
	}
}
