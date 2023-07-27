package config

import (
        "encoding/json"
        "fmt"
        "io/ioutil"
        "os"
)

// Config represents the configuration structure
type Config struct {
        DatabaseURL string `json:"database_url"`
        AppPort     int    `json:"app_port"`
}

// LoadConfig reads the configuration file and returns the Config struct
func LoadConfig(filename string) (*Config, error) {
        configFile, err := os.Open(filename)
        if err != nil {
                return nil, fmt.Errorf("failed to open config file: %w", err)
        }
        defer configFile.Close()

        byteValue, err := ioutil.ReadAll(configFile)
        if err != nil {
                return nil, fmt.Errorf("failed to read config file: %w", err)
        }

        var config Config
        err = json.Unmarshal(byteValue, &config)
        if err != nil {
                return nil, fmt.Errorf("failed to unmarshal config data: %w", err)
        }

        return &config, nil
}
