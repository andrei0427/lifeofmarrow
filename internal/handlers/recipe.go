package handlers

import (
	"net/http"

	"github.com/andrei0427/lifeofmarrow/internal"
	"github.com/andrei0427/lifeofmarrow/view/layout"
	"github.com/andrei0427/lifeofmarrow/view/pages"
	"github.com/andrei0427/lifeofmarrow/view/pages/error"
	"github.com/andrei0427/lifeofmarrow/view/partial"
)

func HandleRecipes(w http.ResponseWriter, r *http.Request) {
	seo := layout.SEOInfo{
		Title:       "Recipes",
		Url:         "/recipes",
		Description: "A wide variety of colorful, healthy and vegan recipes from my collection",
	}

	collection, err := internal.GetRecordCollectionFromStrapi[partial.RecipeEntity](internal.StrapiQueryOptions{
		Endpoint: "recipes",
		Params: []internal.StrapiKeyValue{
			{Key: "populate[0]", Value: "meal_type"},
			{Key: "populate[1]", Value: "cuisine"},
			{Key: "populate[2]", Value: "duration"},
			{Key: "populate[3]", Value: "requirements"},
			{Key: "populate[4]", Value: "Images"},
		},
	})
	if err != nil {
		layout.Base(seo, partial.Header(false), error.InternalServerError()).Render(r.Context(), w)
		return
	}

	props := new([]partial.RecipeEntity)
	for _, r := range collection.Data {
		*props = append(*props, r.Attributes)
	}

	layout.Base(seo, partial.Header(false), pages.Recipes(*props)).Render(r.Context(), w)
}
