package model

type (
	UserLogin struct {
		Username string `json:"username"`
		PassWord string `json:"password"`
		Code     string `json:"code"`
		Uuid     string `json:"uuid"`
	}
)
