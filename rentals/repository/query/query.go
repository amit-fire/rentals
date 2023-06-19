package query

import (
	"strings"

	m "rentals/model"

	u "rentals/utils"
)

func QueryBuilder(qp m.QueryParameter) string {
	query := defaultQuery()

	if qp.Ids != "" {
		query = addIds(query, qp.Ids)
	}

	if qp.PriceMin != "" {
		query = andWhere(query)
		query += u.PricePerDay + ` > ` + qp.PriceMin
	}

	if qp.PriceMax != "" {
		query = andWhere(query)
		query += u.PricePerDay + ` < ` + qp.PriceMax
	}

	if qp.Near != "" {
		query = nearQuery(qp.Near, query)
	}

	query = orderBy(qp.Sort, query)

	if qp.Limit != "" {
		query += ` LIMIT ` + qp.Limit
	}

	if qp.Offset != "" {
		query += ` OFFSET ` + qp.Offset
	}

	return query
}

func IDsQuery(ids string) string {
	query := defaultQuery()
	query = addIds(query, ids)
	return query
}

func defaultQuery() string {
	return `SELECT r.id, name, description, type, vehicle_make, vehicle_model, ` + u.VehicleYear + `, ` + u.VehicleLength + `, ` + u.Sleeps + `, ` + u.PricePerDay + `, primary_image_url, home_city, home_state, home_zip, home_country, ` + u.Lat + `, ` + u.Lng + `, u.id, first_name, last_name FROM rentals r INNER JOIN users u ON user_id = u.id`
}

func addIds(query string, ids string) string {
	query += ` WHERE r.id in (` + ids + `)`
	return query
}

func andWhere(query string) string {
	if strings.Contains(query, "WHERE") {
		query += ` AND `
	} else {
		query += ` WHERE `
	}

	return query
}

func orderBy(sort string, query string) string {
	if sort == "" {
		query += ` ORDER BY r.id ASC`
	} else {
		query += ` ORDER BY ` + u.SortMap()[sort] + ` ASC`
	}
	return query
}

func nearQuery(near string, query string) string {
	if near != "" {
		s := strings.Split(near, ",")
		var radius = "100" // TODO: externalize
		query = andWhere(query)
		query += u.Lat + ` > ` + s[0] + `-` + radius + ` AND ` + u.Lat + ` < ` + s[0] + `+` + radius + ` AND ` + u.Lng + ` > ` + s[1] + `-` + radius + ` AND ` + u.Lng + ` < ` + s[1] + `+` + radius
	}

	return query
}
