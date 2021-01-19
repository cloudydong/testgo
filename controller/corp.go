package controller

import (
	"net/http"

	"github.com/cloudydong/testgo/model"
	"github.com/cloudydong/testgo/storage"
	"github.com/labstack/echo/v4"
)

// GetCorp jsonform get
func GetCorp(c echo.Context) error {
	corp, _ := GetRepoCorp()
	return c.JSON(http.StatusOK, corp)
}

// GetRepoCorp GetDBInstance get corp
func GetRepoCorp() ([]model.Corp, error) {
	db := storage.GetDBInstance()
	corp := []model.Corp{}

	if err := db.Find(&corp).Error; err != nil {
		return nil, err
	}

	return corp, nil
}
