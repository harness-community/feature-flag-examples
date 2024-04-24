package main

import (
	"log"
	"net/http"
	"os"

	harness "github.com/harness/ff-golang-server-sdk/client"
)

var (
	apiKey            string = os.Getenv("FF_SDK_KEY")
	connectionAddress string = os.Getenv("RELAY_PROXY_ADDRESS")
)

func main() {
	log.Println("Harness SDK Getting Started")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {

		log.Printf("Headers: %s", r.Header)

		client, err := harness.NewCfClient(apiKey, harness.WithURL(connectionAddress), harness.WithEventsURL(connectionAddress), harness.WithWaitForInitialized(true), harness.WithMaxAuthRetries(5))

		if err != nil {
			log.Printf("could not connect to FF endpoint %s: %s\n", connectionAddress, err)
			w.WriteHeader(500)
		} else {
			log.Printf("connected to FF endpoint %s: %s\n", connectionAddress, err)
			w.WriteHeader(http.StatusOK)
		}

		client.Close()
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
