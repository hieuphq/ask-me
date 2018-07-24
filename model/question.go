package model

// QuestionType .
type QuestionType interface{}

// Answer .
type Answer string

// Question .
type Question interface {
	GetQuestion() (string, error)
	GetAnswers() ([]Answer, error)
	GetQuestionType() (QuestionType, error)
	GetRating() (int, error)
	UpVote(user User) error
	DownVote(user User) error
}
