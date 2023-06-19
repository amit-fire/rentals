package model

type Rental struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Type            string  `json:"type"`
	Make            string  `json:"make"`
	Model           string  `json:"model"`
	Year            int     `json:"year"`
	Length          float64 `json:"length"`
	Sleeps          int     `json:"sleeps"`
	PrimaryImageURL string  `json:"primary_image_url"`
	Price           int     `json:"price"`
	City            string  `json:"city"`
	State           string  `json:"state"`
	Zip             string  `json:"zip"`
	Country         string  `json:"country"`
	Latitude        float64 `json:"lat"`
	Longitude       float64 `json:"lng"`
}
