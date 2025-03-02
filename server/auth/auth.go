package auth

import (
	"context"
	"fmt"
	"log"
	"main/database"
	cruds "main/server/crud"
	"os"
	"time"

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

func BuildAccessToken(sub uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24 * 365).Unix(), // 365日
	})

	accessToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		log.Printf("error: %v\n", err)
		return "", err
	}

	return accessToken, nil
}

func Auth(ctx context.Context) *database.User {
	db := database.GetDB()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("error: FromIncomingContext")
		return nil
	}

	tokenString := GetTokenFromMD(md)
	if tokenString == "" {
		log.Println("error: GetTokenFromMD")
		return nil
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		log.Printf("error: %v\n", err)
		return nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userId, ok := claims["sub"].(float64)
		if !ok {
			log.Println("error: failed to get userId")
			return nil
		}

		user := cruds.GetUserById(db, int(userId))
		return user
	} else {
		return nil
	}
}
