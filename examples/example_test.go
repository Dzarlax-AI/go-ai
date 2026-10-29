// Package examples demonstrates generation with a synthetic HTTP transport.
package examples_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	ai "github.com/Dzarlax-AI/go-ai"
	"github.com/Dzarlax-AI/go-ai/providers/openai"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func client(response string) *openai.Client {
	return openai.New(ai.Config{APIKey: "example-key", HTTPClient: &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: http.StatusOK, Header: make(http.Header), Body: io.NopCloser(strings.NewReader(response))}, nil
	})}})
}

func Example_text() {
	c := client(`{"status":"completed","output":[{"content":[{"type":"output_text","text":"Hello!"}]}]}`)
	response, err := c.Generate(context.Background(), ai.Request{Model: "example-model", Prompt: "Say hello."})
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Text)
	// Output: Hello!
}

func Example_structuredJSON() {
	c := client(`{"status":"completed","output":[{"content":[{"type":"output_text","text":"{\"greeting\":\"Hello!\"}"}]}]}`)
	schema := json.RawMessage(`{"type":"object","properties":{"greeting":{"type":"string"}},"required":["greeting"]}`)
	response, err := c.Generate(context.Background(), ai.Request{Model: "example-model", Prompt: "Return a greeting.", Schema: schema})
	if err != nil {
		panic(err)
	}
	var result struct {
		Greeting string `json:"greeting"`
	}
	if err := json.Unmarshal([]byte(response.Text), &result); err != nil {
		panic(err)
	}
	fmt.Println(result.Greeting)
	// Output: Hello!
}
