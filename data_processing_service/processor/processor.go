package processor

import (
	"data-processing/models"

	"fmt"
)

type SimpleHandler struct{}

func (h *SimpleHandler) HandleMessage(event models.Review) error {

	keywords, _ := ExtractKeyword(event.Review, 3)
	fmt.Println(keywords)

	sentiment := AnalyzeSentiment(event.Review)
	fmt.Println("Sentiment: %v", sentiment)

	sentences := ExtractSentencesAdvance(event.Review)
	rankedSentences := TextRank(sentences, "english", 0.85, 0.0001, 100)
	topSentences := TopNSentences(rankedSentences, 2)
	fmt.Println(topSentences)

	return nil
}
