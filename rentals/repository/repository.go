package repository

import (
	"database/sql"
	"log"
	m "rentals/model"

	_ "github.com/lib/pq"
)

func FindRentals() []m.Rental {
	// TODO: externalize database properties
	connection, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/rentals?sslmode=disable")
	if err != nil {
		log.Fatal("error creating connection to database. " + err.Error())
	}
	defer connection.Close()

	query := `select id, name, description, type, vehicle_make, vehicle_model, vehicle_year, vehicle_length, sleeps, price_per_day, primary_image_url, home_city, home_state, home_zip, home_country, lat, lng from rentals`
	log.Println("executing query ", query)
	rows, err := connection.Query(query)
	if err != nil {
		log.Fatal("failure when executing the query. " + err.Error())
	}
	defer rows.Close()

	var rentals []m.Rental

	for rows.Next() {
		var r m.Rental
		err := rows.Scan(&r.ID, &r.Name, &r.Description, &r.Type, &r.Make, &r.Model, &r.Year, &r.Length, &r.Sleeps, &r.Price, &r.PrimaryImageURL, &r.City, &r.State, &r.Zip, &r.Country, &r.Latitude, &r.Longitude)

		if err != nil {
			log.Fatal("error scanning rows." + err.Error())
		}

		rentals = append(rentals, r)
	}

	return rentals
}
