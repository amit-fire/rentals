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

	query := `select r.id, name, description, type, vehicle_make, vehicle_model, vehicle_year, vehicle_length, sleeps, price_per_day, primary_image_url, home_city, home_state, home_zip, home_country, lat, lng, u.id, first_name, last_name FROM rentals r INNER JOIN users u ON user_id = u.id ORDER BY r.id ASC`
	log.Println("executing query ", query)
	rows, err := connection.Query(query)
	if err != nil {
		log.Fatal("failure when executing the query. " + err.Error())
	}
	defer rows.Close()

	var rentals []m.Rental

	for rows.Next() {
		var r m.DBRental
		var user m.User
		err := rows.Scan(&r.ID, &r.Name, &r.Description, &r.Type, &r.Make, &r.Model, &r.Year, &r.Length, &r.Sleeps, &r.Price, &r.PrimaryImageURL, &r.City, &r.State, &r.Zip, &r.Country, &r.Latitude, &r.Longitude, &user.ID, &user.FirstName, &user.LastName)

		if err != nil {
			log.Fatal("error scanning rows." + err.Error())
		}

		rental := m.Rental{
			ID:              r.ID,
			Name:            r.Name,
			Description:     r.Description,
			Type:            r.Type,
			Make:            r.Make,
			Model:           r.Model,
			Year:            r.Year,
			Length:          r.Length,
			Sleeps:          r.Sleeps,
			Price:           m.Price{Day: r.Price},
			PrimaryImageURL: r.PrimaryImageURL,
			Location:        m.Location{City: r.City, State: r.State, Zip: r.Zip, Country: r.Country, Latitude: r.Latitude, Longitude: r.Longitude},
			User:            m.User{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName},
		}

		rentals = append(rentals, rental)
	}

	return rentals
}
