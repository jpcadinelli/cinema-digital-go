package security

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"os"
	"time"
)

type JWTService struct {
	secretKey string
	issure    string
}

func NewJWTService() *JWTService {
	return &JWTService{
		secretKey: os.Getenv("JWT_SECRET_KEY"),
		issure:    "api-pattern-go",
	}
}

type Claims struct {
	Sum uuid.UUID `json:"sum"`
	jwt.StandardClaims
}

func (s *JWTService) GenerateToken(id uuid.UUID) (string, error) {
	claims := &Claims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *JWTService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token %v", token)
		}

		return []byte(s.secretKey), nil
	})

	return err == nil
}

func (s *JWTService) GetUserId(token string) (uuid.UUID, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid signing method")
		}

		return []byte(s.secretKey), nil
	})

	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to parse token: %v", err)
	}

	if claims, ok := parsedToken.Claims.(*Claims); ok && parsedToken.Valid {
		return claims.Sum, nil
	}

	return uuid.Nil, fmt.Errorf("invalid token or claims")
}
