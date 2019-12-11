package controllers

import (
	"encoding/json"

	"github.com/IacopoMelani/MusicGang/models/table"
	"github.com/IacopoMelani/musicgang/models/dto"
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

// InsertNewMusic - Controller dedicato all'inserimento di una nuova canzone
func InsertNewMusic(c echo.Context) error {

	decoder := json.NewDecoder(c.Request().Body)

	var m dto.MusicDTO

	if err := decoder.Decode(&m); err != nil {
		return c.JSON(500, Response{
			Status:  1,
			Success: false,
			Message: "Erore lettura JSON",
		})
	}

	if err := table.InsertNewMusic(m); err != nil {
		return c.JSON(500, Response{
			Status:  2,
			Success: false,
			Message: "Erore salvataggio canzone",
		})
	}

	return c.JSON(200, Response{
		Status:  0,
		Success: true,
		Message: "ok",
	})
}
