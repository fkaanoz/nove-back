package web

import (
	"errors"
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
	ApiToken  string
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
	token.Header["kid"] = a.ActiveKid

	key, err := a.KeyStore.PrivateKey(a.ActiveKid)
	if err != nil {
		return "", err
	}

	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (a *Auth) ValidateToken(token string) error {
	parsedToken, err := jwt.Parse(token, jwt.Keyfunc(a.KeyFunc))
	if err != nil {
		return err

	}

	if !parsedToken.Valid {
		return errors.New("token is not valid")
	}

	return nil
}
