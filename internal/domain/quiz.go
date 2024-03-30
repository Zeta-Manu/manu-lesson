package domain

type QuizQuestion struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	ID       int    `json:"id" binding:"required"`
}

type QuizQuestionWithVideo struct {
	VideoURL string
	QuizQuestion
}

type Quiz struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
