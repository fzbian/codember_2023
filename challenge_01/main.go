package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var words []string
	if scanner.Scan() {
		input := scanner.Text()
		words = strings.Fields(input)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	wordsToCount := make(map[string]int)
	var order []string
	for _, s := range words {
		lowerWord := strings.ToLower(s)
		if _, exists := wordsToCount[lowerWord]; !exists {
			order = append(order, lowerWord)
		}
		wordsToCount[lowerWord] += 1
	}

	for _, word := range order {
		fmt.Printf("%s%d", word, wordsToCount[word])
	}
}

