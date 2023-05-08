package processor

import (
	"sort"
	"strings"

	"github.com/jdkato/prose/v2"
)

func getTopNWords(words []string, n int) []string {
	wordCounts := make(map[string]int)

	for _, word := range words {
		wordCounts[word]++
	}

	type wordFrequency struct {
		word  string
		count int
	}

	frequencies := make([]wordFrequency, 0, len(wordCounts))
	for word, count := range wordCounts {
		frequencies = append(frequencies, wordFrequency{word, count})
	}

	sort.Slice(frequencies, func(i, j int) bool {
		return frequencies[i].count > frequencies[j].count
	})

	topWords := make([]string, 0, n)
	for i, freq := range frequencies {
		if i >= n {
			break
		}
		topWords = append(topWords, freq.word)
	}

	return topWords
}

func ExtractKeyword(text string, n int) ([]string, error) {
	// Create new text document
	doc, err := prose.NewDocument(text)
	if err != nil {
		return nil, err
	}

	// Extract keywords
	var keywords []string
	for _, tok := range doc.Tokens() {
		if tok.Tag == "NN" || tok.Tag == "NNS" || tok.Tag == "NNP" || tok.Tag == "NNPS" {
			keywords = append(keywords, strings.ToLower(tok.Text))
		}
	}

	// Get top N words
	topWords := getTopNWords(keywords, n)

	return topWords, nil
}
