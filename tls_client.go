package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	pool := x509.NewCertPool()
	caCertPath := "./cfssl/ca.pem"
	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	pool.AppendCertsFromPEM(caCrt)

	cliCrt, err := tls.LoadX509KeyPair("cfssl/client.pem", "cfssl/client-key.pem")
	if err != nil {
		fmt.Println(err.Error())
		return
	}


	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get("https://127.0.0.1:8889")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
