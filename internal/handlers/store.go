package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/andrei0427/lifeofmarrow/internal"
	"github.com/andrei0427/lifeofmarrow/view/helpers"
	"github.com/andrei0427/lifeofmarrow/view/pages"
	"github.com/andrei0427/lifeofmarrow/view/pages/errors"
	"github.com/andrei0427/lifeofmarrow/view/partial"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/price"
	"github.com/stripe/stripe-go/webhook"
)

func HandleBooks(w http.ResponseWriter, r *http.Request) {
	isHx := r.Header.Get("Hx-Request") == "true"

	collection, err := internal.GetRecordCollectionFromStrapi[helpers.StoreItem](internal.StrapiQueryOptions{
		Endpoint: "/services",
		Params: []internal.StrapiKeyValue{
			{Key: "populate[0]", Value: "Images"},
			{Key: "populate[1]", Value: "Shipping"},
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
			{Key: "populate[1]", Value: "Shipping"},
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
	domain := "https://test.lifeofmarrow.com"
	if os.Getenv("DEV") == "true" {
		domain = "http://localhost:8080"
	}

	itemId := r.PathValue("id")
	if len(itemId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	item, err := internal.GetRecordFromStrapi[helpers.StoreItem](internal.StrapiQueryOptions{
		Endpoint: fmt.Sprintf("/services/%s", itemId),
		Params: []internal.StrapiKeyValue{
			{Key: "populate[0]", Value: "Shipping"},
		},
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
	rawQuantity := r.FormValue("quantity")
	shipTo := r.FormValue("shipTo")
	couponId := r.FormValue("couponId")

	possibleShippingLocations := []string{"MT", "EU", "US", "GB", "AU"}
	if len(shipTo) == 0 {
		shipTo = "MT"
	}

	if !slices.Contains(possibleShippingLocations, shipTo) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	allowableShippingLocations := []string{}
	shippingRateId, _ := item.Data.Attributes.GetShippingRateId(shipTo)
	switch shipTo {
	case "EU":
		{
			allowableShippingLocations = append(allowableShippingLocations,
				"AT",
				"BE",
				"BG",
				"CY",
				"CZ",
				"DK",
				"EE",
				"FI",
				"FR",
				"DE",
				"GR",
				"HU",
				"IT",
				"LV",
				"LT",
				"LU",
				"NL",
				"PL",
				"PT",
				"RO",
				"SK",
				"SI",
				"ES",
				"SE")
			break
		}
	case "GB":
		{
			allowableShippingLocations = append(allowableShippingLocations, "GB", "IE")
			break
		}
	case "US":
		{
			allowableShippingLocations = append(allowableShippingLocations, "US")
			break
		}
	case "AU":
		{
			allowableShippingLocations = append(allowableShippingLocations, "AU")
			break
		}
	case "MT":
		{
			allowableShippingLocations = append(allowableShippingLocations, "MT")
		}
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

	if item.Data.Attributes.MinOrder == 0 {
		item.Data.Attributes.MinOrder = 1
	}

	if item.Data.Attributes.MaxOrder == 0 {
		item.Data.Attributes.MaxOrder = 999
	}

	discounts := []*stripe.CheckoutSessionDiscountParams{}
	if len(couponId) > 0 {
		discounts = append(discounts, &stripe.CheckoutSessionDiscountParams{
			Coupon: stripe.String(couponId),
		})
	}

	s := stripe.CheckoutSessionParams{
		PaymentMethodTypes: []*string{stripe.String(paymType)},
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(priceId),
				Quantity: stripe.Int64(int64(quantity)),
				AdjustableQuantity: &stripe.CheckoutSessionLineItemAdjustableQuantityParams{
					Enabled: stripe.Bool(true),
					Minimum: stripe.Int64(int64(item.Data.Attributes.MinOrder)),
					Maximum: stripe.Int64(int64(item.Data.Attributes.MaxOrder)),
				},
			},
		},
		ShippingAddressCollection: &stripe.CheckoutSessionShippingAddressCollectionParams{
			AllowedCountries: stripe.StringSlice(allowableShippingLocations),
		},
		Mode:                  stripe.String("payment"),
		PhoneNumberCollection: &stripe.CheckoutSessionPhoneNumberCollectionParams{Enabled: stripe.Bool(true)},
		SuccessURL:            stripe.String(domain + "/store/checkout/success"),
		CancelURL:             stripe.String(domain + "/store/checkout/cancel"),
	}

	if len(discounts) > 0 {
		s.Discounts = discounts
	}

	if len(shippingRateId) > 0 {
		s.ShippingOptions = []*stripe.CheckoutSessionShippingOptionParams{
			{ShippingRate: stripe.String(shippingRateId)},
		}
	}

	stripeSession, err := session.New(&s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, stripeSession.URL, http.StatusSeeOther)
}

func HandleCheckoutCancel(w http.ResponseWriter, r *http.Request) {
	errors.PaymentCancelled().Render(r.Context(), w)
}

func HandleCheckoutSuccess(w http.ResponseWriter, r *http.Request) {
	pages.PaymentSuccessful().Render(r.Context(), w)
}

func HandleStripeWebhook(w http.ResponseWriter, req *http.Request) {
	endpointSecret, ok := os.LookupEnv("STRIPE_WEBHOOK_SECRET")
	if !ok {
		fmt.Fprintf(os.Stderr, "Stripe webhook secret not found in .env")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	const MaxBodyBytes = int64(65536)
	req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
	payload, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	// Pass the request body and Stripe-Signature header to ConstructEvent, along
	// with the webhook signing key.
	event, err := webhook.ConstructEvent(payload, req.Header.Get("Stripe-Signature"),
		endpointSecret)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error verifying webhook signature: %v\n", err)
		w.WriteHeader(http.StatusBadRequest) // Return a 400 error on a bad signature
		return
	}

	// Unmarshal the event data into an appropriate struct depending on its Type
	switch event.Type {
	case "checkout.session.completed":
		{
			eventId, ok := event.Data.Object["id"]
			if !ok {
				fmt.Fprintf(os.Stderr, "Id not found when deserialising event data object")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			session, err := session.Get(eventId.(string), &stripe.CheckoutSessionParams{
				Expand: []*string{
					stripe.String("line_items"),
				},
			})
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error retrieving back session: %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			line_items := session.LineItems
			iShipping, ok := event.Data.Object["shipping"]
			if !ok {
				fmt.Fprintf(os.Stderr, "Shipping property not found on event data object")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			shipping, ok := iShipping.(map[string]interface{})
			if !ok {
				fmt.Fprintf(os.Stderr, "Failed to cast shipping object to map")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			customerName, ok := shipping["name"]
			if !ok {
				fmt.Fprintf(os.Stderr, "Customer name property not found on shipping object")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			orderNo, ok := event.Data.Object["payment_intent"]
			if !ok {
				fmt.Fprintf(os.Stderr, "Payment intent property not found on event data object")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			orderDate := time.Now().Format(time.RFC1123)

			strOrderTotal, ok := event.Data.Object["amount_total"]
			if !ok {
				fmt.Fprintf(os.Stderr, "Amount total property not found on event data object")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			orderTotalInt, err := strconv.Atoi(strOrderTotal.(string))
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing order total: %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			orderTotal := fmt.Sprintf("€%.2f", float64(orderTotalInt/100))

			firstLineItem := line_items.Data[0]
			if firstLineItem == nil {
				fmt.Fprintf(os.Stderr, "Could not find line item in order")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			quantity := firstLineItem.Quantity
			price := fmt.Sprintf("€%.2f", float64(firstLineItem.AmountTotal/100))
			shippingPrice := fmt.Sprintf("€%.2f", float64((orderTotalInt-int(firstLineItem.AmountTotal))/100))

			iCustomerDetails, ok := event.Data.Object["customer_details"]
			if !ok {
				fmt.Fprintf(os.Stderr, "Customer details property not found on event data object")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			customerDetails, ok := iCustomerDetails.(map[string]string)
			if !ok {
				fmt.Fprintf(os.Stderr, "Failed to cast customer details obj to map")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			phoneNo, ok := customerDetails["phone"]
			if !ok {
				phoneNo = "No phone number provided"
			}

			iAddress, ok := shipping["address"]
			if !ok {
				fmt.Fprintf(os.Stderr, "Address property not found on shipping object")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			addr, ok := iAddress.(map[string]string)
			if !ok {
				fmt.Fprintf(os.Stderr, "Failed to cast address obj to map")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			line1, ok := addr["line1"]
			if !ok {
				line1 = ""
			}

			line2, ok := addr["line2"]
			if !ok {
				line2 = ""
			}

			postCode, ok := addr["postal_code"]
			if !ok {
				postCode = ""
			}

			city, ok := addr["city"]
			if !ok {
				city = ""
			}

			state, ok := addr["state"]
			if !ok {
				state = ""
			}

			country, ok := addr["country"]
			if !ok {
				country = ""
			}

			address := strings.Join([]string{
				line1,
				line2,
				postCode,
				city,
				state,
				country,
			}, "<br />")

			customerEmail, ok := customerDetails["email"]
			if !ok {
				customerEmail = ""
			}

			bodyBuf := []byte{}
			bodyWriter := bytes.NewBuffer(bodyBuf)
			model := partial.EmailContent{
				CustomerName:  customerName.(string),
				Address:       address,
				OrderNo:       orderNo.(string),
				OrderDate:     orderDate,
				OrderTotal:    orderTotal,
				Quantity:      int(quantity),
				Price:         price,
				ShippingPrice: shippingPrice,
				PhoneNo:       phoneNo,
				DanielEmail:   "daniel@lifeofmarrow.com",
			}

			partial.OrderEmail(model).Render(req.Context(), bodyWriter)

			internal.SendEmail(customerEmail, fmt.Sprintf(
				"Life of Marrow Order - Receipt [%s]", orderDate),
				string(bodyBuf),
				true)
		}

	default:
		fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
	}

	w.WriteHeader(http.StatusOK)
}
