package main

import (
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net/http"
)

func main() {
	commit := flag.String("commit", "", "commit hash")
	pipelineURL := flag.String("pipeline", "", "pipeline URL")
	env := flag.String("env", "", "environment")
	flag.Parse()

	log.Println("listening on :8080")

	err := http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := getResponse(*commit, *pipelineURL, *env)
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

func getResponse(commit, pipelineURL, env string) ([]byte, error) {
	if commit == "" || pipelineURL == "" || env == "" {
		return nil, errors.New("commit, pipeline URL and env must be provided")
	}
	type resp struct {
		Commit      string `json:"commit"`
		PipelineURL string `json:"pipeline_url"`
		Env         string `json:"env"`
	}
	return json.Marshal(resp{
		Commit:      commit,
		PipelineURL: pipelineURL,
		Env:         env,
	})
}
