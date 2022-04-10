package jwt

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateJWT(username string) (string, error)
	DecodeJWT(tokenString string) (*TokenClaim, bool, error)
}

type TokenClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey          string
	expireTimeInMinute int
}

func (self *jwtService) GenerateJWT(username string) (string, error) {

	jwtKey := []byte(self.secretKey)

	claims := TokenClaim{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(self.expireTimeInMinute) * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func (self *jwtService) DecodeJWT(tokenString string) (*TokenClaim, bool, error) {

	var secret = []byte(self.secretKey)

	token, err := jwt.ParseWithClaims(tokenString, &TokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		if strings.Contains(err.Error(), "expired") {
			return nil, false, nil
		}
		return nil, false, err
	}

	if claims, ok := token.Claims.(*TokenClaim); ok {
		return claims, token.Valid, nil
	}

	return nil, false, err
}

func NewJWTService(secretKey string, expireTimeInMinute int) JWTService {
	return &jwtService{secretKey: secretKey, expireTimeInMinute: expireTimeInMinute}
}
