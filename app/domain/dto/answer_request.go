package dto

type CreateAnswerItemDto struct {
	AnswerDesc string `json:"answerDesc"`
	QuestionID string `json:"questionId"`
}
