package model

type (
	Config struct {
		Mysql Mysql `mapstructure:"mysql"`
	}
	Mysql struct {
		Host     string
		Port     string
		Username string
		Password string
		Dbname   string
	}
)
