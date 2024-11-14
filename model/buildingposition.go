package model

type (
	BuidingRow struct {
		Id       int     `json:"id"`
		Model_id int     `json:"modelId"`
		Status   int     `json:"status"`
		Lng      float64 `json:"lng"`
		Lat      float64 `json:"lat"`
		Angle    float64 `json:"angle"`
	}
)
