package server

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	measureHandlers "github.com/viictormg/fribeer-v2/internal/infraesctructure/entrypoints/api/measure_unit"
	productHandlers "github.com/viictormg/fribeer-v2/internal/infraesctructure/entrypoints/api/product"
)

type Server struct {
	Port               string
	ProductHandler     productHandlers.ProductHandler
	MeasuerUnitHandler measureHandlers.MeasuerUnitHandler
}

func NewServer(
	port string,
	productHandler productHandlers.ProductHandler,
	MeasuerUnitHandler measureHandlers.MeasuerUnitHandler,
) *Server {
	return &Server{
		Port:               port,
		ProductHandler:     productHandler,
		MeasuerUnitHandler: MeasuerUnitHandler,
	}
}

func (s *Server) RunServer() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	apiPulic := e.Group("/api")
	apiPulic.POST("/product", s.ProductHandler.CreateProductHandler)
	apiPulic.POST("/service", s.ProductHandler.CreateServiceHandler)
	apiPulic.GET("/measureUnit", s.MeasuerUnitHandler.GetMeasureUnitHandler)

	err := e.Start(":" + s.Port)
	if err != nil {
		log.Fatal(err)
	}
}
