package main

import (
	"flag"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"log"
)

const (
	staticUrl = "/static"
)

var (
	httpAddr = flag.String("http", defaultAddr, "Listen for HTTP connections on this address.")
)

func init() {
	flag.Parse()
	if err := cacheTemplates([]string{
		"templates/home.tmpl",
	}); err != nil {
		log.Fatal(err)
	}

}

func app() *echo.Echo {
	e := echo.New()

	// Default Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())

	// Groups
	h := e.Group("")
	s := e.Group(staticUrl)

	h.Get("/", home())

	// Static
	s.Get("/:folder/:file", serveFile())
	s.Get("/favicon.ico", serveFavicon())

	return e
}

func main() {
	e := app()
	// Start server
	e.Run(standard.New(*httpAddr))
}
