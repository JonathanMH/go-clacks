package goClacks

import "github.com/gobuffalo/buffalo"

func Terrify(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		c.Response().Header().Add("X-Clacks-Overhead", "GNU Terry Pratchett")
		err := next(c)
		return err
	}
}
