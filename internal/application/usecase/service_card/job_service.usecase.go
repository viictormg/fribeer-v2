package usecase

import "fmt"

func (s *ServiceCardUsecase) JobServiceCardUsecase() error {
	serviceCards, err := s.serviceCerdService.GetAllServiceCardService()

	// for _, service := range serviceCards {

	// }

	fmt.Println(serviceCards)
	return err
}
