package network

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func BuildRequest(method string, url string) (*http.Request, error) {
	r, err := http.NewRequest(method, url, nil)
	r.Header.Set("User-Agent", "Blazer")
	return r, err
}

func Download(url string, thread int) {
	fmt.Println("Initiating download... dispatching workers")
	r, _ := BuildRequest(http.MethodGet, url)
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, _ := ioutil.ReadAll(resp.Body)
	ioutil.WriteFile("out/test.pdf", data, os.ModePerm)
}
