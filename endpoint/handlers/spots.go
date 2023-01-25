package handlers

import (
	"endpoint/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"endpoint/models"
	"encoding/json"
)

type params = models.Params;
type spot = models.Spot

func generateQuery(p *params) string{

	longitude := fmt.Sprintf("%f", p.Longitude)
	latitude := fmt.Sprintf("%f", p.Latitude)
	
	if p.Type == "circle" {
		return "SELECT * " +
		"FROM (" +
		"SELECT *, ST_DistanceSphere(ST_Centroid(coordinates::geometry), ST_GeomFromText('POINT(" + longitude + " " + latitude + ")')) as distance " +
		"FROM \"MY_TABLE\" " +
		"WHERE ST_DistanceSphere(ST_Centroid(coordinates::geometry), ST_GeomFromText('POINT(" + longitude + " " + latitude + ")')) <= $1 " +
		") sub " +
		"ORDER BY CASE WHEN distance < 50 THEN rating END ASC, CASE WHEN distance >= 50 THEN distance END ASC;"
	}
	return "SELECT * from \"MY_TABLE\" WHERE " +
		"ST_Within(" +
		"ST_GeomFromText(ST_AsText(coordinates)), " +
		"ST_MakeEnvelope(" +
			"ST_X(ST_Transform(ST_SetSRID(ST_Point(" + longitude + ", " + latitude + "), 4326), 3857)) - $1, " +
			"ST_Y(ST_Transform(ST_SetSRID(ST_Point(" + longitude + ", " + latitude + "), 4326), 3857)) - $1, " +
			"ST_X(ST_Transform(ST_SetSRID(ST_Point(" + longitude + ", " + latitude + "), 4326), 3857)) + $1, " +
			"ST_Y(ST_Transform(ST_SetSRID(ST_Point(" + longitude + ", " + latitude + "), 4326), 3857)) + $1" +
			 "));";
}

func GetSpots(c *fiber.Ctx) error {

	tmp := c.Locals("spot")
	spotParam := tmp.(*params)
	query := generateQuery(spotParam)
	database.DB.Ping()
	rows, err := database.DB.Query(query, spotParam.Radius)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()
	var spots []spot
	for rows.Next() {
        var sp spot
        if err := rows.Scan(&sp.Id, &sp.Name, &sp.Website,
            &sp.Coordinates, &sp.Description, &sp.Rating, &sp.Distance); err != nil {
            return err
        }
        spots = append(spots, sp)
    }
	jsonSpots, err := json.Marshal(spots)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	c.Send(jsonSpots)
	return nil
}