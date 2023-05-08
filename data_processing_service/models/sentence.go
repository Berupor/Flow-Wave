package models

type Sentence struct {
	Text string
	Len  int
}

type WeightSentence struct {
	Text   string  `json:"text"`
	Weight float64 `json:"weight"`
}
