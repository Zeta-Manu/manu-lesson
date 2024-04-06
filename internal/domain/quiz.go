package domain

type QuizQuestion struct {
	Question string  `json:"question"`
	Answer   string  `json:"answer"`
	ID       int     `json:"id" binding:"required"`
	VideoURL *string `json:"videoURL,omitempty"`
}

type Quiz struct {
	Question string  `json:"question"`
	Answer   string  `json:"answer"`
	VideoURL *string `json:"videoURL,omitempty"`
}
