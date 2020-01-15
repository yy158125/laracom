package service

import (
	"github.com/dgrijalva/jwt-go"
	pb "github.com/yy158125/laracom/user-service/proto/user"
	"github.com/yy158125/laracom/user-service/repo"
	"time"
)

var (
	secret = []byte("laracomUserTokenKeySecret")
)

type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}


type Authable interface {
	Parse(token string) (*CustomClaims, error)
	Sign(user *pb.User) (string, error)
}
type TokenService struct {
	//Repo repo.Repository
}

func (t *TokenService) Parse(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString,&CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return secret, nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (t *TokenService) Sign(user *pb.User) (string,error) {
	expireToken := time.Now().Add(time.Hour * 72).Unix()
	// Create the Claims
	claims := CustomClaims{
		User:           user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "laracom.user.service",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(secret)
}
