package response

import (
	m_account "movie-api/model/account"
	"time"
)

type Account struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
}

func FromModel(model m_account.Account) Account {
	return Account{
		ID:        int(model.ID),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		Username:  model.Username,
		Password:  model.Password,
		Email:     model.Email,
	}
}

func FromModelSlice(model []m_account.Account) []Account {
	var accountArr []Account
	for key := range model {
		accountArr = append(accountArr, FromModel(model[key]))
	}
	return accountArr
}
