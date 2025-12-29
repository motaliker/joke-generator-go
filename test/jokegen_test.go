package test

import (
	"testing"

	"github.com/motaliker/joke-generator-go/pkg/jokegen"
)

func TestFetchRandomJoke(t *testing.T) {
	gen := jokegen.NewJokeGenerator()
	joke, err := gen.FetchRandomJoke()

	if err != nil {
		t.Fatalf("FetchRandomJoke failed: %v", err)
	}

	if joke == nil {
		t.Error("Expected joke, got nil")
	}

	if joke.ID == 0 {
		t.Error("Expected non-zero ID")
	}

	if joke.Setup == "" {
		t.Error("Expected non-empty Setup")
	}

	if joke.Punchline == "" {
		t.Error("Expected non-empty Punchline")
	}

	if joke.Type == "" {
		t.Error("Expected non-empty Type")
	}

	t.Logf("Fetched joke: %+v", joke)
}

func TestFetchJokesByType(t *testing.T) {
	gen := jokegen.NewJokeGenerator()
	joke, err := gen.FetchJokesByType("knock-knock")

	if err != nil {
		t.Fatalf("FetchJokesByType failed: %v", err)
	}

	if joke == nil {
		t.Error("Expected joke, got nil")
	}

	if joke.Type != "knock-knock" {
		t.Errorf("Expected type 'knock-knock', got '%s'", joke.Type)
	}

	t.Logf("Fetched knock-knock joke: %+v", joke)
}

func TestFetchJokeByID(t *testing.T) {
	gen := jokegen.NewJokeGenerator()
	joke, err := gen.FetchJokeByID(1)

	if err != nil {
		t.Fatalf("FetchJokeByID failed: %v", err)
	}

	if joke == nil {
		t.Error("Expected joke, got nil")
	}

	if joke.ID != 1 {
		t.Errorf("Expected ID 1, got %d", joke.ID)
	}

	t.Logf("Fetched joke with ID 1: %+v", joke)
}

func TestFetchMultipleJokes(t *testing.T) {
	gen := jokegen.NewJokeGenerator()
	jokes, err := gen.FetchMultipleJokes(3)

	if err != nil {
		t.Fatalf("FetchMultipleJokes failed: %v", err)
	}

	if jokes == nil {
		t.Error("Expected jokes slice, got nil")
	}

	if len(jokes) != 3 {
		t.Errorf("Expected 3 jokes, got %d", len(jokes))
	}

	for i, joke := range jokes {
		if joke.ID == 0 || joke.Setup == "" || joke.Punchline == "" {
			t.Errorf("Joke %d is invalid: %+v", i, joke)
		}
	}

	t.Logf("Fetched %d jokes successfully", len(jokes))
}

func TestFetchMultipleJokesInvalidCount(t *testing.T) {
	gen := jokegen.NewJokeGenerator()
	_, err := gen.FetchMultipleJokes(0)

	if err == nil {
		t.Error("Expected error for count=0")
	}

	_, err = gen.FetchMultipleJokes(-1)
	if err == nil {
		t.Error("Expected error for count=-1")
	}

	t.Log("Invalid count validation passed")
}

func TestJokeFormat(t *testing.T) {
	joke := &jokegen.Joke{
		ID:        1,
		Type:      "general",
		Setup:     "Why don't scientists trust atoms? ",
		Punchline: "Because they make up everything!",
	}

	str := joke.String()
	if str == "" {
		t.Error("Expected non-empty string representation")
	}

	format := joke.Format()
	if format == "" {
		t.Error("Expected non-empty formatted string")
	}

	t.Logf("Formatted joke: %s", format)
}
