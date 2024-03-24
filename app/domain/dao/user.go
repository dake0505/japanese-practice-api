package dao

type User struct {
	ID       string `gorm:"column:id; primary_key; not null" json:"id"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password;->:false" json:"-"`
	BaseModel
}
