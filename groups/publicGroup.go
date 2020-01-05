package groups

import (
	"pkkkk/handlers"

	"github.com/labstack/echo"
)

// DefaultGroup default group
func DefaultGroup(e *echo.Echo) {
	g := e.Group("")

	// config info
	g.GET("/start", handlers.Start)
	g.GET("/start/:idprovinsi", handlers.InProvinsi)
	g.GET("/start/:idprovinsi/:idkota", handlers.InKota)
	g.GET("/start/:idprovinsi/:idkota/:idkecamatan", handlers.InKecamatan)
	g.GET("/start/:idprovinsi/:idkota/:idkecamatan/:idkelurahan", handlers.InKelurahan)
}
