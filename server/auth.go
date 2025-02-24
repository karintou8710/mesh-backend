package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/metadata"
)

// TODO: 後で環境変数に置き換える
var SECRET = []byte("secret")

func GetTokenFromMD(md metadata.MD) string {
	tokenStringList := md.Get("token")
	if len(tokenStringList) > 0 {
		return tokenStringList[0]
	} else {
		return ""
	}
}

func BuildAccessToken(sub uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24 * 365).Unix(), // 365日
	})

	accessToken, err := token.SignedString(SECRET)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return "", err
	}

	return accessToken, nil
}

func Auth(ctx context.Context) jwt.MapClaims {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}

	tokenString := GetTokenFromMD(md)
	if tokenString == "" {
		return nil
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return SECRET, nil
	})

	if err != nil {
		log.Println(err)
		return nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims
	} else {
		return nil
	}
}
