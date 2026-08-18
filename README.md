# Stock-Price-Viewer

A simple command-line tool, written in Go, that looks up a stock ticker and displays its live price and basic financial metrics using the [Finnhub](https://finnhub.io/) API.

This project was built as a hands-on exercise in Go fundamentals: REST API calls, JSON parsing, error handling, environment-based configuration, and unit testing (including mocking an external API with `httptest`).

## Features

- Interactive CLI loop — keeps prompting for a ticker until you type `exit`
- Validates ticker input (1–5 letters) before making any network call
- Fetches the live quote (current price, change, high/low, open, previous close) from Finnhub
- Detects and reports unknown tickers, rather than silently showing `$0.00`
- API key is never hardcoded or committed — loaded from an environment variable

## Prerequisites

- [Go](https://go.dev/dl/) 1.26 or later
- A free Finnhub API key (see below)

## Getting a Finnhub API Key

1. Go to [finnhub.io/register](https://finnhub.io/register) and create a free account.
2. Once logged in, your API key is shown on the [Dashboard](https://finnhub.io/dashboard) page.
3. Copy that key — you'll need it in the next step.

The free tier is more than enough to run this program.

## Setting Up Your Environment Variable

This program reads your API key from an environment variable named `FINNHUB_API_KEY`. It is **never** stored in a file that gets committed to this repository.

### Linux / macOS / WSL (bash or zsh)

For a quick, one-off test in your current terminal session:

```bash
export FINNHUB_API_KEY="your_api_key_here"
```

To make it permanent (available every time you open a new terminal), add that same line to your shell's startup file, then reload it:

```bash
# bash
echo 'export FINNHUB_API_KEY="your_api_key_here"' >> ~/.bashrc
source ~/.bashrc

# zsh
echo 'export FINNHUB_API_KEY="your_api_key_here"' >> ~/.zshrc
source ~/.zshrc
```

Verify it's set:

```bash
echo $FINNHUB_API_KEY
```

### Windows (PowerShell)

```powershell
setx FINNHUB_API_KEY "your_api_key_here"
```

Note: `setx` sets the variable for future terminal sessions — you'll need to open a new PowerShell window for it to take effect.

## Installation

```bash
git clone https://github.com/tylerheller27/Stock-Price-Viewer.git
cd Stock-Price-Viewer
```

## Usage

Run the program from the repository root:

```bash
go run ./cmd/stockviewer
```

Or build a binary and run that:

```bash
go build -o stockviewer ./cmd/stockviewer
./stockviewer
```

Example session:

```
Please Enter A Stock Ticker
Please Type Exit to Terminate Program
VOO
Ticker: VOO
Current Price: 710.27
Change: -3.34 (-0.47%)
Open: 713.50
High: 714.04
Low: 710.10
Previous Close: 713.61
ZZZZZ
Ticker not found
exit
Goodbye!
```

If `FINNHUB_API_KEY` isn't set, the program will tell you immediately and exit rather than attempting any API calls.

## Running Tests

All tests run entirely offline — network calls to Finnhub are mocked with `httptest`, so no API key is required to run the test suite.

```bash
go test ./...
```

For verbose, per-test output:

```bash
go test -v ./...
```

## Project Structure

```
cmd/stockviewer/
├── main.go               # Entry point; wires everything together in the input loop
├── displayCLI.go          # Reads a line of input from the user, detects "exit"
├── inputValidation.go     # Validates that a ticker is 1-5 letters
├── config.go               # Loads the API key from the environment
├── getQuote.go             # Calls Finnhub's /quote endpoint and parses the response
├── printStockInfo.go       # Formats and prints the fetched quote
└── *_test.go               # Unit tests for the corresponding file
```

## License

This is a personal learning project and does not currently include a license.
