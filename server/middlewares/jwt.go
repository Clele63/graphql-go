package middlewares

import (
	"context"
	"net/http"
	"strings"

	"workbench/graphql-app/graph/model"
	"workbench/graphql-app/utils"
)

type contextKey struct {
	email string
}

var AuthCtxKey = &contextKey{"auth"}

func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqToken := r.Header.Get("Authorization")
			if reqToken == "" {
				next.ServeHTTP(w, r)
				return
			}

			splitToken := strings.Split(reqToken, "Bearer ")
			if len(splitToken) != 2 {
				next.ServeHTTP(w, r)
				return
			}
			tokenString := splitToken[1]

			token, err := utils.DecodeJwt(tokenString)
			if err != nil || !token.Valid {
				next.ServeHTTP(w, r)
				return
			}

			claims, ok := token.Claims.(*model.UserClaims)
			if !ok || claims == nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), AuthCtxKey, claims)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func GetUserClaimsFromContext(ctx context.Context) (*model.UserClaims, bool) {
	claims, ok := ctx.Value(AuthCtxKey).(*model.UserClaims)
	return claims, ok
}
