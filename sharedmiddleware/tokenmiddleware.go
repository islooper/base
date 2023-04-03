package sharedmiddleware

import (
	"context"
	"net/http"
)

type TokensMiddleware struct {
}

func NewTokenMiddleware() *TokensMiddleware {
	return &TokensMiddleware{}
}

func (m *TokensMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(context.WithValue(r.Context(), "token", r.Header.Get("token")))
		next(w, r)
	}
}
