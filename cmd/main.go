package main

import (
	"os"

	productUsecases "github.com/viictormg/fribeer-v2/internal/application/usecase/product-service"
	productServices "github.com/viictormg/fribeer-v2/internal/domain/service/product-service"
	productAdapters "github.com/viictormg/fribeer-v2/internal/infraesctructure/adapters/database/product-service"
	productHandlers "github.com/viictormg/fribeer-v2/internal/infraesctructure/entrypoints/api/product"

	measureUnitUsecases "github.com/viictormg/fribeer-v2/internal/application/usecase/measure_unit"
	measureUnitServices "github.com/viictormg/fribeer-v2/internal/domain/service/measure_unit"
	measureUnitAdapters "github.com/viictormg/fribeer-v2/internal/infraesctructure/adapters/database/measure_unit"
	measureUnitHandlers "github.com/viictormg/fribeer-v2/internal/infraesctructure/entrypoints/api/measure_unit"

	unitTimeUsecases "github.com/viictormg/fribeer-v2/internal/application/usecase/unit_time"
	unitTimeServices "github.com/viictormg/fribeer-v2/internal/domain/service/unit_time"
	unitTimeAdapters "github.com/viictormg/fribeer-v2/internal/infraesctructure/adapters/database/unit_time"
	unitTimeHandlers "github.com/viictormg/fribeer-v2/internal/infraesctructure/entrypoints/api/unit_time"

	database "github.com/viictormg/fribeer-v2/internal/infraesctructure/pkg/database"
	"github.com/viictormg/fribeer-v2/internal/server"
)

func main() {
	port := "4000"

	os.Setenv("USER_DB", "root")
	os.Setenv("PASSWORD_DB", "")
	os.Setenv("HOST_DB", "127.0.0.1")
	os.Setenv("NAME_DB", "FribeerDB")

	db, _ := database.InitConnectionDB()

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

	srv := server.NewServer(port, *productHandler, *measureUnitHandler, *unitTimeHandler)

	srv.RunServer()
}
