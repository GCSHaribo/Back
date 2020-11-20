package main

import (
	"github.com/GCSHaribo/Back/src/rest"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()
	h := rest.NewHandler()
	e.Use(middleware.Logger())
	e.GET("/", h.GetMainPage) //서버 검사
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/signup", h.Signup)
	e.POST("/signin", h.Signin)
	e.Logger.Fatal(e.Start(":1323"))
}
