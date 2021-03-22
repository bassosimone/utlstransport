package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/bassosimone/utlstransport"
	utls "github.com/refraction-networking/utls"
)

func main() {
	clnt := &http.Client{Transport: &utlstransport.UHTTPTransport{
		UTLSClientHelloID: &utls.HelloFirefox_63,
	}}
	wg := &sync.WaitGroup{}
	for _, target := range os.Args[1:] {
		wg.Add(1)
		go func(wg *sync.WaitGroup, target string) {
			defer wg.Done()
			resp, err := clnt.Get(target)
			if err != nil {
				log.Printf("%s: err=%s", target, err)
				return
			}
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("%s: ioutil.ReadAll err=%s", target, err)
				resp.Body.Close()
				return
			}
			log.Printf("%s: ioutil.ReadAll bytes=%d", target, len(data))
		}(wg, target)
	}
	wg.Wait()
}
