package download

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func Download(downloadurl string, timeout time.Duration, proxy string) ([]byte, error) {
	client := http.Client{
		Timeout: timeout * time.Second,
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	if proxy != "" {
		proxy, err := url.Parse(proxy)
		if err != nil {
			fmt.Println("Download failed")
			fmt.Println(err)
			return nil, err
		}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxy)}
	}

	resp, err := client.Get(downloadurl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		fmt.Println("Download failed")
		fmt.Println(resp.Status)
		return nil, errors.New(resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read received response")
		fmt.Println(err)
		return nil, err
	}

	return data, nil
}
