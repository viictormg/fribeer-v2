package server

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	measureHandlers "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/measure_unit"
	productHandlers "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/product"
	unitTimeHandlers "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/unit_time"
)

type Server struct {
	Port               string
	ProductHandler     productHandlers.ProductHandler
	MeasuerUnitHandler measureHandlers.MeasuerUnitHandler
	UnitTimeHandler    unitTimeHandlers.UnitTimeHandler
}

func NewServer(
	port string,
	productHandler productHandlers.ProductHandler,
	MeasuerUnitHandler measureHandlers.MeasuerUnitHandler,
	UnitTimeHandler unitTimeHandlers.UnitTimeHandler,
) *Server {
	return &Server{
		Port:               port,
		ProductHandler:     productHandler,
		MeasuerUnitHandler: MeasuerUnitHandler,
		UnitTimeHandler:    UnitTimeHandler,
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
	apiPulic.GET("/product", s.ProductHandler.GetProductsHandler)

	apiPulic.POST("/service", s.ProductHandler.CreateServiceHandler)
	apiPulic.GET("/measureUnit", s.MeasuerUnitHandler.GetMeasureUnitHandler)
	apiPulic.GET("/unitTime", s.UnitTimeHandler.GetUnitTimeHandler)

	err := e.Start(":" + s.Port)
	if err != nil {
		log.Fatal(err)
	}
}
