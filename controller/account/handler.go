package c_account

import (
	"net/http"

	"github.com/labstack/echo/v4"

	r_account "movie-api/controller/account/response"
	m_account "movie-api/model/account"
)

func GetAllAccount(c echo.Context) error {
	result, err := m_account.SelectAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "hope all feeling well",
		"data":    r_account.FromModelSlice(result),
	})
}

// TRY FROM LINE 26 - 67
// func AddAccount(c echo.Context) error {
// 	newAccount := Account{}
// 	c.Bind(&newAccount)

// 	if err := DB.Save(&newAccount).Error; err != nil {
// 		return c.JSON(http.StatusNotFound, err)
// 	}

// 	// data = append(data, newAccount)
// 	return c.JSON(http.StatusAccepted, map[string]interface{}{
// 		"message": "hope all feeling well",
// 		"data":    newAccount,
// 	})
// }

// func UpdateAccount(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	reqAccount := Account{}
// 	c.Bind(&reqAccount)

// 	// var dataAccount Account
// 	// reqAccount.ID = uint(id)

// 	// DB.Where("id = ?", id).First(&dataAccount)

// 	// if reqAccount.Username != "" {
// 	// 	dataAccount.Username = reqAccount.Username
// 	// }

// 	if err := DB.Where("id = ?", id).Updates(&reqAccount).Error; err != nil {
// 		return c.JSON(http.StatusBadRequest, err)
// 	}

// 	var dataAccount Account
// 	reqAccount.ID = uint(id)

// 	return c.JSON(http.StatusAccepted, map[string]interface{}{
// 		"message": "hope all feeling well",
// 		"data":    dataAccount,
// 	})
// }

// func getAccount(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	if id <= len(data) && id > 0 {
// 		return c.JSON(http.StatusOK, data[id-1])
// 	} else {
// 		return c.JSON(http.StatusOK, data)
// 	}
// }
