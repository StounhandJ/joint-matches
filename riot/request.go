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
	q.Add("api_key", "RGAPI-a82f7846-be49-44ae-b695-0eb0483a43c5")

	if data != nil {
		for key, value := range data {
			q.Add(key, value)
		}
	}
	req.URL.RawQuery = q.Encode()

nextCheck:
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode == 429 {
		fmt.Println("Request: go sleep")
		time.Sleep(10 * time.Second)
		goto nextCheck
	}

	return resp.Body
}
