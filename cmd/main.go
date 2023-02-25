package main

import (
	"os"

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
	saleHandlers "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api/sale"

	database "github.com/viictormg/fribeer-v2/internal/infrastructure/pkg/database"
	"github.com/viictormg/fribeer-v2/internal/infrastructure/server"
)

func main() {
	port := "4000"

	os.Setenv("USER_DB", "root")
	os.Setenv("PASSWORD_DB", "")
	os.Setenv("HOST_DB", "127.0.0.1")
	os.Setenv("NAME_DB", "FribeerDB")

	db, _ := database.InitConnectionDB()

	saleUsecase := saleUsecases.NewSaleUsecase()
	saleHandler := saleHandlers.NewSaleHandler(saleUsecase)

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

	srv := server.NewServer(
		port,
		*productHandler,
		*measureUnitHandler,
		*unitTimeHandler,
		*typeDocumentHandler,
		*customerHandler,
		*saleHandler,
	)

	srv.RunServer()
}
