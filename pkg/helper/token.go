package helper

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/cast"
)

type TokenInfo struct {
	UserID     string `json:"user_id"`
	ClientType string `json:"client_type"`
}

func GenerateJWT(m map[string]interface{}, tokenExpireTime time.Duration, tokenSecretKey string) (tokenString string, err error) {
	var token = jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	for key, value := range m {
		claims[key] = value
	}

	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(tokenExpireTime).Unix()

	tokenString, err = token.SignedString([]byte(tokenSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseClaims(token string, secretKey string) (result TokenInfo, err error) {
	var claims jwt.MapClaims

	claims, err = ExtractClaims(token, secretKey)
	if err != nil {
		return result, err
	}

	result.UserID = cast.ToString(claims["user_id"])
	result.ClientType = cast.ToString(claims["client_type "])
	if len(result.UserID) <= 0 {
		err = errors.New("cannot parse fild 'UserID'")
		return result, err
	}

	return
}

func ExtractClaims(tokenString string, secretKey string) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)

	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func ExtractToken(bearer string) (token string, err error) {
	strArr := strings.Split(bearer, " ")

	if len(strArr) == 2 {
		return strArr[1], nil
	}

	return token, errors.New("wrong token format")
}
