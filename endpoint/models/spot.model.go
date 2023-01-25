package models

import "database/sql"

type Spot struct {
	Id			string `json:"id,omitempty"`
	Name		sql.NullString `json:"name,omitempty"` 
	Website		sql.NullString `json:"website,omitempty"`
	Coordinates	sql.NullString `json:"coordinates,omitempty"`
	Description sql.NullString `json:"description,omitempty"`
	Rating		float64
	Distance	float64
}