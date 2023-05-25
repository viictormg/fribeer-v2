package api

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (cron *CronJobHandler) ServiceCardJob(c echo.Context) error {
	response := cron.serviceCardUsecase.JobServiceCardUsecase()

	logrus.Error(response)
	return response
}
