package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const port = ":8080"

var (
	hostApp1 = "http://localhost:8081"
	hostApp2 = "http://localhost:8082"
	countR   = 0
)

func main() {
	log.Printf("Starting Proxy on %s...", port)
	http.HandleFunc("/", redirect)
	log.Fatal(http.ListenAndServe(port, nil))
}

func redirect(w http.ResponseWriter, r *http.Request) {
	countR++
	if countR%2 == 1 {
		r.Host = hostApp1
	} else {
		r.Host = hostApp2
	}
	u, _ := url.Parse(r.Host)
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.Transport = &myTransport{}
	proxy.ServeHTTP(w, r)
}

type myTransport struct{}

func (t *myTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	response, err := http.DefaultTransport.RoundTrip(request)
	if err != nil {
		return nil, err
	}
	body, err := httputil.DumpResponse(response, true)
	if err != nil {
		return nil, err
	}
	log.Print(string(body))
	return response, err
}
