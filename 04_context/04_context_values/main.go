package main

import (
	"context"
	"log"
	"net/http"
)

// ???

type requestIDKey struct {
}

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := "123"
		ctx := context.WithValue(r.Context(), requestIDKey{}, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func processOrder(ctx context.Context, orderID string) error {
	requestID, _ := ctx.Value(requestIDKey{}).(string)
	log.Printf("[%s] Processing order %s", requestID, orderID)

	// Передаем контекст дальше
	// return db.QueryContext(ctx, "INSERT INTO orders ...")
	return nil
}

// + tx example
