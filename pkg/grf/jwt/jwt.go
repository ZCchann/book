package jwt

import (
	"book/initalize/database/redis"
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type CustomClaims struct {
	AccessID string
	Username string
	jwt.StandardClaims
}

type internalJwt struct {
	SigningKey []byte
}

// err
var (
	errTokenExpired     = errors.New("token is expired")
	errTokenNotValidYet = errors.New("token not active yet")
	errTokenMalformed   = errors.New("that's not even a token")
	errTokenInvalid     = errors.New("couldn't handle this token")
)

var jwtSecret string

func Init(key string) {
	jwtSecret = key
}

func JWT() *internalJwt {
	return &internalJwt{
		[]byte(jwtSecret),
	}
}
func (j *internalJwt) CreateToken(claims CustomClaims) (string, error) {
	if claims.Id == "" {
		claims.Id = uuid.New().String()
	}
	redis.Redis().Set(context.Background(), "jwt:"+claims.Id, time.Now().String(), 7*24*time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}
func (j *internalJwt) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errTokenNotValidYet
			} else {
				return nil, errTokenInvalid
			}
		}
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errTokenInvalid
}
func (j *internalJwt) RefreshToken(tokenString string) (string, error) {
	var expires = 6 * time.Hour
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(expires).Unix()
		return j.CreateToken(*claims)
	}
	return "", errTokenInvalid
}
