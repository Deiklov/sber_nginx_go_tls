package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// load client cert
	home, _ := os.Getwd()
	err := os.Chdir(filepath.Join(home, "certs"))
	if err != nil {
		log.Fatal(err)
	}

	cert, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		log.Fatal(err)
	}

	// load CA cert
	caCert, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// https client tls config

	// export GODEBUG="x509ignoreCN=0"
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}

	transport := &http.Transport{TLSClientConfig: tlsConfig}

	// https client request
	url := "https://localhost:8443/"

	//пример любого запроса
	j := []byte(`{"id": "3232323", "name": "lambda"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
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
