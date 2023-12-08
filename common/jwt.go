package common

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/trikrama/Depublic/internal/app/user/entity"
	"github.com/trikrama/Depublic/internal/config"
)



type JwtCustomClaims struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func JWTProtected(secretKey string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtCustomClaims)
		},
		SigningKey: []byte(secretKey),
	})
}

func GenerateAccessToken(c context.Context, user *entity.User) (string, error) {
	expiredTime := time.Now().Local().Add(10 * time.Minute)
	claims := JwtCustomClaims{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var cfg *config.Config
	cfg, _ = config.NewConfig(".env")
	encodedToken, err := token.SignedString([]byte(cfg.JWT.SecretKey))

	if err != nil {
		fmt.Println("salah di generate access token fungsi jwt")
		return "", err
	}

	return encodedToken, nil
}

//migrate create -ext sql -dir db/migrations create_table_user
//migrate -database postgres://postgres:trikrama@localhost:5432/depublic?sslmode=disable -path db/migration-golang up
