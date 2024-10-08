package repositories

import (
	"fmt"

	"github.com/Zeta-Manu/manu-lesson/internal/adapters/db"
	"github.com/Zeta-Manu/manu-lesson/internal/domain"
)

var _ IQuizRepository = &QuizRepository{}

type IQuizRepository interface {
	GetQuizQuestion(id string) (*domain.QuizQuestion, error)
	PostQuizQuestion(quiz string, answer string, video *string) error
	GetAllQuestions() ([]*domain.QuizQuestion, error)
}

type QuizRepository struct {
	dbAdapter *db.Database
}

func NewQuizRepository(dbAdapter *db.Database) *QuizRepository {
	return &QuizRepository{
		dbAdapter: dbAdapter,
	}
}

func (repo *QuizRepository) GetQuizQuestion(id string) (*domain.QuizQuestion, error) {
	query := "SELECT id, question, answer, video FROM quiz WHERE id = ?"
	rows, err := repo.dbAdapter.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quiz domain.QuizQuestion
	if rows.Next() {
		var videoURL *string
		err = rows.Scan(&quiz.ID, &quiz.Question, &quiz.Answer, &videoURL)
		if err != nil {
			return nil, err
		}
		quiz.VideoURL = videoURL
	} else {
		return nil, err
	}

	return &quiz, nil
}

func (repo *QuizRepository) PostQuizQuestion(quiz string, answer string, video *string) error {
	query := "INSERT INTO quiz (question, answer, video) VALUES (?, ?, ?)"

	_, err := repo.dbAdapter.Exec(query, quiz, answer, video)
	if err != nil {
		return err
	}

	return nil
}

func (repo *QuizRepository) GetAllQuestions() ([]*domain.QuizQuestion, error) {
	query := "SELECT id, question, answer, video FROM quiz"
	rows, err := repo.dbAdapter.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quizQuestions []*domain.QuizQuestion
	for rows.Next() {
		var quiz domain.QuizQuestion
		var videoURL *string
		err = rows.Scan(&quiz.ID, &quiz.Question, &quiz.Answer, &videoURL)
		if err != nil {
			return nil, err
		}
		quiz.VideoURL = videoURL
		quizQuestions = append(quizQuestions, &quiz)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return quizQuestions, nil
}

func (repo *QuizRepository) UpdateQuizQuestion(id string, question *string, answer *string, video *string) error {
	query := "UPDATE quiz SET "
	args := make([]interface{}, 0)

	if question != nil {
		query += "question = ?,"
		args = append(args, *question)
	}
	if answer != nil {
		query += "answer = ?,"
		args = append(args, *answer)
	}
	if video != nil {
		query += "video = ?,"
		args = append(args, *video)
	}

	if question != nil || answer != nil || video != nil {
		query = query[:len(query)-1]
	}

	query += " WHERE id = ?"
	args = append(args, id)

	_, err := repo.dbAdapter.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to update quiz %v query %v", id, query)
	}

	return nil
}
