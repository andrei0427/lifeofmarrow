package internal

import (
	"testing"

	"github.com/andrei0427/lifeofmarrow/view/helpers"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

type Home struct {
	CTA           string                        `json:"CTA"`
	ImageCarousel helpers.StrapiMediaCollection `json:"ImageCarousel"`
}

func TestStrapiApiUnwrapByAttributes(t *testing.T) {
	initEnv(t)

	res, err := GetRecordFromStrapi[Home](StrapiQueryOptions{
		Endpoint: "/home",
	})

	assert.Nil(t, err)
	assert.True(t, len(res.Data.Attributes.CTA) > 0)
}

func TestStrapiApiUnwrapByAttributesWithRelations(t *testing.T) {
	initEnv(t)

	res, err := GetRecordFromStrapi[Home](StrapiQueryOptions{
		Endpoint: "/home",
		Params: []StrapiKeyValue{
			{
				Key: "populate[0]", Value: "ImageCarousel",
			},
		},
	})

	assert.Nil(t, err)
	assert.True(t, len(res.Data.Attributes.ImageCarousel.Data) > 0)
}

func initEnv(t *testing.T) {
	err := godotenv.Load("../.env")
	assert.Nil(t, err)
}
