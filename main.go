package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// CircleCIPayload is the data CircleCI sends as a JSON payload in its POST requests whenever a workflow's status changes (i.e. when it starts, finishes, etc.)
type CircleCIPayload struct {
	Pipeline struct {
		Vcs struct {
			Revision string `json:"revision"`
		} `json:"vcs"`
	} `json:"pipeline"`
}

func isRequestValid(payload []byte, signature string, key []byte) bool {
	asBytes, _ := hex.DecodeString(signature)
	mac := hmac.New(sha256.New, key)
	mac.Write(payload)
	code := mac.Sum(nil)
	return hmac.Equal(code, asBytes)
}

var CircleciSharedSecret = []byte(os.Getenv("CIRCLECI_SECRET"))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var signature string
		sigString := r.Header.Get("CircleCI-Signature")
		sigs := strings.Split(sigString, ",")
		for _, sig := range sigs {
			parts := strings.Split(sig, "=")
			if parts[0] == "v1" {
				signature = parts[1]
			}
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			return
		}

		if !isRequestValid(body, signature, CircleciSharedSecret) {
			return
		}

		decoder := json.NewDecoder(bytes.NewReader(body))
		payload := CircleCIPayload{}
		err = decoder.Decode(&payload)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%+v\n", payload)
	})

	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
