package helpers

type StrapiMediaCollection struct {
	Data []struct {
		Attributes StrapiMediaType `json:"attributes"`
	} `json:"data"`
}

type StrapiMedia struct {
	Data struct {
		Attributes StrapiMediaType `json:"attributes"`
	} `json:"data"`
}

type StrapiMediaType struct {
	Name   string `json:"name"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Url    string `json:"url"`
}

type CTA struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}
