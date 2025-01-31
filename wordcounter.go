package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Reads a file and counts occurrences of specific words
func processFile(filepath string, words []string) map[string]int {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo %s: %v\n", filepath, err)
		os.Exit(1)
	}
	defer file.Close()

	// Map to store the occurrence frequency of each word
	// Keys are the words and values are the counts (initialized to zero)
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word] = 0
	}

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", filepath, err)
		os.Exit(1)
	}

	// Read file
	for scanner.Scan() {
		split := strings.Fields(scanner.Text())
		for _, token := range split {
			if _, exists := wordCount[token]; exists {
				wordCount[token]++
			}
		}
	}

	return wordCount
}

func main() {
	// Directory containing files (provided via command-line)
	dir := os.Args[1]
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}

	// Ask user for words to count
	fmt.Println("Informe as palavras que deseja contabilizar, separadas por espaço:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	words := strings.Split(input, " ")

	for _, word := range words {
		fmt.Printf("\nPalavra: \"%s\"\n", word)
		for _, file := range files {
			if !file.IsDir() {
				results := processFile(filepath.Join(dir, file.Name()), words)
				fmt.Printf("- %s: %d\n", file.Name(), results[word])
			}
		}
	}
}