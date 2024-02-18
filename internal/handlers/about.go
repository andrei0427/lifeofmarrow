package handlers

import (
	"net/http"

	"github.com/andrei0427/lifeofmarrow/internal"
	"github.com/andrei0427/lifeofmarrow/view/layout"
	"github.com/andrei0427/lifeofmarrow/view/pages"
	"github.com/andrei0427/lifeofmarrow/view/pages/error"
	"github.com/andrei0427/lifeofmarrow/view/partial"
)

func HandleAbout(w http.ResponseWriter, r *http.Request) {
	seo := layout.SEOInfo{
		Title: "About",
		Url:   "/about",
	}

	props, err := internal.GetRecordFromStrapi[pages.AboutEntity](internal.StrapiQueryOptions{
		Endpoint: "about",
		Params: []internal.StrapiKeyValue{
			{Key: "populate[0]", Value: "Sections"},
			{Key: "populate[1]", Value: "Sections.Image"},
			{Key: "populate[2]", Value: "Sections.CTALink"},
		},
	})
	if err != nil {
		layout.Base(seo, partial.Header(false), error.InternalServerError()).Render(r.Context(), w)
		return
	}

	layout.Base(seo, partial.Header(seo.IsHome), pages.About(props.Data.Attributes)).Render(r.Context(), w)
}
