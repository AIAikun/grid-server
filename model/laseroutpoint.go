package model

type (
	LaserOutPoint struct {
		Id           int     `json:"id"`
		Lng          float64 `json:"lng"`
		Lat          float64 `json:"lat"`
		Laser_out_id int     `json:"laserOutId"`
	}
)
