package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	accountServices "github.com/viictormg/fribeer-v2/internal/domain/service/account"
	accountAdapters "github.com/viictormg/fribeer-v2/internal/infrastructure/adapters/database/account"

	productUsecases "github.com/viictormg/fribeer-v2/internal/application/usecase/product-service"
	productServices "github.com/viictormg/fribeer-v2/internal/domain/service/product-service"
	productAdapters "github.com/viictormg/fribeer-v2/internal/infrastructure/adapters/database/product-service"
	productHandlers "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/product"

	measureUnitUsecases "github.com/viictormg/fribeer-v2/internal/application/usecase/measure_unit"
	measureUnitServices "github.com/viictormg/fribeer-v2/internal/domain/service/measure_unit"
	measureUnitAdapters "github.com/viictormg/fribeer-v2/internal/infrastructure/adapters/database/measure_unit"
	measureUnitHandlers "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/measure_unit"

	unitTimeUsecases "github.com/viictormg/fribeer-v2/internal/application/usecase/unit_time"
	unitTimeServices "github.com/viictormg/fribeer-v2/internal/domain/service/unit_time"
	unitTimeAdapters "github.com/viictormg/fribeer-v2/internal/infrastructure/adapters/database/unit_time"
	unitTimeHandlers "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/unit_time"

	typeDocumentUsecases "github.com/viictormg/fribeer-v2/internal/application/usecase/type_document"
	typeDocumentServices "github.com/viictormg/fribeer-v2/internal/domain/service/type_document"
	typeDocumentAdapters "github.com/viictormg/fribeer-v2/internal/infrastructure/adapters/database/type_document"
	typeDocumentHandlers "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/type_document"

	customerUsecases "github.com/viictormg/fribeer-v2/internal/application/usecase/customer"
	peopleServices "github.com/viictormg/fribeer-v2/internal/domain/service/people"
	peopleAdapters "github.com/viictormg/fribeer-v2/internal/infrastructure/adapters/database/people"
	customerHandlers "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/customer"

	saleUsecases "github.com/viictormg/fribeer-v2/internal/application/usecase/sale"
	saleServices "github.com/viictormg/fribeer-v2/internal/domain/service/sale"
	campusAdapters "github.com/viictormg/fribeer-v2/internal/infrastructure/adapters/database/campus"
	saleAdapters "github.com/viictormg/fribeer-v2/internal/infrastructure/adapters/database/sale"

	saleHandlers "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/sale"

	AuthUsecases "github.com/viictormg/fribeer-v2/internal/application/usecase/auth"
	authHandlers "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/auth"

	campusServices "github.com/viictormg/fribeer-v2/internal/domain/service/campus"

	database "github.com/viictormg/fribeer-v2/internal/infrastructure/pkg/database"
	"github.com/viictormg/fribeer-v2/internal/infrastructure/server"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	db, _ := database.InitConnectionDB()

	accountAdapter := accountAdapters.NewAccountAdapter(db)

	accountService := accountServices.NewAccountService(accountAdapter)

	productAdapter := productAdapters.NewProductAdapter(db)
	productService := productServices.NewProductServie(productAdapter)
	productUsecase := productUsecases.NewProductUsecase(productService)
	productHandler := productHandlers.NewProductHandler(productUsecase)

	measureUnitAdapter := measureUnitAdapters.NewMeasureAdaterUnit(db)
	measureUnitService := measureUnitServices.NewMeasureUnitService(measureUnitAdapter)
	measureUnitUsecase := measureUnitUsecases.NewMeasureUnitUsecase(measureUnitService)
	measureUnitHandler := measureUnitHandlers.NewMeasuerUnitHandler(measureUnitUsecase)

	unitTimeAdapter := unitTimeAdapters.NewUnitTimeAdapter(db)
	unitTimeService := unitTimeServices.NewUnitTimeService(unitTimeAdapter)
	unitTimeUsecase := unitTimeUsecases.NewUnitTimeUseCase(unitTimeService)
	unitTimeHandler := unitTimeHandlers.NewUnitTimeHandler(unitTimeUsecase)

	typeDocumentAdapter := typeDocumentAdapters.NewTypeDocumentAdapter(db)
	typeDocumentService := typeDocumentServices.NewTypeDocumentService(typeDocumentAdapter)
	typeDocumentUsecase := typeDocumentUsecases.NewTypeDocumentUseCase(typeDocumentService)
	typeDocumentHandler := typeDocumentHandlers.NewTypeDocumentHandler(typeDocumentUsecase)

	peopleAdapter := peopleAdapters.NewPeopleAdapter(db)
	peopleService := peopleServices.NewPeopleService(peopleAdapter)
	customerUsecase := customerUsecases.NewCustomerUsecase(peopleService)
	customerHandler := customerHandlers.NewCustomerHandler(customerUsecase)

	saleAdapter := saleAdapters.NewSaleAdapter(db)
	saleService := saleServices.NewSaleService(saleAdapter, productAdapter)
	saleUsecase := saleUsecases.NewSaleUsecase(saleService)
	saleHandler := saleHandlers.NewSaleHandler(saleUsecase)

	campusAdapter := campusAdapters.NewCampusAdapter(db)
	campuseService := campusServices.NewCampusService(campusAdapter)

	AuthUsecase := AuthUsecases.NewAuthUsecase(accountService, campuseService)
	authHandler := authHandlers.NewAuthHandler(AuthUsecase)

	srv := server.NewServer(
		port,
		*productHandler,
		*measureUnitHandler,
		*unitTimeHandler,
		*typeDocumentHandler,
		*customerHandler,
		*saleHandler,
		*authHandler,
	)

	srv.RunServer()
}
