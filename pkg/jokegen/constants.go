package jokegen

// API endpoints for different joke services
const (
	// Official Joke API - Random joke endpoint
	OfficialJokeAPIURL = "https://official-joke-api.appspot.com/random_joke"
	
	// Joke of the Day API
	JokeOfDayURL = "https://jokes.one/joke/random"
	
	// Default timeout in seconds
	DefaultTimeout = 10
)

// JokeType represents the type of joke
type JokeType string

const (
	TypeGeneral JokeType = "general"
	TypeKnock   JokeType = "knock-knock"
	TypeProgram JokeType = "programming"
)