package keystore

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"io/fs"
	"path"
	"strings"
	"sync"
)

type KeyStore struct {
	mu    sync.RWMutex
	store map[string]*rsa.PrivateKey
}

func New() *KeyStore {
	return &KeyStore{
		store: make(map[string]*rsa.PrivateKey),
	}
}

func NewMap(store map[string]*rsa.PrivateKey) *KeyStore {
	return &KeyStore{
		store: store,
	}
}

func NewFS(fileSystem fs.FS) (*KeyStore, error) {
	k := New()

	// Function for WalkDir, it walks through for given path in WalkDir function. It is called each of them.
	f := func(filePath string, dirEntry fs.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if dirEntry.IsDir() {
			return nil
		}

		if path.Ext(filePath) != ".pem" {
			return nil
		}

		file, err := fileSystem.Open(filePath)
		if err != nil {
			return err
		}

		defer file.Close()

		privatePem, err := io.ReadAll(io.LimitReader(file, 1024*1024))
		if err != nil {
			return err
		}

		privateKeyFromPEM, err := jwt.ParseRSAPrivateKeyFromPEM(privatePem)
		if err != nil {
			return err
		}

		k.store[strings.TrimSuffix(dirEntry.Name(), ".pem")] = privateKeyFromPEM

		return nil

	}

	if err := fs.WalkDir(fileSystem, ".", f); err != nil {
		return nil, fmt.Errorf("walking fs err : %w", err)
	}

	return k, nil
}

func (k *KeyStore) Add(kid string, key *rsa.PrivateKey) {
	k.mu.Lock()
	defer k.mu.Unlock()

	k.store[kid] = key
}

func (k *KeyStore) Remove(kid string) {
	k.mu.RLock()
	defer k.mu.RUnlock()

	delete(k.store, kid)
}

// retrievers for Public and Private Keys.

func (k *KeyStore) PrivateKey(kid string) (*rsa.PrivateKey, error) {
	k.mu.RLock()
	defer k.mu.RUnlock()

	key, found := k.store[kid]
	if !found {
		return nil, fmt.Errorf("private key not found for %s", kid)
	}

	return key, nil
}

func (k *KeyStore) PublicKey(kid string) (*rsa.PublicKey, error) {
	private, err := k.PrivateKey(kid)
	if err != nil {
		return nil, fmt.Errorf("public key not found for %s", kid)
	}

	return &private.PublicKey, nil
}

func (k *KeyStore) InMemKeyFunc(token *jwt.Token) (interface{}, error) {
	kid, ok := token.Header["kid"]
	if !ok {
		return nil, errors.New("kid is not provided")
	}

	return &k.store[kid.(string)].PublicKey, nil
}
