package domain

type QuizQuestion struct {
	Question string
	Answer   string
	ID       int
}

type QuizQuestionWithVideo struct {
	VideoURL string
	QuizQuestion
}
