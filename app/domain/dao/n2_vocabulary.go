package dao

type N2VocabularySubject struct {
	Id            int    `gorm:"column:id; primary_key; not null"`
	QuestionId    int    `gorm:"column:questionid"`
	QuestionTitle string `gorm:"column:questiontitle"`
	BaseModel
}

type N2VocabularySubjectDetail struct {
}

func (N2VocabularySubject) TableName() string {
	return "n2_vocabulary_subject"
}
