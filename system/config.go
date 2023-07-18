package system

import (
	"encoding/json"
	"os"
	"strconv"
)

type ConfigInterface struct {
	DatabaseUrl string `json:"databaseUrl"`
	Port        int    `json:"port"`
	JwtSecret   string `json:"JwtSecret"`
	AppName     string `json:"AppName"`
}

func LoadConfig(configFile string) error {

	dat, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}

	config := &ConfigInterface{}
	if err := json.Unmarshal([]byte(dat), config); err != nil {
		return err
	}

	os.Setenv("DATABASE_URL", config.DatabaseUrl)
	os.Setenv("PORT", strconv.Itoa(config.Port))
	os.Setenv("JWT_SECRET", config.JwtSecret)
	os.Setenv("APP_NAME", config.AppName)

	return nil

}
