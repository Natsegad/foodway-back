package jwt

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type SClaims struct {
	jwt.StandardClaims
	Id uint32 `json:"id"`
}

// Это значение должно браться из cfg
var secret string = "secret"

func GenerateJwtById(id uint32) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &SClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Id: id,
	})

	tkStr, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Printf("Error SignedString jwt %s !\n", err.Error())
		panic(nil)
	}

	return tkStr
}

func GenerateRefeshTokenById(id uint32) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &SClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(72 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Id: id,
	})

	tkStr, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Printf("Error SignedString jwt %s !\n", err.Error())
		panic(nil)
	}

	return tkStr
}

func GetJwt(tok string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tok, &SClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	return token, err
}

func TokenExpired(token string) (bool, error) {
	_, err := GetJwt(token)
	if err != nil {
		vErr, ok := err.(*jwt.ValidationError)
		if ok {
			if vErr.Errors == 16 {
				return true, nil
			}
		}
		return false, err
	}
	return false, nil
}

func ParseJwt(token *jwt.Token) (*SClaims, error) {
	if claims, ok := token.Claims.(*SClaims); ok && token.Valid {
		return claims, nil
	}

	fmt.Println("Error parse token !")
	return nil, nil
}
