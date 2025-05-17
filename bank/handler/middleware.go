package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/wycliff-ochieng/bank/data"
)

type bookkey struct{}

func BookValidateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bk := data.Book{}

		//decording json body into product struct
		err := json.NewDecoder(r.Body).Decode(&bk)
		if err != nil {
			http.Error(w, "Invalid json input", http.StatusInternalServerError)
		}

		//validate the book fields
		var validate = validator.New()

		err = validate.Struct(bk)
		if err != nil {
			http.Error(w, "validation failed: ", http.StatusNotImplemented)
			return
		}

		//Add product to the context
		ctx := context.WithValue(r.Context(), bookkey{}, bk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
