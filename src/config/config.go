package config

import "encoding/json" //needed for turning structs into redisable JSON

type Config struct {
	Database DatabaseServer `json:"database"`
}

type DatabaseServer struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

// /Our Application Config Variable [EXPORTABLE]
var AppConfig *Config

// Simple function to handle unmarshaling the Application Config
func HandleConfig(body []byte) (*Config, error) {
	cfg := new(Config)
	err := json.Unmarshal(body, &cfg)
	//store for use later
	AppConfig = cfg
	return cfg, err

}

func GetConfig() (config *Config) {

	return AppConfig
}
