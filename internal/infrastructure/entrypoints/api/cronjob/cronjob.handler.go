package api

import "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/port"

type CronJobHandler struct {
	serviceCardUsecase port.IServiceCardUsecase
}

func NewCronJobHandler(serviceCardUsecase port.IServiceCardUsecase) *CronJobHandler {
	return &CronJobHandler{
		serviceCardUsecase: serviceCardUsecase,
	}
}
