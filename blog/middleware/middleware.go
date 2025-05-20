package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/wycliff-ochieng/blog/data"
	"golang.org/x/net/context"
)

type PostKey struct{}

func ValidatePostMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var pst data.Post

		err := json.NewDecoder(r.Body).Decode(&pst)
		if err != nil {
			log.Println("[ERROR},[ERROR],[ERROR]")
			http.Error(w, "Unable to parse json", http.StatusBadRequest)
		}

		validate := validator.New()

		err = validate.Struct(pst)
		if err != nil {
			http.Error(w, "Validation failed", http.StatusBadRequest)
		}

		ctx := context.WithValue(r.Context(), PostKey{}, pst)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
