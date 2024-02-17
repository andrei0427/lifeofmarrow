package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type StrapiKeyValue struct {
	Key   string
	Value string
}
type StrapiQueryOptions struct {
	Endpoint           string
	UnwrapByAttributes bool
	Params             []StrapiKeyValue
}

func FetchStrapi[T any](opts StrapiQueryOptions) (*T, error) {
	if opts.Endpoint[0] == '/' {
		opts.Endpoint = strings.TrimPrefix(opts.Endpoint, "/")
	}

	strapiUrl := os.Getenv("STRAPI_URL")
	if len(strapiUrl) == 0 {
		log.Printf("strapi URL missing from env")
		return nil, fmt.Errorf("internal server error")
	}

	uri := fmt.Sprintf("%s/api/%s", strapiUrl, opts.Endpoint)

	if len(opts.Params) > 0 {
		params := url.Values{}
		for _, kv := range opts.Params {
			params.Set(kv.Key, kv.Value)
		}

		uri = fmt.Sprintf("%s?%s", uri, params.Encode())
	}

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Printf("error loading strapi api: %s", err)
		return nil, fmt.Errorf("internal server error")
	}

	token := os.Getenv("STRAPI_KEY")
	if len(token) == 0 {
		log.Printf("strapi token missing from env")
		return nil, fmt.Errorf("internal server error")
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("error when sending request to strapi: %s", err)
		return nil, fmt.Errorf("internal server error")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error when reading response from strapi: %s", err)
		return nil, fmt.Errorf("internal server error")
	}

	data, err := unmarshalStrapiResponse[T](body, opts.UnwrapByAttributes)
	if err != nil {
		log.Printf("error when unmarshalling response from strapi: %s", err)
		return nil, fmt.Errorf("internal server error")
	}

	return data, nil
}

type StrapiResponse[T any] struct {
	Data T `json:"data"`
}

type StrapiResponseWrappedByAttributes[T any] struct {
	Data struct {
		Attributes T `json:"attributes"`
	} `json:"data"`
}

func unmarshalStrapiResponse[T any](body []byte, unwrapByAttributes bool) (*T, error) {
	if unwrapByAttributes {
		data := new(StrapiResponseWrappedByAttributes[T])
		err := json.Unmarshal(body, &data)
		if err != nil {
			return nil, err
		}

		return &data.Data.Attributes, nil
	} else {
		data := new(StrapiResponse[T])
		err := json.Unmarshal(body, &data)
		if err != nil {
			return nil, err
		}

		return &data.Data, nil
	}
}
