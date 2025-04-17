package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

func checkBlacklistedURL(url *url.URL) bool {

	IPs, err := net.LookupIP(url.Hostname())
	if err != nil {
		return true
	}

	for _, IP := range IPs {
		if IP.IsLoopback() || IP.IsPrivate() || IP.IsUnspecified() {
			return true
		}
	}

	return false
}

func urlHandler(w http.ResponseWriter, r *http.Request) {

	url, err := url.Parse(r.URL.Query().Get("url"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid URL"))
		return
	}

	if url.String() == "" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("URL is empty"))
		return
	} else if checkBlacklistedURL(url) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("URL is forbidden"))
		return
	} else {
		httpClient := http.Client{
			Timeout: 5 * time.Second,
		}

		fmt.Printf("Sending request to: %s\n", url)
		resp, err := httpClient.Get(url.String())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Connection error"))
			return
		}

		defer resp.Body.Close()

		httpResponse, _ := io.ReadAll(resp.Body)
		w.Write(httpResponse)
	}

}

func main() {
	http.HandleFunc("/", urlHandler)
	http.ListenAndServe(":8080", nil)
}
