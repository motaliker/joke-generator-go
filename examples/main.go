package main

import (
	"fmt"
	"log"

	"github.com/motaliker/joke-generator-go/pkg/jokegen"
)

func main() {
	// Create a new JokeGenerator instance
	generator := jokegen.NewJokeGenerator()

	fmt.Println("=== Random Joke Generator ===\n")

	// Example 1: Fetch a random joke
	fmt.Println("Example 1: Random Joke")
	fmt.Println("----------------------")
	joke, err := generator.FetchRandomJoke()
	if err != nil {
		log.Fatalf("Error fetching random joke: %v", err)
	}
	fmt.Println(joke.Format())

	// Example 2: Fetch a joke by type
	fmt.Println("Example 2: Knock-Knock Joke")
	fmt.Println("---------------------------")
	knockJoke, err := generator.FetchJokesByType("knock-knock")
	if err != nil {
		log.Fatalf("Error fetching knock-knock joke: %v", err)
	}
	fmt.Println(knockJoke.Format())

	// Example 3: Fetch a joke by ID
	fmt.Println("Example 3: Joke by ID (ID: 1)")
	fmt.Println("-----------------------------")
	specificJoke, err := generator.FetchJokeByID(1)
	if err != nil {
		log.Fatalf("Error fetching specific joke: %v", err)
	}
	fmt.Println(specificJoke.Format())

	// Example 4: Fetch multiple jokes
	fmt.Println("Example 4: Multiple Jokes (5 jokes)")
	fmt.Println("-----------------------------------")
	jokes, err := generator.FetchMultipleJokes(5)
	if err != nil {
		log.Fatalf("Error fetching multiple jokes: %v", err)
	}
	for i, j := range jokes {
		fmt.Printf("Joke %d:\n%s\n", i+1, j.Format())
	}
}