// Работа с JWT в контексте
// Условие:
//
// Напиши две функции:
//
// 1. `AddJWTToContext(ctx context.Context, userID int) (context.Context, error)`
// 2. `ExtractUserIDFromContext(ctx context.Context) (int, error)`
//
// `AddJWTToContext` должен:
//
// - Создавать JWT-токен (используй `github.com/golang-jwt/jwt/v5`).
// - Зашифровывать в него `userID`.
// - Возвращать новый `context.Context`, в который записан JWT-токен.
//
// `ExtractUserIDFromContext` должен:
//
// - Извлекать JWT-токен из контекста.
// - Расшифровывать `userID`.
// - Вывести его на экран
//
// Дополнительное условие:
//
// Создай горутину, в которой будет использоваться `ExtractUserIDFromContext`. Покажи, что передача контекста работает и данные можно безопасно извлекать между горутинами.

package main

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

var hmacSampleSecret = []byte("my_secret_key")

type contextKey string

const jwtContextKey contextKey = "jwt"

func AddJWTToContext(ctx context.Context, userID int) (context.Context, error) {
	claims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		return ctx, err
	}

	ctx = context.WithValue(ctx, jwtContextKey, tokenString)
	return ctx, nil
}

func ExtractUserIDFromContext(ctx context.Context) (int, error) {
	val := ctx.Value(jwtContextKey)
	tokenString, ok := val.(string)
	if !ok {
		return 0, fmt.Errorf("token not found in context or wrong type")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if idFloat, ok := claims["id"].(float64); ok {
			return int(idFloat), nil
		}
		return 0, fmt.Errorf("user id not found in token")
	}
	return 0, fmt.Errorf("invalid token claims")
}

func main() {
	ctx := context.Background()
	ctx, err := AddJWTToContext(ctx, 42)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go func(ctx context.Context) {
		userID, err := ExtractUserIDFromContext(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("User ID from context:", userID)
		close(done)
	}(ctx)

	<-done
}
