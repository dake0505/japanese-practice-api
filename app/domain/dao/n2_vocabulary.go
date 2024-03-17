package dao

type N2VocabularySubject struct {
	Id            int    `gorm:"column:id; primary_key; not null"`
	QuestionId    int    `gorm:"column:questionid"`
	QuestionTitle string `gorm:"column:questiontitle"`
	BaseModel
}

type N2VocabularySubjectOption struct {
	Id          int    `gorm:"column:id; primary_key; not null"`
	QuestionId  int    `gorm:"column:questionid"`
	OptionId    int    `gorm:"column:optionid"`
	OptionTitle string `gorm:"column:optiontitle"`
	IsAnswer    bool   `gorm:"column:isanswer"`
	BaseModel
}

func (N2VocabularySubject) TableName() string {
	return "n2_vocabulary_subject"
}

func (N2VocabularySubjectOption) TableName() string {
	return "n2_vocabulary_subject_option"
}
