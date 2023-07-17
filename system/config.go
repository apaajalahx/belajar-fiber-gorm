package system

import (
	"encoding/json"
	"os"
	"strconv"
)

type ConfigInterface struct {
	DatabaseUrl string `json:"databaseUrl"`
	Port        int    `json:"port"`
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

	return nil

}
