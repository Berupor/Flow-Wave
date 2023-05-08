package processor

import (
	"data-processing/models"
	"data-processing/storage"
)

type SimpleHandler struct {
	Storage storage.Storage
}

func (h *SimpleHandler) HandleMessage(event models.Review) error {
	keywords, _ := ExtractKeyword(event.Review, 3)

	sentiment := AnalyzeSentiment(event.Review)

	sentences := ExtractSentencesAdvance(event.Review)
	rankedSentences := TextRank(sentences, "english", 0.85, 0.0001, 100)
	topSentences := TopNSentences(rankedSentences, 2)

	reviewCreate := models.ReviewCreate{
		ProductID: event.ProductID,
		PlaceID:   event.PlaceID,
		AuthorID:  event.AuthorID,
		Rating:    event.Rating,
		Review:    event.Review,
		Timestamp: event.Timestamp,
		Keywords:  keywords,
		Sentiment: sentiment,
		Sentences: topSentences,
	}

	h.Storage.CreateReview(reviewCreate)

	return nil
}
