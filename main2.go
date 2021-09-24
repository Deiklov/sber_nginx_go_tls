package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// load client cert
	cert, err := tls.LoadX509KeyPair("certslocalhost/client.crt", "certslocalhost/client.key")
	if err != nil {
		log.Fatal(err)
	}

	// load CA cert
	caCert, err := ioutil.ReadFile("certslocalhost/RootCA.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// https client tls config
	// InsecureSkipVerify true means not validate server certificate (so no need to set RootCAs)
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caCertPool,
		InsecureSkipVerify: false,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}

	// https client request
	url := "https://localhost:7070/mef"
	j := []byte(`{"id": "3232323", "name": "lambda"}`)
	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(j))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Transport: transport}

	// read response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	contents, err := ioutil.ReadAll(resp.Body)
	log.Println(string(contents))
}
