package usecase

import (
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTManager struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	ttl        time.Duration
}

func NewJWTManager(privateKey string, publicKey string, ttl time.Duration) (*JWTManager, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode private key: %w", err)
	}

	rsaPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse rsa private key from pem: %w", err)
	}

	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode public key: %w", err)
	}

	rsaPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse rsa public key from pem: %w", err)
	}

	return &JWTManager{
		privateKey: rsaPrivateKey,
		publicKey:  rsaPublicKey,
		ttl:        ttl,
	}, nil
}

func (manager *JWTManager) CreateToken(payload interface{}) (string, error) {
	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["sub"] = payload
	claims["exp"] = now.Add(manager.ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(manager.privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to create signed jwt token: %w", err)
	}

	return token, nil
}

func (manager *JWTManager) ValidateToken(token string) (interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method of jwt token: %s", t.Header["alg"])
		}

		return manager.publicKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse jwt token: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("invalid jwt token")
	}

	return claims["sub"], nil
}
