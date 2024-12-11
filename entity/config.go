package entity

import "time"

type Config struct {
	API struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"api"`
	Auth struct {
		JwtSecretKey string `mapstructure:"jwt_secret_key"`
	} `mapstructure:"auth"`
	Context struct {
		Timeout time.Duration `mapstructure:"timeout"`
	} `mapstructure:"context"`
	Database struct {
		Mysql Mysql `mapstructure:"mysql"`
	} `mapstructure:"database"`
}

type Mysql struct {
	Host     string `mapstructure:"host"`
	Dbname   string `mapstructure:"dbname"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}
