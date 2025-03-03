package model

type (
	PolygonPoint struct {
		Id        int     `json:"id"`
		Lng       float64 `json:"lng"`
		Lat       float64 `json:"lat"`
		PolygonId int     `json:"polygon_id"`
	}
)
