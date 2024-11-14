package model

type (
	CircuitState struct {
		Id               int64   `json:"id"`
		Normal           bool    `json:"normal"`
		FaultProbability float64 `json:"faultProbability"`
	}
	CableRow struct {
		Id                int `json:"id"`
		Point_id1         int `json:"pointId1"`
		Point_id2         int `json:"pointId2"`
		Join_point_index1 int `json:"joinPointIndex1"`
		Join_point_index2 int `json:"joinPointIndex2"`
		Circuit_id        int `json:"circuitId"`
	}
)
