package exercise

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type TopThreeWords struct {
	First  string
	Second string
	Third  string
}

func wordCount(s string) map[string]int {
	splittedString := strings.Split(s, " ")
	words := make(map[string]int)
	for _, word := range splittedString {
		sw := string(word)
		v, ok := words[sw]
		if ok {
			words[sw] = v + 1
		} else {
			words[sw] = 1
		}
	}
	return words
}

func readFileContentAsString(fname string) (string, error) {
	f, err := os.Open(fname)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("file %s does not exist", fname)
		} else {
			return "", fmt.Errorf("file %s open error: %w", fname, err)
		}
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	// Use a strings.Builder for efficient string concatenation
	var builder strings.Builder

	// Read the file line by line
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
		builder.WriteString("\n") // add newline character to preserve line breaks
	}

	// Get the concatenated string
	fileContent := builder.String()
	return fileContent, nil
}

func getTopThreeWords(countedWords map[string]int) (TopThreeWords, error) {
	if len(countedWords) < 3 {
		return TopThreeWords{}, fmt.Errorf("not enough words")
	}
	// get list of words from countedWords map
	keys := make([]string, 0, len(countedWords))
	for key := range countedWords {
		keys = append(keys, key)
	}
	// sort the list of words using occurances from countedWords mapping
	sort.Slice(keys, func(i, j int) bool { return countedWords[keys[i]] > countedWords[keys[j]] })

	return TopThreeWords{
		First:  keys[0],
		Second: keys[1],
		Third:  keys[2],
	}, nil
}

func CalculateTopThree(fname string) (TopThreeWords, error) {
	fileContent, error := readFileContentAsString(fname)
	if error != nil {
		return TopThreeWords{}, error
	}
	var countedWords map[string]int = wordCount(fileContent)
	return getTopThreeWords(countedWords)
}
