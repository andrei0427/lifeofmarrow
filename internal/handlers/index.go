package handlers

import (
	"net/http"

	"github.com/andrei0427/lifeofmarrow/internal"
	"github.com/andrei0427/lifeofmarrow/view/layout"
	"github.com/andrei0427/lifeofmarrow/view/pages"
	"github.com/andrei0427/lifeofmarrow/view/pages/errors"
	"github.com/andrei0427/lifeofmarrow/view/partial"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	// Handle attempts to browse pages that are not handled in the mux
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", http.StatusMovedPermanently)
		return
	}

	seo := layout.SEOInfo{
		Title:  "Home",
		Url:    "/",
		IsHome: true,
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
		layout.Base(seo, partial.Header(false), errors.InternalServerError()).Render(r.Context(), w)
	}

	layout.Base(seo, partial.Header(seo.IsHome), pages.Index(props.Data.Attributes)).Render(r.Context(), w)
}
