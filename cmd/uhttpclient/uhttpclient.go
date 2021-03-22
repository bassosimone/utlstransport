package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/bassosimone/utlstransport"
	utls "github.com/refraction-networking/utls"
)

func main() {
	clnt := &http.Client{Transport: &utlstransport.UHTTPTransport{
		UTLSClientHelloID: &utls.HelloFirefox_63,
	}}
	for _, target := range os.Args[1:] {
		log.Printf("> GET %s", target)
		resp, err := clnt.Get(target)
		if err != nil {
			log.Printf("< err=%s", err)
			continue
		}
		log.Printf("< %d", resp.StatusCode)
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("ioutil.ReadAll err=%s", err)
			resp.Body.Close()
			continue
		}
		log.Printf("ioutil.ReadAll bytes=%d", len(data))
	}
}
