package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	// ====== Здесь добавили лог =====
	logrus.Info("run app")
	// ===============================
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("Hello, %s", c.QueryParam("name")))
	})

	e.Logger.Fatal(e.Start(":8080"))
}
