package handler

import (
	"simple-CRUD/pkg/app"
	"simple-CRUD/pkg/entity"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateJWT(r *entity.GenerateTokenRequest) (string, error) {
	appConfig := app.GetConfig().App
	now := time.Now()
	expTime := now.Add(time.Duration(appConfig.JWT_exp_minutes) * time.Minute)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":     strconv.Itoa(r.Id),
		"iss":     "simple-CRUD",
		"exp":     expTime.Unix(),
		"iat":     now.Unix(),
		"isAdmin": r.IsAdmin,
		"id":      r.Id,
	})

	tokenString, err := claims.SignedString([]byte(appConfig.JWT_secret))
	return tokenString, err
}
