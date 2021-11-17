package m_account

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
