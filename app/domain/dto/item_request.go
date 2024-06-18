package dto

type CreateItemRequest struct {
	QuestionTitle string `json:"questionTitle"`
	AnswerId      string `json:"answerId"`
}
