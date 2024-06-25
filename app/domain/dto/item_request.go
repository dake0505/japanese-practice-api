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

type QuestionDetailDto struct {
	ID            uint         `json:"id"`
	QuestionID    string       `json:"questionId"`
	QuestionTitle string       `json:"questionTitle"`
	AnswerId      string       `json:"answerId"`
	AnswerItems   []AnswerItem `json:"answerItems"`
	NextId        *uint        `json:"nextId"`
	PreId         *uint        `json:"preId"`
}

type AnswerItem struct {
	AnswerId   string `json:"answerId"`
	AnswerDesc string `json:"answerDesc"`
}
