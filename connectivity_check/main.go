package main

import (
	"log"
	"net/http"
	"os"

	harness "github.com/harness/ff-golang-server-sdk/client"
	"github.com/harness/ff-golang-server-sdk/evaluation"
)

var (
	apiKey            string            = os.Getenv("FF_SDK_KEY")
	connectionAddress string            = os.Getenv("RELAY_PROXY_ADDRESS")
	flagId            string            = os.Getenv("FF_ID")
	target            evaluation.Target = evaluation.Target{
		Identifier: "ffProxyCanary",
		Name:       "ffProxyCanary",
	}
)

func main() {
	log.Println("Harness SDK Getting Started")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {

		client, err := harness.NewCfClient(
			apiKey,
			harness.WithURL(connectionAddress),
			harness.WithEventsURL(connectionAddress),
			harness.WithWaitForInitialized(true),
			harness.WithMaxAuthRetries(5),
		)

		if err != nil {
			log.Printf("could not connect to FF endpoint %s: %s\n", connectionAddress, err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			log.Printf("connected to FF endpoint %s\n", connectionAddress)

			if flagId != "" {
				result, err := client.BoolVariation(flagId, &target, false)
				if err != nil {
					log.Printf("failed to get evaluation: %s", err)
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					log.Printf("got evaluation: %t", result)
					w.WriteHeader(http.StatusNoContent)
				}
			}
		}

		client.Close()
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
