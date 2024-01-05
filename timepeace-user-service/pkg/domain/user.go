package domain

type Users struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"validate:required" `
	Email    string `json:"email" gorm:"validate:required,email" `
	PhNo     string `json:"mobile_number" gorm:"validate:required"`
	Password string `json:"password" gorm:"validate:required"`
	Role     string `json:"role"`
	Block    bool   `json:"block" gorm:"default:false"`
}
