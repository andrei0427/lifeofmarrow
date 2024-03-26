package handlers

import (
	"fmt"
	"net/http"
	"os"
	"slices"
	"strconv"

	"github.com/andrei0427/lifeofmarrow/internal"
	"github.com/andrei0427/lifeofmarrow/view/helpers"
	"github.com/andrei0427/lifeofmarrow/view/pages"
	"github.com/andrei0427/lifeofmarrow/view/pages/errors"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/price"
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
		book := b.Attributes

		if len(book.StripePriceId) > 0 {
			bookPriceCacheKey := fmt.Sprintf("price-%s", b.Attributes.StripePriceId)
			if ok := internal.Cache.Has(bookPriceCacheKey); !ok {
				p, err := price.Get(b.Attributes.StripePriceId, &stripe.PriceParams{})
				if err != nil {
					errors.InternalServerError().Render(r.Context(), w)
					return
				}

				price := fmt.Sprintf("%.2f", p.UnitAmountDecimal/100)

				internal.Cache.Set(bookPriceCacheKey, price)
				book.Price = price
			} else {
				price, _ := internal.Cache.Get(bookPriceCacheKey)
				book.Price = (*price).(string)
			}
		}

		if len(book.StripeDiscountedPriceId) > 0 {
			bookPriceCacheKey := fmt.Sprintf("price-%s", b.Attributes.StripeDiscountedPriceId)
			if ok := internal.Cache.Has(bookPriceCacheKey); !ok {
				p, err := price.Get(b.Attributes.StripeDiscountedPriceId, &stripe.PriceParams{})
				if err != nil {
					errors.InternalServerError().Render(r.Context(), w)
					return
				}

				price := fmt.Sprintf("%.2f", p.UnitAmountDecimal/100)

				internal.Cache.Set(bookPriceCacheKey, price)
				book.DiscountedPrice = price
			} else {
				price, _ := internal.Cache.Get(bookPriceCacheKey)
				book.DiscountedPrice = (*price).(string)
			}
		}

		book.Id = b.Id
		books = append(books, book)
	}

	pages.Books(books).Render(r.Context(), w)
}

func HandleFood(w http.ResponseWriter, r *http.Request) {
	isHx := r.Header.Get("Hx-Request") == "true"

	collection, err := internal.GetRecordCollectionFromStrapi[helpers.StoreItem](internal.StrapiQueryOptions{
		Endpoint: "/services",
		Params: []internal.StrapiKeyValue{
			{Key: "populate[0]", Value: "Images"},
		},
		Filters: []internal.StrapiFilter{
			{FieldName: "Type", Operator: "$eq", Value: "Food"},
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

	food := make([]helpers.StoreItem, 0)
	for _, f := range collection.Data {
		f.Attributes.Id = f.Id
		food = append(food, f.Attributes)
	}

	pages.Food(food).Render(r.Context(), w)
}

func HandleCheckout(w http.ResponseWriter, r *http.Request) {
	domain := "https://lifeofmarrow.com"
	if os.Getenv("DEV") == "true" {
		domain = "http://localhost:7331"
	}

	itemId := r.PathValue("id")
	if len(itemId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	item, err := internal.GetRecordFromStrapi[helpers.StoreItem](internal.StrapiQueryOptions{
		Endpoint: fmt.Sprintf("/services/%s", itemId),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if item == nil || len(item.Data.Attributes.StripePriceId) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	paymType := "card"
	rawQuantity := r.Form.Get("quantity")
	shipTo := r.Form.Get("shipTo")
	// couponId := r.Form.Get("couponId")

	allowableShippingLocations := []string{"MT", "EU", "US", "GB", "AU"}

	if len(shipTo) > 0 && !slices.Contains(allowableShippingLocations, shipTo) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	quantity := 1
	if len(rawQuantity) > 0 {
		parsed, err := strconv.Atoi(rawQuantity)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		quantity = parsed
	}

	priceId := item.Data.Attributes.StripePriceId
	if len(item.Data.Attributes.StripeDiscountedPriceId) > 0 {
		priceId = item.Data.Attributes.StripeDiscountedPriceId
	}

	s := stripe.CheckoutSessionParams{
		PaymentMethodTypes: []*string{stripe.String(paymType)},
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(priceId),
				Quantity: stripe.Int64(int64(quantity)),
			},
		},
		Mode:                  stripe.String("payment"),
		PhoneNumberCollection: &stripe.CheckoutSessionPhoneNumberCollectionParams{Enabled: stripe.Bool(true)},
		SuccessURL:            stripe.String(domain + "/store/checkout/success"),
		CancelURL:             stripe.String(domain + "/store/checkout/cancel"),
	}

	stripeSession, err := session.New(&s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, stripeSession.URL, http.StatusSeeOther)
}
