package query

func query() string {
	return `select r.id, name, description, type, vehicle_make, vehicle_model, vehicle_year, vehicle_length, sleeps, price_per_day, primary_image_url, home_city, home_state, home_zip, home_country, lat, lng, u.id, first_name, last_name FROM rentals r INNER JOIN users u ON user_id = u.id`
}

func DefaultQuery() string {
	return query() + orderBy()
}

func IDQuery(id string) string {
	return query() + ` WHERE r.id in (` + id + `)` + orderBy()
}

func orderBy() string {
	return ` ORDER BY r.id ASC`
}
