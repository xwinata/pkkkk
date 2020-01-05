package router

import (
	"pkkkk/groups"

	"github.com/labstack/echo"
)

// NewRouter define new router
func NewRouter() *echo.Echo {
	e := echo.New()

	groups.DefaultGroup(e)

	return e
}
