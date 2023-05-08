package processor

import (
	"data-processing/models"
	"github.com/kljensen/snowball"
	"sort"
	"strings"
)

func ExtractSentencesAdvance(text string) []models.WeightSentence {
	sentences := strings.Split(text, ". ")
	var result []models.WeightSentence
	for _, sent := range sentences {
		result = append(result, models.WeightSentence{Text: strings.TrimSpace(sent)})
	}

	return result
}

func TextRank(sentences []models.WeightSentence, language string, damping, epsilon float64, maxIter int) []models.WeightSentence {
	matrix := buildSimilarityMatrix(sentences, language)
	rankVector := powerIteration(matrix, damping, epsilon, maxIter)

	for i, rank := range rankVector {
		sentences[i].Weight = rank
	}

	return sentences
}

func buildSimilarityMatrix(sentences []models.WeightSentence, language string) [][]float64 {
	n := len(sentences)
	matrix := make([][]float64, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]float64, n)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			score := similarity(sentences[i].Text, sentences[j].Text, language)
			matrix[i][j] = score
			matrix[j][i] = score
		}
	}

	return matrix
}

func similarity(s1, s2, language string) float64 {
	words1 := strings.Split(s1, " ")
	words2 := strings.Split(s2, " ")

	stemmedWords1 := stemWords(words1, language)
	stemmedWords2 := stemWords(words2, language)

	intersect := intersection(stemmedWords1, stemmedWords2)
	union := union(stemmedWords1, stemmedWords2)

	return float64(len(intersect)) / float64(len(union))
}

func stemWords(words []string, language string) []string {
	var stemmedWords []string
	for _, word := range words {
		stem, _ := snowball.Stem(word, language, true)
		stemmedWords = append(stemmedWords, stem)
	}
	return stemmedWords
}

func intersection(a, b []string) []string {
	var result []string
	for _, elemA := range a {
		for _, elemB := range b {
			if elemA == elemB {
				result = append(result, elemA)
				break
			}
		}
	}
	return result
}

func union(a, b []string) []string {
	m := make(map[string]bool)
	for _, elem := range a {
		m[elem] = true
	}
	for _, elem := range b {
		m[elem] = true
	}

	var result []string
	for elem := range m {
		result = append(result, elem)
	}
	return result
}

func powerIteration(matrix [][]float64, damping, epsilon float64, maxIter int) []float64 {
	n := len(matrix)
	rankVector := make([]float64, n)
	for i := range rankVector {
		rankVector[i] = 1.0 / float64(n)
	}
	for iter := 0; iter < maxIter; iter++ {
		newRankVector := make([]float64, n)
		for i := range newRankVector {
			newRankVector[i] = (1.0 - damping) / float64(n)
		}

		for i := range matrix {
			for j := range matrix[i] {
				newRankVector[i] += damping * matrix[i][j] * rankVector[j]
			}
		}

		delta := 0.0
		for i := range rankVector {
			delta += (rankVector[i] - newRankVector[i]) * (rankVector[i] - newRankVector[i])
		}

		rankVector = newRankVector
		if delta < epsilon {
			break
		}
	}

	return rankVector

}

func TopNSentences(sentences []models.WeightSentence, n int) []models.WeightSentence {
	sort.Slice(sentences, func(i, j int) bool {
		return sentences[i].Weight > sentences[j].Weight
	})
	if n > len(sentences) {
		n = len(sentences)
	}
	return sentences[:n]
}
