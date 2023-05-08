package processor

import (
	vader "github.com/grassmudhorses/vader-go"
	"math"
)

func AnalyzeSentiment(text string) float64 {
	sentiment := vader.GetSentiment(text)

	return math.Round(sentiment.Compound*1000) / 1000
}
