package handlers

import (
	"net/http"

	"github.com/andrei0427/lifeofmarrow/internal"
	"github.com/andrei0427/lifeofmarrow/view/pages"
	"github.com/andrei0427/lifeofmarrow/view/pages/errors"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	// Handle attempts to browse pages that are not handled in the mux
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", http.StatusMovedPermanently)
		return
	}

	props, err := internal.GetRecordFromStrapi[pages.HomeEntity](internal.StrapiQueryOptions{
		Endpoint: "home",
		Params: []internal.StrapiKeyValue{
			{
				Key: "populate[0]", Value: "ImageCarousel",
			},
		},
	})
	if err != nil {
		errors.InternalServerError().Render(r.Context(), w)
	}

	pages.Index(props.Data.Attributes).Render(r.Context(), w)
}
