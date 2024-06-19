package dto

type CreateItemRequest struct {
	QuestionTitle string `json:"questionTitle"`
	AnswerId      string `json:"answerId"`
}

type UpdateItemRequest struct {
	ID            uint   `json:"id"`
	QuestionID    string `json:"questionId"`
	QuestionTitle string `json:"questionTitle"`
	AnswerId      string `json:"answerId"`
}

type CreateTypeDto struct {
	TypeName string `json:"typeName"`
}
