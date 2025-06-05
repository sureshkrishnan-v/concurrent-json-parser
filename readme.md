## JSON Validator
   This Go program recursively scans the tests directory for .json files, then concurrently validates each JSON file using a custom lexer and parser. It prints the validation results for each file.

## Features
 - Recursively finds all .json files under the tests directory.
 - Concurrently validates multiple JSON files to speed up processing.
 - Prints whether each file contains valid or invalid JSON.
 - Graceful error handling for file read failures.

## Prerequisites
Go 1.18+ installed and configured in your system.

Your project must have implemented the NewLexer and NewParser functions (for tokenizing and parsing JSON).

## How to Run
Place the .json files you want to validate inside the tests folder or its subfolders.

## Run the program:
```go
go run .
```


## Example:
✅ tests/config.json: Valid JSON

❌ tests/invalid.json: Invalid JSON

❌ tests/missing.json: Read error: open tests/missing.json: no such file or directory

## Code Overview
Uses filepath.Walk to find .json files.

Uses goroutines and a WaitGroup to validate files concurrently.

Results from all validations are sent through a channel and printed sequentially.

The validation logic is inside validateFile, which reads the file, lexes, parses, and returns a validation message.

## Notes
Ensure your lexer and parser handle JSON format correctly.

The tests directory path is relative to where you run the program.

Modify the directory path in filepath.Walk("tests", ...) if needed.