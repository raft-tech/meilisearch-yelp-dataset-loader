package main

import (
	"encoding/json"
	"fmt"
	"github.com/meilisearch/meilisearch-go"
	"io"
	"log"
	"os"
)

type Input struct {
	Index      string
	Filename   string
	PrimaryKey string
	BatchSize  int
}

type Result struct {
	Index      string
	PrimaryKey string
	BatchSize  int
	JSON       []map[string]interface{}
}

func decodeFile(input Input, channel chan Result) {
	file, err := os.Open(input.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Printf("Successfully opened `%s`\n", input.Filename)

	decoder := json.NewDecoder(file)

	var docs []map[string]interface{}

	for {
		var doc map[string]interface{}

		err := decoder.Decode(&doc)
		if err == io.EOF {
			// all done
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		docs = append(docs, doc)
	}

	channel <- Result{
		Index:      input.Index,
		PrimaryKey: input.PrimaryKey,
		BatchSize:  input.BatchSize,
		JSON:       docs,
	}
}

func main() {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://localhost:7700",
		APIKey: "aSampleMasterKey",
	})

	results := make(chan Result)

	inputs := []Input{
		{"businesses", "./data/yelp_academic_dataset_business.json", "business_id", 50000},
		{"checkins", "./data/yelp_academic_dataset_checkin.json", "business_id", 10000},
		{"reviews", "./data/yelp_academic_dataset_review.json", "review_id", 10000},
		{"users", "./data/yelp_academic_dataset_user.json", "user_id", 10000},
	}

	for _, input := range inputs {
		go decodeFile(input, results)
	}

	for i := 0; i < len(inputs); i++ {
		result := <-results
		log.Printf("Adding %d %s documents in batches of %d", len(result.JSON), result.Index, result.BatchSize)
		if responses, err := client.Index(result.Index).AddDocumentsInBatches(result.JSON, result.BatchSize, result.PrimaryKey); err != nil {
			log.Fatalf("Failed to unmarshal message value into json: %s", err)
		} else {
			for i, resp := range responses {
				log.Printf("Waiting for %s task %d to be completed...", result.Index, i)
				if t, err := client.WaitForTask(resp.TaskUID); err != nil {
					if err.Error() == "context deadline exceeded" {
						continue
					}
					log.Fatalf("Task %d failed to be processed: %s", i, err)
				} else {
					log.Printf("Task %d complete. Took %s", i, t.Duration)
				}
			}
		}
	}
	log.Println("done")
}
