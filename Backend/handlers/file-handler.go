package handlers

import (
	"net/http"

	"github.com/apo-theddy/file-manager/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type FileHandler struct {
	Db *gorm.DB
}

func (fh FileHandler) GetFilesHandler(c echo.Context) error {
	var files []models.File
	fh.Db.Where("DeleteDate IS NULL").Find(&files)
	return c.JSON(http.StatusOK, files)
}
