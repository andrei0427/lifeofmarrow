package internal

import (
	"testing"

	"github.com/andrei0427/lifeofmarrow/view/pages"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestStrapiApiUnwrapByAttributes(t *testing.T) {
	initEnv(t)

	res, err := FetchStrapi[pages.HomeEntity](StrapiQueryOptions{
		Endpoint:           "/home",
		UnwrapByAttributes: true,
	})

	assert.Nil(t, err)
	assert.True(t, len(res.CTA) > 0)
}

func TestStrapiApiUnwrapByAttributesWithRelations(t *testing.T) {
	initEnv(t)

	res, err := FetchStrapi[pages.HomeEntity](StrapiQueryOptions{
		Endpoint:           "/home",
		UnwrapByAttributes: true,
		Params: []StrapiKeyValue{
			{
				Key: "populate[0]", Value: "ImageCarousel",
			},
		},
	})

	assert.Nil(t, err)
	assert.True(t, len(res.ImageCarousel.Data) > 0)
}

func initEnv(t *testing.T) {
	err := godotenv.Load("../.env")
	assert.Nil(t, err)
}
