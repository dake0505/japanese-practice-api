package dto

type CreateRecordDto struct {
	AnswerId     string `json:"answerId"`
	QuestionId   string `json:"questionId"`
	UserAnswerId string `json:"userAnswerId"`
	RecordType   string `json:"recordType"`
	CreatedBy    string `json:"createdBy"`
}
