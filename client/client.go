package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
)

func main() {
	keyLogFile := flag.String("keylog", "", "key log file")
	insecure := flag.Bool("insecure", false, "skip certificate verification")
	flag.Parse()
	urls := []string{"https://localhost1119.ml:8080/"}

	var keyLog io.Writer
	if len(*keyLogFile) > 0 {
		f, err := os.Create(*keyLogFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		keyLog = f
	}

	pool, err := x509.SystemCertPool()
	if err != nil {
		log.Fatal(err)
	}

	var qconf quic.Config

	roundTripper := &http3.RoundTripper{
		TLSClientConfig: &tls.Config{
			RootCAs:            pool,
			InsecureSkipVerify: *insecure,
			KeyLogWriter:       keyLog,
		},
		QuicConfig: &qconf,
	}
	defer roundTripper.Close()
	hclient := &http.Client{
		Transport: roundTripper,
	}

	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, addr := range urls {
		go func(addr string) {
			rsp, err := hclient.Get(addr)
			if err != nil {
				log.Fatal(err)
			}
			body := &bytes.Buffer{}
			_, err = io.Copy(body, rsp.Body)
			if err != nil {
				log.Fatal(err)
			}
			wg.Done()
		}(addr)
	}
	wg.Wait()
}
