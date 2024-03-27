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

type StrapiFilter struct {
	FieldName string
	Operator  string
	Value     string
}

type StrapiQueryOptions struct {
	Endpoint string
	Params   []StrapiKeyValue
	Filters  []StrapiFilter
}

type StrapiResponse[T any] struct {
	Data struct {
		id         int
		Attributes T `json:"attributes"`
	} `json:"data"`
}

type StrapiCollectionResponse[T any] struct {
	Data []struct {
		Id         int `json:"id"`
		Attributes T   `json:"attributes"`
	} `json:"data"`
}

func GetRecordFromStrapi[T any](opts StrapiQueryOptions) (*StrapiResponse[T], error) {
	if cached, ok := Cache.Get(opts.Endpoint); ok {
		cast := (*cached).(*StrapiResponse[T])
		return cast, nil
	}

	body, err := fetchStrapi(opts)
	if err != nil {
		return nil, err
	}

	data := new(StrapiResponse[T])
	err = json.Unmarshal(*body, data)
	if err != nil {
		log.Printf("error when unmarshalling response from strapi: %s\n%s\n\n", err, *body)
		return nil, fmt.Errorf("internal server error")
	}

	Cache.Set(opts.Endpoint, data)

	return data, nil
}

func GetRecordCollectionFromStrapi[T any](opts StrapiQueryOptions) (*StrapiCollectionResponse[T], error) {
	if cached, ok := Cache.Get(opts.Endpoint); ok {
		cast := (*cached).(*StrapiCollectionResponse[T])
		return cast, nil
	}

	body, err := fetchStrapi(opts)
	if err != nil {
		return nil, err
	}

	data := new(StrapiCollectionResponse[T])
	err = json.Unmarshal(*body, data)
	if err != nil {
		log.Printf("error when unmarshalling response from strapi: %s", err)
		return nil, fmt.Errorf("internal server error")
	}

	Cache.Set(opts.Endpoint, data)

	return data, nil
}

func fetchStrapi(opts StrapiQueryOptions) (*[]byte, error) {
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

	if len(opts.Filters) > 0 {
		prefix := "?"
		if strings.Contains(uri, "?") {
			prefix = "&"
		}

		for _, f := range opts.Filters {
			uri = fmt.Sprintf("%s%sfilters[%s][%s]=%s", uri, prefix, f.FieldName, f.Operator, f.Value)
			prefix = "&"
		}
	}

	prefix := "?"
	if strings.Contains(uri, "?") {
		prefix = "&"
	}
	uri = fmt.Sprintf("%s%spagination[pageSize]=1000", uri, prefix)

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

	return &body, nil
}
