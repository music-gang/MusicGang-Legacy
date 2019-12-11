package routes

import (
	"github.com/IacopoMelani/musicgang/controllers"
	"github.com/labstack/echo"
)

// InitGetRoutes - Dichiara tutte le route GET
func InitGetRoutes(e *echo.Echo) {
	e.GET("/music/all", controllers.GetAllMusics)
}

// InitPostRoutes - Dichiara tutte le route POST
func InitPostRoutes(e *echo.Echo) {
	e.POST("/music/new", controllers.InsertNewMusic)
}

// InitStaticRoutes - inizializza le route per i file statici
func InitStaticRoutes(e *echo.Echo) {
	e.Static("/", "./dist/")
}
