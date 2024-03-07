package handlers

import (
	"net/http"

	"github.com/andrei0427/lifeofmarrow/internal"
	"github.com/andrei0427/lifeofmarrow/view/helpers"
	"github.com/andrei0427/lifeofmarrow/view/pages"
	"github.com/andrei0427/lifeofmarrow/view/pages/errors"
)

func HandleBooks(w http.ResponseWriter, r *http.Request) {
	isHx := r.Header.Get("Hx-Request") == "true"

	collection, err := internal.GetRecordCollectionFromStrapi[helpers.StoreItem](internal.StrapiQueryOptions{
		Endpoint: "/services",
		Params: []internal.StrapiKeyValue{
			{Key: "populate[0]", Value: "Images"},
		},
		Filters: []internal.StrapiFilter{
			{FieldName: "Type", Operator: "$eq", Value: "Book"},
		},
	})

	if err != nil {
		if isHx {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			errors.InternalServerError().Render(r.Context(), w)
		}

		return
	}

	books := make([]helpers.StoreItem, 0)
	for _, b := range collection.Data {
		books = append(books, b.Attributes)
	}

	pages.Books(books).Render(r.Context(), w)
}
