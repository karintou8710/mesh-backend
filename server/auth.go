package main

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/metadata"
)

func GetTokenFromMD(md metadata.MD) string {
	tokenStringList := md.Get("token")
	if len(tokenStringList) > 0 {
		return tokenStringList[0]
	} else {
		return ""
	}
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

		return []byte("secret"), nil
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
