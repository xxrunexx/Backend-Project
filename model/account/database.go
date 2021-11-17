package m_account

import (
	// Import Files
	"movie-api/config"
)

func SelectAll() (account []Account, err error) {
	if err = config.DB.Find(&account).Error; err != nil {
		return []Account{}, err
	}
	return
}
