package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/wycliff-ochieng/todo/data"
)

func TaskValidateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tsk data.Task

		err := json.NewDecoder(r.Body).Decode(&tsk)
		if err != nil {
			log.Println("[ERROR]")
			http.Error(w, "inavalid json response", http.StatusMethodNotAllowed)
		}

		//field validation
		validate := validator.New()

		err = validate.Struct(tsk)
		if err != nil {
			http.Error(w, "validation failed", http.StatusBadRequest)
		}

		ctx := context.WithValue(r.Context(), taskkey{}, tsk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
