package goClacks

import "github.com/labstack/echo"

func Terrify(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Add("X-Clacks-Overhead", "GNU Terry Pratchett")
		err := next(c)
		return err
	}
}
