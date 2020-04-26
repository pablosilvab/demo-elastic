package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
)

func Log(index string, logger interface{}) {
	log.SetFlags(0)

	var (
		wg sync.WaitGroup
	)

	// Initialize a client with the default settings.
	//
	// An `ELASTICSEARCH_URL` environment variable will be used when exported.
	//

	urlElastic := validateEnv(os.Getenv("ELASTIC_URL"))

	cfg := elasticsearch.Config{
		Addresses: []string{
			urlElastic,
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	fmt.Println(index)

	wg.Add(1)
	go func() {
		defer wg.Done()

		// Build the request body.

		// Set up the request object.
		body, err := json.Marshal(logger)
		if err != nil {
			log.Fatal("Failed @marshal")
		}

		req := esapi.IndexRequest{
			Index:      index,
			DocumentID: strconv.Itoa(int(time.Now().UnixNano())),
			Body:       bytes.NewReader(body),
			Refresh:    "true",
		}

		// Perform the request with the client.
		res, err := req.Do(context.Background(), es)
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		defer res.Body.Close()

	}()

	wg.Wait()

}

func validateEnv(url string) string {
	if url == "" {
		return "http://localhost:9200/"
	} else {
		return url
	}
}
