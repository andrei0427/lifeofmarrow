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
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Url    string `json:"url"`
}

type CTA struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Link  string `json:"link"`
}

type ContentSection struct {
	Id      int         `json:"id"`
	Title   string      `json:"title"`
	Image   StrapiMedia `json:"image"`
	Text    string      `json:"text"`
	CTALink CTA         `json:"ctalink"`
}
