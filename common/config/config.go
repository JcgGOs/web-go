package config

import (
	"os"
)

//AppConfig config for env
// conn, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/test?parseTime=true")
type AppConfig struct {
	DriverName string
	URL        string
}

//DefaultConfig default config
var DefaultConfig AppConfig

func (config *AppConfig) init() {
	config.DriverName = Env("db.driverName", "mysql")
	config.URL = Env("db.url", "root:mysql@tcp(127.0.0.1:3306)/test?parseTime=true")
}

//Env Read OS Env variable
func Env(key string, defaultValue string) string {
	v := os.Getenv(key)
	if len(v) == 0 {
		return defaultValue
	}
	return v
}
