package jokegen

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Joke represents a joke structure
type Joke struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

// JokeGenerator manages joke fetching
type JokeGenerator struct {
	client  *http.Client
	timeout time.Duration
}

// NewJokeGenerator creates a new JokeGenerator instance
func NewJokeGenerator() *JokeGenerator {
	return &JokeGenerator{
		client: &http.Client{
			Timeout: DefaultTimeout * time.Second,
		},
		timeout: DefaultTimeout * time.Second,
	}
}

// FetchRandomJoke fetches a random joke from the Official Joke API
func (jg *JokeGenerator) FetchRandomJoke() (*Joke, error) {
	return jg.fetchFromURL(OfficialJokeAPIURL)
}

// FetchJokesByType fetches a joke of a specific type
func (jg *JokeGenerator) FetchJokesByType(jokeType string) (*Joke, error) {
	url := fmt.Sprintf("https://official-joke-api.appspot.com/jokes/%s/random", jokeType)
	return jg.fetchFromURL(url)
}

// FetchJokeByID fetches a specific joke by ID
func (jg *JokeGenerator) FetchJokeByID(id int) (*Joke, error) {
	url := fmt.Sprintf("https://official-joke-api.appspot.com/jokes/%d", id)
	return jg.fetchFromURL(url)
}

// FetchMultipleJokes fetches multiple jokes at once
func (jg *JokeGenerator) FetchMultipleJokes(count int) ([]*Joke, error) {
	if count <= 0 {
		return nil, fmt.Errorf("count must be greater than 0")
	}
	if count > 10 {
		count = 10 // Limit to 10 per API restrictions
	}

	url := fmt.Sprintf("https://official-joke-api.appspot.com/jokes/%d", count)
	resp, err := jg.makeRequest(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	var jokes []*Joke
	if err := json.Unmarshal(body, &jokes); err != nil {
		return nil, fmt.Errorf("failed to decode jokes: %v", err)
	}

	return jokes, nil
}

// fetchFromURL is a helper method to fetch and decode a joke from any URL
func (jg *JokeGenerator) fetchFromURL(url string) (*Joke, error) {
	resp, err := jg.makeRequest(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var joke Joke
	if err := json.NewDecoder(resp.Body).Decode(&joke); err != nil {
		return nil, fmt.Errorf("failed to decode joke: %v", err)
	}

	return &joke, nil
}

// makeRequest performs an HTTP GET request with proper error handling
func (jg *JokeGenerator) makeRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set User-Agent header to avoid being blocked
	req.Header.Set("User-Agent", "JokeGenerator-Go/1.0")

	resp, err := jg.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch joke: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("API returned status code: %d", resp.StatusCode)
	}

	return resp, nil
}

// String returns a formatted string representation of the joke
func (j *Joke) String() string {
	return fmt.Sprintf("Setup: %s\nPunchline: %s\n", j.Setup, j.Punchline)
}

// Format returns a more readable format
func (j *Joke) Format() string {
	return fmt.Sprintf("😄 Joke #%d (%s)\n%s\n%s\n", j.ID, j.Type, j.Setup, j.Punchline)
}
