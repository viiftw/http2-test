package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

const url = "https://localhost:9080"

func main() {
	// client := &http.Client{}

	// Create a pool with the server certificate since it is not signed
	// by a known CA
	caCert, err := ioutil.ReadFile("../cert/server.crt")
	if err != nil {
		log.Fatalf("Reading server certificate: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create TLS configuration with the certificate of the server
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	client := http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
	resp, _ := client.Get(url)
	fmt.Printf("Client Proto: %d\n", resp.ProtoMajor)
}
