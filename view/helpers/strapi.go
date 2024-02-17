package helpers

import (
	"os"
)

func GetStrapiMediaUrl(url string) string {
	strapiUrl := os.Getenv("STRAPI_URL")

	if len(url) == 0 || len(strapiUrl) == 0 {
		return "/static/img/logo-responsive.png"
	}

	return strapiUrl + url
}
