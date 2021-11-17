package migrate

import (
	// Import files
	"movie-api/config"
	m_account "movie-api/model/account"
)

func AutoMigrate() {
	config.DB.AutoMigrate(&m_account.Account{}) //Automatically checking for changes in struct
}
