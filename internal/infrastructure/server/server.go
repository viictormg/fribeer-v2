package server

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	authHandlers "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/auth"
	serviceCardJobs "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/cronjob"
	customerHandlers "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/customer"
	measureHandlers "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/measure_unit"
	PeopleHandlers "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/people"
	productHandlers "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/product"
	saleHandlers "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/sale"
	serviceCards "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/service_card"
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
	authHandler         authHandlers.AuthHandler
	ServiceCard         serviceCards.ServiceCardHandler
	serviceCardJob      serviceCardJobs.CronJobHandler
	peopleHandlers      PeopleHandlers.PeopleHandler
}

func NewServer(
	port string,
	productHandler productHandlers.ProductHandler,
	MeasuerUnitHandler measureHandlers.MeasuerUnitHandler,
	UnitTimeHandler unitTimeHandlers.UnitTimeHandler,
	TypeDocumentHandler typeDocumentHandlers.TypeDocumentHandler,
	CustomerHandlers customerHandlers.CustomerHandler,
	SaleHandlers saleHandlers.SaleHandler,
	authHandler authHandlers.AuthHandler,
	ServiceCard serviceCards.ServiceCardHandler,
	serviceCardJob serviceCardJobs.CronJobHandler,
	peopleHandlers PeopleHandlers.PeopleHandler,

) *Server {
	return &Server{
		Port:                port,
		ProductHandler:      productHandler,
		MeasuerUnitHandler:  MeasuerUnitHandler,
		UnitTimeHandler:     UnitTimeHandler,
		TypeDocumentHandler: TypeDocumentHandler,
		CustomerHandlers:    CustomerHandlers,
		SaleHandlers:        SaleHandlers,
		authHandler:         authHandler,
		ServiceCard:         ServiceCard,
		serviceCardJob:      serviceCardJob,
		peopleHandlers:      peopleHandlers,
	}
}

func (s *Server) RunServer() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	apiPulic := e.Group("/api")
	apiPrivate := e.Group("/api")

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(dto.CustomClaims)
		},
		SigningKey: []byte(os.Getenv("SECRET_PASSWORD")),
	}

	// e.Use(echojwt.WithConfig(config))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	apiPrivate.Use(echojwt.WithConfig(config))
	apiPrivate.GET("/test-auth", s.authHandler.TestJWT)

	apiPrivate.GET("/getUser", s.authHandler.GetUserHandler)

	apiPrivate.POST("/product", s.ProductHandler.CreateProductHandler)
	apiPrivate.GET("/product", s.ProductHandler.GetProductsHandler)

	apiPulic.POST("/service", s.ProductHandler.CreateServiceHandler)
	apiPulic.GET("/measureUnit", s.MeasuerUnitHandler.GetMeasureUnitHandler)
	apiPulic.GET("/unitTime", s.UnitTimeHandler.GetUnitTimeHandler)
	apiPulic.GET("/typeDocument", s.TypeDocumentHandler.GetTypeDocumentHandler)

	apiPulic.POST("/customer", s.CustomerHandlers.CreateCustomerHandler)
	apiPulic.GET("/customer", s.CustomerHandlers.GetCustomerHandler)

	apiPulic.POST("/sale", s.SaleHandlers.CreateSaleHandler)
	apiPulic.GET("/sale", s.SaleHandlers.GetSalesHandler)
	apiPulic.POST("/login", s.authHandler.LoginHandler)

	apiPrivate.GET("/serviceCard", s.ServiceCard.GetServiceCardHandler)
	apiPrivate.GET("/serviceCard/:id", s.ServiceCard.GetServiceCardByIDHandler)

	apiPrivate.GET("/serviceCardJob", s.serviceCardJob.ServiceCardJob)

	apiPrivate.GET("/person/:id", s.peopleHandlers.GetPersonByIDHandler)

	err := e.Start(":" + s.Port)
	if err != nil {
		log.Fatal(err)
	}
}
