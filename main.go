package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
)

func main() {
	commit := flag.String("commit", "", "commit hash")
	pipelineURL := flag.String("pipeline", "", "pipeline URL")
	flag.Parse()

	log.Println("listening on :8080")

	err := http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type resp struct {
			Commit      string `json:"commit"`
			PipelineURL string `json:"pipeline_url"`
		}
		data, err := json.Marshal(resp{
			Commit:      *commit,
			PipelineURL: *pipelineURL,
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("error while encoding response:", err)
			return
		}
		_, err = w.Write(data)
		if err != nil {
			log.Println("error while writing response:", err)
			return
		}
	}))
	log.Fatalln("error while listening on :8080", err)
}
