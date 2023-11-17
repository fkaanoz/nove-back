package web

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"shtil/business/keystore"
	"time"
)

type KeyFunc func(token *jwt.Token) (interface{}, error)

type Auth struct {
	KeyFolder string
	ActiveKid string
	KeyStore  *keystore.KeyStore
	KeyFunc   KeyFunc
}

func (a *Auth) GenerateToken(hour int) (string, error) {
	claims := struct {
		jwt.RegisteredClaims
	}{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(hour))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["activeKid"] = a.ActiveKid

	key, err := a.KeyStore.PrivateKey(a.ActiveKid)
	if err != nil {
		return "", err
	}

	fmt.Println("3")

	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (a *Auth) ValidateToken(token string) bool {
	parsedToken, err := jwt.Parse(token, jwt.Keyfunc(a.KeyFunc))
	if err != nil {
		fmt.Print("parser error", err)
		return false
	}

	return parsedToken.Valid
}
