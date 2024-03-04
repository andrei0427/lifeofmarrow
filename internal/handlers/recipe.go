package handlers

import (
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/andrei0427/lifeofmarrow/internal"
	"github.com/andrei0427/lifeofmarrow/view/pages"
	"github.com/andrei0427/lifeofmarrow/view/pages/errors"
	"github.com/andrei0427/lifeofmarrow/view/partial"
	"github.com/gosimple/slug"
)

const pageSize = 6

func getRecipes(pageNo int, filters partial.SelectedRecipeFilters) (*[]partial.RecipeEntity, error) {
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
		return nil, err
	}

	filtered := new([]partial.RecipeEntity)
	for _, r := range collection.Data {

		if len(filters.SelectedMealTypes.Data) > 0 && !slices.Contains(filters.SelectedMealTypes.Data, r.Attributes.MealType.Data) {
			continue
		}

		if len(filters.SelectedCuisines.Data) > 0 && !slices.Contains(filters.SelectedCuisines.Data, r.Attributes.Cuisine.Data) {
			continue
		}

		if len(filters.SelectedDurations.Data) > 0 && !slices.Contains(filters.SelectedDurations.Data, r.Attributes.Duration.Data) {
			continue
		}

		if len(filters.Requirements.Data) > 0 {
			hasRequirement := false

			for _, req := range r.Attributes.Requirements.Data {
				if slices.Contains(filters.Requirements.Data, req) {
					hasRequirement = true
				}
			}

			if !hasRequirement {
				continue
			}
		}

		if len(filters.SearchText) > 0 {
			lowerCase := strings.ToLower(filters.SearchText)
			if !strings.Contains(strings.ToLower(r.Attributes.Title), lowerCase) &&
				!strings.Contains(strings.ToLower(r.Attributes.Ingredients), lowerCase) &&
				!strings.Contains(strings.ToLower(r.Attributes.Method), lowerCase) {
				continue
			}
		}

		if len(filters.Slug) > 0 {
			if slug.Make(r.Attributes.Title) != filters.Slug {
				continue
			}
		}

		*filtered = append(*filtered, r.Attributes)
	}

	c := len(*filtered)
	start := (pageNo - 1) * pageSize
	end := min(c, start+pageSize)

	props := new([]partial.RecipeEntity)
	if end > c {
		return props, nil
	}

	for i := start; i < end; i++ {
		*props = append(*props, (*filtered)[i])
	}

	return props, nil
}

func HandleRecipesPage(w http.ResponseWriter, r *http.Request) {
	sPageNo := r.PathValue("p")
	isHx := r.Header.Get("Hx-Request") == "true"

	filters, err := partial.NewSelectedRecipeFilters(r)
	if err != nil {
		if isHx {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			errors.InternalServerError().Render(r.Context(), w)
		}
		return
	}

	dirtyPageNo, err := strconv.Atoi(sPageNo)
	if err != nil {
		dirtyPageNo = 1
	}
	pageNo := max(1, dirtyPageNo)

	props, err := getRecipes(pageNo, *filters)
	if err != nil {
		if isHx {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			errors.InternalServerError().Render(r.Context(), w)
		}
		return
	}

	if isHx {
		partial.RecipeFilters(*filters).Render(r.Context(), w)
		for i, rec := range *props {
			partial.RecipeTile(rec, pageNo, len(*props) == i+1).Render(r.Context(), w)
		}
	} else {
		pages.Recipes(*props, *filters).Render(r.Context(), w)
	}
}

func HandleRecipePage(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if len(slug) == 0 {
		errors.NotFoundError().Render(r.Context(), w)
		return
	}

	recipe, err := getRecipes(1, partial.SelectedRecipeFilters{
		Slug: slug,
	})
	if err != nil || len(*recipe) == 0 {
		errors.NotFoundError().Render(r.Context(), w)
		return
	}

	pages.Recipe((*recipe)[0]).Render(r.Context(), w)
}
