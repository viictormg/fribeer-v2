package usecase

import (
	"encoding/hex"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	"golang.org/x/crypto/argon2"
)

var hmacSampleSecret []byte

const HASH_MEMORY = 64 * 1024
const HASH_KEY_LENGTH = 36
const HASH_ITERATIONS = 3
const HASH_PARALELLEISM = 1
const TIME_EXPIRED_TOKEN = 1 * time.Hour

func (authUsecase *AuthUsecase) LoginUsecase(user, password string) (string, error) {
	response, err := authUsecase.accountService.GetSignTokenLoginService(user, authUsecase.hashPassword(password))

	if err != nil {
		logrus.Error(err)
		return "", err
	}

	return authUsecase.signToken(response), nil
}

func (authUsecase *AuthUsecase) hashPassword(password string) string {
	hash := argon2.IDKey([]byte(password), []byte(os.Getenv("SECRET_PASSWORD")), HASH_ITERATIONS, HASH_MEMORY, HASH_PARALELLEISM, HASH_KEY_LENGTH)
	return hex.EncodeToString(hash)
}

func (authUsecase *AuthUsecase) signToken(signToken dto.SignTokenDTO) string {
	// expireTime := time.Now().Add(TIME_EXPIRED_TOKEN)

	claims := dto.CustomClaims{
		SignTokenDTO: signToken,
		RegisteredClaims: jwt.RegisteredClaims{
			// ExpiresAt: jwt.NewNumericDate(time.Unix(expireTime.Unix(), 0)),
			Issuer: "system",
		},
	}

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// tokenString, _ := token.SignedString([]byte("todoesculpadelabareta"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET_PASSWORD")))

	return tokenString
}
