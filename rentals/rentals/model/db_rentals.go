package model

type DBRental struct {
	ID              int
	Name            string
	Description     string
	Type            string
	Make            string
	Model           string
	Year            int
	Length          float64
	Sleeps          int
	PrimaryImageURL string
	Price           int
	City            string
	State           string
	Zip             string
	Country         string
	Latitude        float64
	Longitude       float64
}
