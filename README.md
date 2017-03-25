# Go Clacks Middleware

Middleware to support the GNU Clacks in various Go web frameworks.

*Mostly an educational exercise to try to write a quick middleware for multiple go frameworks*

## Usage

```go
package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/jonathanmh/goClacks/echo"
)

func main() {
	e := echo.New()
	e.Use(goClacks.Terrify)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
```

## Currently Contains

* [buffalo][]
* [echo][]
* [gin][]
* [net/http][]

## Educational Findings

They're all pretty similar, but at least now there will be more Clacks throughout the net.

[buffalo]: http://gobuffalo.io/
[echo]: https://echo.labstack.com/
[gin]: https://gin-gonic.github.io/gin/
[net/http]: https://golang.org/pkg/net/http/
