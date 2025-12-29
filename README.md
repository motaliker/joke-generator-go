# Random Joke Generator in Go

A simple yet feature-rich Go application that fetches random jokes from the Official Joke API. It demonstrates HTTP requests, JSON parsing, and testing best practices.

## Features

- 🎭 **Fetch Random Jokes** - Get a random joke with one simple call
- 🎯 **Fetch by Type** - Get jokes of specific types (general, knock-knock, programming)
- 🔍 **Fetch by ID** - Retrieve a specific joke by its ID
- 📚 **Fetch Multiple** - Get multiple jokes at once
- ✅ **Comprehensive Tests** - Full test coverage with unit tests
- 🛠️ **Error Handling** - Robust error handling and validation
- ⚡ **Timeout Support** - Configurable timeouts for API requests

## Project Structure

```
joke-generator-go/
├── pkg/
│   └── jokegen/
│       ├── jokegen.go         # Core implementation
│       └── constants.go       # API endpoints and constants
├── examples/
│   └── main.go               # Usage examples
├── test/
│   └── jokegen_test.go       # Unit tests
├── go.mod                    # Go module definition
├── .gitignore               # Git ignore file
└── README.md               # This file
```

## Quick Start

### Installation

```bash
# Clone the repository
git clone https://github.com/motaliker/joke-generator-go.git
cd joke-generator-go

# Download dependencies
go mod tidy
```

### Basic Usage

```bash
# Run the example
go run examples/main.go
```

### Running Tests

```bash
# Run all tests
go test -v ./test

# Run tests with coverage
go test -v -cover ./test

# Run specific test
go test -v -run TestFetchRandomJoke ./test
```

## API Documentation

### JokeGenerator

```go
// Create a new generator
gen := jokegen.NewJokeGenerator()

// Fetch a random joke
joke, err := gen.FetchRandomJoke()

// Fetch a joke by type
knockJoke, err := gen.FetchJokesByType("knock-knock")

// Fetch a joke by ID
specificJoke, err := gen.FetchJokeByID(1)

// Fetch multiple jokes
jokes, err := gen.FetchMultipleJokes(5)
```

### Joke Structure

```go
type Joke struct {
    ID        int    // Unique identifier
    Type      string // Type of joke (general, knock-knock, programming)
    Setup     string // The setup/question
    Punchline string // The punchline/answer
}

// Format as readable string
joke.Format()  // Returns formatted output with emoji and details
joke.String()  // Returns simple setup/punchline format
```

## Usage Examples

### Example 1: Get a Random Joke

```go
generator := jokegen.NewJokeGenerator()
joke, err := generator.FetchRandomJoke()
if err != nil {
    log.Fatal(err)
}
fmt.Println(joke.Format())
```

### Example 2: Get Jokes by Type

```go
// Available types: "general", "knock-knock", "programming"
joke, err := generator.FetchJokesByType("knock-knock")
if err != nil {
    log.Fatal(err)
}
fmt.Println(joke.Setup)
fmt.Println(joke.Punchline)
```

### Example 3: Fetch Multiple Jokes

```go
jokes, err := generator.FetchMultipleJokes(5)
if err != nil {
    log.Fatal(err)
}

for i, joke := range jokes {
    fmt.Printf("Joke %d: %s\n", i+1, joke.Setup)
}
```

## API Source

This project uses the [Official Joke API](https://official-joke-api.appspot.com/):
- **Base URL**: `https://official-joke-api.appspot.com`
- **Random Joke**: `GET /random_joke`
- **Jokes by Type**: `GET /jokes/{type}/random`
- **Specific Joke**: `GET /jokes/{id}`
- **Multiple Jokes**: `GET /jokes/{count}`

Supported joke types:
- `general` - General jokes
- `knock-knock` - Knock-knock jokes
- `programming` - Programming jokes

## Error Handling

The library handles various error cases:

```go
joke, err := generator.FetchRandomJoke()
if err != nil {
    // Handle errors like network timeouts, API failures, etc.
}
```

## Configuration

You can customize the timeout for API requests:

```go
gen := jokegen.NewJokeGenerator()
// Default timeout is 10 seconds
// Modify by creating custom client if needed
```

## Testing

The project includes comprehensive test coverage:

- `TestFetchRandomJoke` - Tests random joke fetching
- `TestFetchJokesByType` - Tests fetching jokes by type
- `TestFetchJokeByID` - Tests fetching specific joke by ID
- `TestFetchMultipleJokes` - Tests fetching multiple jokes
- `TestFetchMultipleJokesInvalidCount` - Tests error handling
- `TestJokeFormat` - Tests formatting functions

Run tests with:
```bash
go test -v ./test
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - feel free to use this project for any purpose.

## Author

motaliker

## Performance Notes

- Response time: ~500ms - 1s per request (depends on internet connection)
- Rate limiting: The Official Joke API has reasonable rate limits
- Concurrency: Safe to use with goroutines

## Future Enhancements

- [ ] Caching mechanism to reduce API calls
- [ ] Support for multiple joke APIs
- [ ] CLI tool with additional options
- [ ] Joke filtering and search functionality
- [ ] Docker support
- [ ] Concurrent fetching for multiple jokes