package model

type (
	LaserOut struct {
		Id        int     `json:"id"`
		CenterLng float64 `json:"centerLng"`
		CenterLat float64 `json:"centerLat"`
		Name      string  `json:"name"`
	}
)
