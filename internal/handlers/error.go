package handlers

import (
	"net/http"

	"github.com/andrei0427/lifeofmarrow/view/pages/errors"
)

func Handle404(w http.ResponseWriter, r *http.Request) {
	errors.NotFoundError().Render(r.Context(), w)
}

func Handle503(w http.ResponseWriter, r *http.Request) {
	errors.InternalServerError().Render(r.Context(), w)
}
