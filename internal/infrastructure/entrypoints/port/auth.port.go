package port

type IAuthUsecase interface {
	LoginUsecase(user, password string) (string, error)
}
