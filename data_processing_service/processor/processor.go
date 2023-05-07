package processor

import (
	"data-processing/models"

	"fmt"
)

type SimpleHandler struct{}

func (h *SimpleHandler) HandleMessage(event models.Review) error {

	keywords, err := ExtractKeyword(event.Review, 3)
	if err != nil {
		fmt.Println("Error")
	}

	fmt.Println(keywords)
	return nil
}
