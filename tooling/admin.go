package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"os"
	"time"
)

func main() {

	//if err := GenerateKeys("test"); err != nil {
	//	fmt.Println("generate key error")
	//	os.Exit(1)
	//}

	token, err := GenerateToken()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(token)
	//
	//fmt.Println("token is sending to parser")
	//
	//err = ParseToken(token)
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

}

func GenerateKeys(fileName string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return err
	}

	blk := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	err = pem.Encode(file, &blk)
	if err != nil {
		return err
	}

	publicBlk := pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
	}

	pfile, err := os.Create(fileName + ".pub")
	if err != nil {
		return err
	}

	err = pem.Encode(pfile, &publicBlk)

	return err
}

func GenerateToken() (string, error) {
	claim := struct {
		jwt.RegisteredClaims
		CustomClaim string
	}{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "",
			Audience:  nil,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 120)),
			NotBefore: nil,
			IssuedAt:  nil,
			ID:        "",
		},
		CustomClaim: "custom-claim",
	}

	privatefile, err := os.Open("/Users/fkaanoz/Desktop/dev/nove/back/zarf/keys/test")
	if err != nil {
		return "", err
	}

	privateKey, err := io.ReadAll(privatefile)
	if err != nil {
		return "", err
	}

	parsedPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	token.Header["kid"] = "test"

	signedToken, err := token.SignedString(parsedPrivateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseToken(token string) error {

	var claim struct {
		jwt.RegisteredClaims
		CustomClaim string
	}

	parsed, err := jwt.ParseWithClaims(token, &claim, keyFunc)
	if err != nil {
		return err
	}

	fmt.Println("is valid ? ", parsed.Valid)

	return nil
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	privateFile, err := os.Open("/Users/fkaanoz/Desktop/dev/nove/back/zarf/keys/test")
	if err != nil {
		return nil, err
	}

	privateByte, err := io.ReadAll(privateFile)
	if err != nil {
		return nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateByte)
	if err != nil {
		return nil, err
	}

	return &privateKey.PublicKey, nil
}
