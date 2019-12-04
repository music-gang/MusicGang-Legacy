package routes

import (
	"github.com/labstack/echo"
)

// InitGetRoutes - Dichiara tutte le route GET
func InitGetRoutes(e *echo.Echo) {
}

// InitPostRoutes - Dichiara tutte le route POST
func InitPostRoutes(e *echo.Echo) {
}

// InitStaticRoutes - inizializza le route per i file statici
func InitStaticRoutes(e *echo.Echo) {
	e.Static("/", "./dist/")
}
