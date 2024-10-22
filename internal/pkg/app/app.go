package app

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"go.mod/internal/app/endpoint"
	"go.mod/internal/app/mw"
	"go.mod/internal/app/service"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()
	a.e = endpoint.New(a.s)

	a.echo = echo.New()
	a.echo.Use(mw.RoleCheck)
	a.echo.GET("/status", a.e.Status)
	return a, nil
}

func (a *App) Run() error {
	fmt.Println("server starting...")
	err := a.echo.Start(":8080")
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
