package middleware

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/holycann/smart-parking-backend/config"
)

func CreateJWT(secret []byte, UserID int) (string, error) {
    expiration := time.Second * time.Duration(config.Env.JWTExpirationInSecond);

    token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
        "user_id": strconv.Itoa(UserID),
        "expired_at" : time.Now().Add(expiration).Unix(),
    })

    tokenString, err := token.SignedString(secret)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}