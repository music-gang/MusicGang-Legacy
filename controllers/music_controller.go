package controllers

import (
	"github.com/IacopoMelani/MusicGang/models/table"
	"github.com/labstack/echo"
)

// GetAllMusics - Restituisce tutte le canzoni salvate
func GetAllMusics(c echo.Context) error {

	musicsList, err := table.LoadAllMusic()
	if err != nil {
		return c.JSON(500, Response{
			Status:  1,
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(200, Response{
		Status:  0,
		Success: true,
		Message: "ok!",
		Content: musicsList,
	})
}
