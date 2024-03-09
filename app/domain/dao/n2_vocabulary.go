package dao

type N2Vocabulary struct {
	Id         int `gorm:"column:id; primary_key; not null"`
	QuestionId int `gorm:"column:questionid"`
	// BaseModel
}
