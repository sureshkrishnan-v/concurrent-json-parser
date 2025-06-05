package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func main() {

	files := []string{}

	filepath.Walk("tests", func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".json" {
			files = append(files, path)
		}
		return nil
	})

	results := make(chan string)

	var wg sync.WaitGroup
	go func() {
		wg.Wait()
		close(results)
	}()

	for _, file := range files {
		wg.Add(1)

		go func(f string) {
			defer wg.Done()
			res := validateFile(f)
			results <- res
		}(file)
	}
	// Print results sequentially from the channel
	for r := range results {
		fmt.Println(r)
	}
}

func validateFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Sprintf("❌ %s: Read error: %v", path, err)
	}

	lexer := NewLexer(string(content))
	parser := NewParser(lexer.Lex())

	if parser.Parse() {
		return fmt.Sprintf("✅ %s: Valid JSON", path)
	}
	return fmt.Sprintf("❌ %s: Invalid JSON", path)

}
