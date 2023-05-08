package processor

import (
	"data-processing/models"
	"data-processing/storage"
	"log"
)

type ReviewAnalyzerHandler struct {
	Storage storage.Storage
}

func (h *ReviewAnalyzerHandler) HandleMessage(event models.Review) error {
	keywords, err := ExtractKeyword(event.Review, 3)
	if err != nil {
		log.Printf("Error extracting keywords: %v", err)
	}

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

	if err := h.Storage.CreateReview(reviewCreate); err != nil {
		log.Printf("Error adding review: %v", err)
		return err
	}

	return nil
}
