package network

import (
	"net"
	"net/http"
	"time"
)

func BuildRequest(method string, url string) (*http.Request, error) {
	r, err := http.NewRequest(method, url, nil)
	r.Header.Set("User-Agent", "Blazer")
	return r, err
}

// To avoid TLS handshake
// https://stackoverflow.com/questions/41719797/tls-handshake-timeout-on-requesting-data-concurrently-from-api
func HTTPClient() *http.Client {
	t := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		// We use ABSURDLY large keys, and should probably not.
		TLSHandshakeTimeout: 600 * time.Second,
	}
	c := &http.Client{
		Transport: t,
	}
	return c
}
