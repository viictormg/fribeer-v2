package server

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	customerHandlers "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/customer"
	measureHandlers "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/measure_unit"
	productHandlers "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/product"
	saleHandlers "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/sale"
	typeDocumentHandlers "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/type_document"
	unitTimeHandlers "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/unit_time"
)

type Server struct {
	Port                string
	ProductHandler      productHandlers.ProductHandler
	MeasuerUnitHandler  measureHandlers.MeasuerUnitHandler
	UnitTimeHandler     unitTimeHandlers.UnitTimeHandler
	TypeDocumentHandler typeDocumentHandlers.TypeDocumentHandler
	CustomerHandlers    customerHandlers.CustomerHandler
	SaleHandlers        saleHandlers.SaleHandler
}

func NewServer(
	port string,
	productHandler productHandlers.ProductHandler,
	MeasuerUnitHandler measureHandlers.MeasuerUnitHandler,
	UnitTimeHandler unitTimeHandlers.UnitTimeHandler,
	TypeDocumentHandler typeDocumentHandlers.TypeDocumentHandler,
	CustomerHandlers customerHandlers.CustomerHandler,
	SaleHandlers saleHandlers.SaleHandler,

) *Server {
	return &Server{
		Port:                port,
		ProductHandler:      productHandler,
		MeasuerUnitHandler:  MeasuerUnitHandler,
		UnitTimeHandler:     UnitTimeHandler,
		TypeDocumentHandler: TypeDocumentHandler,
		CustomerHandlers:    CustomerHandlers,
		SaleHandlers:        SaleHandlers,
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
	apiPulic.GET("/typeDocument", s.TypeDocumentHandler.GetTypeDocumentHandler)

	apiPulic.POST("/customer", s.CustomerHandlers.CreateCustomerHandler)
	apiPulic.GET("/customer", s.CustomerHandlers.GetCustomerHandler)

	apiPulic.POST("/sale", s.SaleHandlers.CreateSaleHandler)

	err := e.Start(":" + s.Port)
	if err != nil {
		log.Fatal(err)
	}
}
