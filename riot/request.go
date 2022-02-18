package riot

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func Get(server string, uri string, data map[string]string) io.ReadCloser {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s.api.riotgames.com/%s", server, uri), nil)
	if err != nil {
		return nil
	}

	q := req.URL.Query()
	q.Add("api_key", "RGAPI-68e10570-991c-4acc-8a58-27c1b8c08ab7")

	if data != nil {
		for key, value := range data {
			q.Add(key, value)
		}
	}
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode == 429 {
		time.Sleep(5 * time.Second)
		return Get(server, uri, data)
	}

	return resp.Body
}