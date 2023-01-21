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

	unitTimeServices "github.com/viictormg/fribeer-v2/internal/domain/service/unit-time"
	unitTimeAdapters "github.com/viictormg/fribeer-v2/internal/infraesctructure/adapters/database/unit-time"

	database "github.com/viictormg/fribeer-v2/internal/infraesctructure/pkg/database"
	"github.com/viictormg/fribeer-v2/internal/server"
)

func main() {
	port := "4000"

	os.Setenv("USER_DB", "user-fribeer-remote")
	os.Setenv("PASSWORD_DB", "2f690df5cc514129881588fee0393")
	os.Setenv("HOST_DB", "131.0.136.53")
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

	// unitTimeAdapter := unitTimeAdapters.NewMeasureAdaterUnit(db)
	// measureUnitService := measureUnitServices.NewMeasureUnitService(unitTimeAdapter)
	// measureUnitUsecase := measureUnitUsecases.NewMeasureUnitUsecase(measureUnitService)
	// measureUnitHandler := measureUnitHandlers.NewMeasuerUnitHandler(measureUnitUsecase)

	srv := server.NewServer(port, *productHandler, *measureUnitHandler)

	srv.RunServer()
}
