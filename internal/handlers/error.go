package handlers

import (
	"net/http"

	"github.com/andrei0427/lifeofmarrow/view/layout"
	"github.com/andrei0427/lifeofmarrow/view/pages/errors"
	"github.com/andrei0427/lifeofmarrow/view/partial"
)

func Handle404(w http.ResponseWriter, r *http.Request) {
	layout.Base(
		layout.SEOInfo{
			Title: "404",
			Url:   "/404",
		},
		partial.Header(false),
		errors.NotFoundError(),
	).Render(r.Context(), w)
}

func Handle503(w http.ResponseWriter, r *http.Request) {
	layout.Base(
		layout.SEOInfo{
			Title: "503",
			Url:   "/503",
		},
		partial.Header(false),
		errors.InternalServerError(),
	).Render(r.Context(), w)
}
