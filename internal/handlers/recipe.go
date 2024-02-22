package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/andrei0427/lifeofmarrow/internal"
	"github.com/andrei0427/lifeofmarrow/view/layout"
	"github.com/andrei0427/lifeofmarrow/view/pages"
	"github.com/andrei0427/lifeofmarrow/view/pages/errors"
	"github.com/andrei0427/lifeofmarrow/view/partial"
)

const pageSize = 6

type RecipeFilters struct {
	MealType     string
	Cuisine      string
	Duration     string
	Requirements string
	Search       string
}

func readFilters(r *http.Request) *RecipeFilters {
	filters := RecipeFilters{}

	mealtype := r.URL.Query().Get("mealtype")
	if len(mealtype) > 0 {
		filters.MealType = mealtype
	}

	cuisine := r.URL.Query().Get("cuisine")
	if len(cuisine) > 0 {
		filters.Cuisine = cuisine
	}

	duration := r.URL.Query().Get("duration")
	if len(duration) > 0 {
		filters.Duration = duration
	}

	reqs := r.URL.Query().Get("requirements")
	if len(reqs) > 0 {
		filters.Requirements = reqs
	}

	search := r.URL.Query().Get("search")
	if len(search) > 0 {
		filters.Search = search
	}

	return &filters
}

func getRecipes(pageNo int, filters RecipeFilters) (*[]partial.RecipeEntity, error) {
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
		if len(filters.MealType) > 0 && !strings.Contains(strings.ToLower(filters.MealType), strings.ToLower(r.Attributes.MealType.Data.Attributes.Type)) {
			continue
		}

		if len(filters.Cuisine) > 0 && !strings.Contains(strings.ToLower(filters.Cuisine), strings.ToLower(r.Attributes.Cuisine.Data.Attributes.Cuisine)) {
			continue
		}

		if len(filters.Duration) > 0 && !strings.Contains(strings.ToLower(filters.Duration), strings.ToLower(r.Attributes.Duration.Data.Attributes.Time)) {
			continue
		}

		if len(filters.Requirements) > 0 {
			hasRequirement := false
			lowerCase := strings.ToLower(filters.Requirements)
			for _, req := range r.Attributes.Requirements.Data {
				if strings.Contains(lowerCase, strings.ToLower(req.Attributes.Requirement)) {
					hasRequirement = true
				}
			}

			if !hasRequirement {
				continue
			}
		}

		if len(filters.Search) > 0 {
			lowerCase := strings.ToLower(filters.Search)
			if !strings.Contains(strings.ToLower(r.Attributes.Title), lowerCase) &&
				!strings.Contains(strings.ToLower(r.Attributes.Ingredients), lowerCase) &&
				!strings.Contains(strings.ToLower(r.Attributes.Method), lowerCase) {
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

func HandleRecipes(w http.ResponseWriter, r *http.Request) {
	seo := layout.SEOInfo{
		Title:       "Recipes",
		Url:         "/recipes",
		Description: "A wide variety of colorful, healthy and vegan recipes from my collection",
	}

	mealTypes, err := internal.GetRecordCollectionFromStrapi[partial.MealType](internal.StrapiQueryOptions{
		Endpoint:     "meal-types",
		IsCollection: true,
	})
	if err != nil {
		layout.Base(seo, partial.Header(false), errors.InternalServerError()).Render(r.Context(), w)
		return
	}

	cuisines, err := internal.GetRecordCollectionFromStrapi[partial.Cuisine](internal.StrapiQueryOptions{
		Endpoint:     "cuisines",
		IsCollection: true,
	})
	if err != nil {
		layout.Base(seo, partial.Header(false), errors.InternalServerError()).Render(r.Context(), w)
		return
	}

	durations, err := internal.GetRecordCollectionFromStrapi[partial.Duration](internal.StrapiQueryOptions{
		Endpoint:     "durations",
		IsCollection: true,
	})
	if err != nil {
		layout.Base(seo, partial.Header(false), errors.InternalServerError()).Render(r.Context(), w)
		return
	}

	requirements, err := internal.GetRecordCollectionFromStrapi[partial.Requirements](internal.StrapiQueryOptions{
		Endpoint:     "requirements",
		IsCollection: true,
	})
	if err != nil {
		layout.Base(seo, partial.Header(false), errors.InternalServerError()).Render(r.Context(), w)
		return
	}

	props, err := getRecipes(1, *readFilters(r))
	if err != nil {
		layout.Base(seo, partial.Header(false), errors.InternalServerError()).Render(r.Context(), w)
		return
	}

	filterOpts := pages.RecipeFilters{
		MealTypes:    *mealTypes,
		Cuisines:     *cuisines,
		Durations:    *durations,
		Requirements: *requirements,
	}

	layout.Base(seo, partial.Header(false), pages.Recipes(*props, filterOpts)).Render(r.Context(), w)
}

func HandleRecipesPage(w http.ResponseWriter, r *http.Request) {
	sPageNo := r.PathValue("p")
	if len(sPageNo) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dirtyPageNo, err := strconv.Atoi(sPageNo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pageNo := max(1, dirtyPageNo)

	props, err := getRecipes(pageNo, *readFilters(r))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for i, rec := range *props {
		partial.RecipeTile(rec, pageNo, len(*props) == i+1).Render(r.Context(), w)
	}
}
