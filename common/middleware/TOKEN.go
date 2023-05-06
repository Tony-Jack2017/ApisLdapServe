package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type ApisClaims struct {
	Account string `json:"account"`
	AdminID string `json:"adminId"`
	jwt.RegisteredClaims
}

func GenerateToken(account string, accountId string) (string, error) {
	claim := ApisClaims{
		Account: account,
		AdminID: accountId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Issuer:    "ApisLdapServe",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
	return token.SignedString([]byte("secret-key"))
}

func VerifyToken(tokenString string) (*ApisClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &ApisClaims{}, func(token *jwt.Token) (interface{}, error) {
		return "", nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*ApisClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func TOKENMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.Request.Header.Get("access_auth_token")
		if _, err := VerifyToken(tokenString); err != nil {
			fmt.Println(err)
			context.JSON(401, err)
		}
		context.Next()
	}
}
