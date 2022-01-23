package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// I didn't type all this out. It was auto-generated from an actual JSON object by GoLand.
type CircleCIPayload struct {
	Type       string    `json:"type"`
	Id         string    `json:"id"`
	HappenedAt time.Time `json:"happened_at"`
	Webhook    struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"webhook"`
	Workflow struct {
		Id        string    `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		StoppedAt time.Time `json:"stopped_at"`
		Url       string    `json:"url"`
		Status    string    `json:"status"`
	} `json:"workflow"`
	Pipeline struct {
		Id        string    `json:"id"`
		Number    int       `json:"number"`
		CreatedAt time.Time `json:"created_at"`
		Trigger   struct {
			Type string `json:"type"`
		} `json:"trigger"`
		Vcs struct {
			ProviderName        string `json:"provider_name"`
			OriginRepositoryUrl string `json:"origin_repository_url"`
			TargetRepositoryUrl string `json:"target_repository_url"`
			Revision            string `json:"revision"`
			Commit              struct {
				Subject string `json:"subject"`
				Body    string `json:"body"`
				Author  struct {
					Name  string `json:"name"`
					Email string `json:"email"`
				} `json:"author"`
				AuthoredAt time.Time `json:"authored_at"`
				Committer  struct {
					Name  string `json:"name"`
					Email string `json:"email"`
				} `json:"committer"`
				CommittedAt time.Time `json:"committed_at"`
			} `json:"commit"`
			Branch string `json:"branch"`
		} `json:"vcs"`
	} `json:"pipeline"`
	Project struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"project"`
	Organization struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"organization"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		body := CircleCIPayload{}
		err := decoder.Decode(&body)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%+v\n", body)
	})

	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
