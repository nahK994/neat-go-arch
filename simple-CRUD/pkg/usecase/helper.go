package usecase

import (
	"simple-CRUD/pkg/app"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateJWT(isAdmin bool, id int) (string, error) {
	appConfig := app.GetConfig().App
	now := time.Now()
	expTime := now.Add(time.Duration(appConfig.JWT_exp_minutes) * time.Minute)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":     strconv.Itoa(id),
		"iss":     "simple-CRUD",
		"exp":     expTime.Unix(),
		"iat":     now.Unix(),
		"isAdmin": isAdmin,
		"id":      id,
	})

	tokenString, err := claims.SignedString([]byte(appConfig.JWT_secret))
	return tokenString, err
}
