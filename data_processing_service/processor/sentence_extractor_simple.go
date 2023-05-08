package processor

import (
	"github.com/jdkato/prose/v2"
	"sort"

	"data-processing/models"
)

// ExtractSentences extracts top N sentences from a given text based on their length.
// It uses a simple algorithm that sorts sentences by their length in descending
// order and returns the top N sentences.
func ExtractSentences(text string, n int) ([]models.Sentence, error) {
	doc, err := prose.NewDocument(text)
	if err != nil {
		return nil, err
	}

	var sentences []models.Sentence
	for _, sent := range doc.Sentences() {
		sentences = append(sentences, models.Sentence{Text: sent.Text, Len: len(sent.Text)})
	}

	sentences = topNSentences(sentences, n)
	return sentences, nil
}

func topNSentences(sentences []models.Sentence, n int) []models.Sentence {
	sort.Slice(sentences, func(i, j int) bool {
		return sentences[i].Len > sentences[j].Len
	})

	if n > len(sentences) {
		n = len(sentences)
	}

	return sentences[:n]
}
