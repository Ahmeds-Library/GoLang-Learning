package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

	func main() {
	start := time.Now()

	content, err := os.ReadFile("../File.txt") 
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	text := string(content)
	fmt.Println("Total characters:", len(text))

	//****************************************************(Chunking of file)******************************************************************
	chunkCount := 20
	chunkSize := len(text) / chunkCount
	chunks := make([]string, chunkCount)

	for i := 0; i < chunkCount; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == chunkCount-1 {
			end = len(text)
		}
		chunks[i] = text[start:end]
	}

	//****************************************************(Apply Routines)******************************************************************
	var wg sync.WaitGroup
	results := make([]map[string]int, chunkCount)
	for i := 0; i < chunkCount; i++ {
		wg.Add(1)
		go func(index int, chunk string) {
			defer wg.Done()
			results[index] = analyze(chunk)
		}(i, chunks[i])
	}

	wg.Wait()

	//****************************************************(Combining Result)******************************************************************
	finalResult := map[string]int{
		"words":        0,
		"lines":        0,
		"spaces":       0,
		"vowels":       0,
		"digits":       0,
		"paragraphs":   0,
		"punctuations": 0,
		"special":      0,
		"consonants":   0,
	}

	for _, result := range results {
		for key, value := range result {
			finalResult[key] += value
		}
	}

	fmt.Println("Final analysis:", finalResult)
	fmt.Println("Total time:", time.Since(start))
}

// ****************************************************(Analyze the chunks)******************************************************************
func analyze(text string) map[string]int {
	counts := map[string]int{
		"words":        0,
		"lines":        0,
		"spaces":       0,
		"vowels":       0,
		"digits":       0,
		"paragraphs":   0,
		"punctuations": 0,
		"special":      0,
		"consonants":   0,
	}

	for i := 0; i < len(text); i++ {
		c := text[i]
		switch {
		case c == ' ':
			counts["spaces"]++
			counts["words"]++
		case c == '\n':
			counts["lines"]++
			counts["words"]++
			counts["paragraphs"]++
		case c == '\t':
			counts["words"]++
		case c == '.' || c == ',' || c == '!' || c == '?' || c == ';' || c == ':' || c == ')' || c == '(' || c == '-' || c == '_' || c == '"' || c == '\'':
			counts["punctuations"]++
		case c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U':
			counts["vowels"]++
		case c >= '0' && c <= '9':
			counts["digits"]++
		case c == '@' || c == '#' || c == '$' || c == '%' || c == '^' || c == '&' || c == '*' || c == '+' || c == '=' || c == '{' || c == '}' || c == '[' || c == ']':
			counts["special"]++
		case (c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z') && !(c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'):
			counts["consonants"]++
		}
	}
	return counts
}
