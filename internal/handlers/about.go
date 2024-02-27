package handlers

import (
	"net/http"

	"github.com/andrei0427/lifeofmarrow/internal"
	"github.com/andrei0427/lifeofmarrow/view/pages"
	"github.com/andrei0427/lifeofmarrow/view/pages/errors"
)

func HandleAbout(w http.ResponseWriter, r *http.Request) {
	props, err := internal.GetRecordFromStrapi[pages.AboutEntity](internal.StrapiQueryOptions{
		Endpoint: "about",
		Params: []internal.StrapiKeyValue{
			{Key: "populate[0]", Value: "Sections"},
			{Key: "populate[1]", Value: "Sections.Image"},
			{Key: "populate[2]", Value: "Sections.CTALink"},
		},
	})
	if err != nil {
		errors.InternalServerError().Render(r.Context(), w)
		return
	}

	pages.About(props.Data.Attributes).Render(r.Context(), w)
}
