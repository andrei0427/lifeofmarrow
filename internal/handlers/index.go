package handlers

import (
	"github.com/a-h/templ"
	"github.com/andrei0427/lifeofmarrow/internal"
	"github.com/andrei0427/lifeofmarrow/view/layout"
	"github.com/andrei0427/lifeofmarrow/view/pages"
	"github.com/andrei0427/lifeofmarrow/view/pages/error"
	"github.com/andrei0427/lifeofmarrow/view/partial"
)

func HandleIndex() *templ.ComponentHandler {
	seo := layout.SEOInfo{
		Title:  "Home",
		Url:    "/",
		IsHome: true,
	}

	props, err := internal.FetchStrapi[pages.HomeEntity](internal.StrapiQueryOptions{
		Endpoint:           "home",
		UnwrapByAttributes: true,
		Params: []internal.StrapiKeyValue{
			{
				Key: "populate[0]", Value: "ImageCarousel",
			},
		},
	})
	if err != nil {
		return templ.Handler(layout.Base(seo, partial.Header(false), error.InternalServerError()))
	}

	return templ.Handler(layout.Base(seo, partial.Header(seo.IsHome), pages.Index(*props)))
}
