package configs

import (
	"log"

	"github.com/spf13/viper"
)

// Config is the struct representation of your config.yaml properties
type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Database struct {
		Driver   string `mapstructure:"driver"`
		Postgres struct {
			Host     string `mapstructure:"host"`
			Port     int    `mapstructure:"port"`
			User     string `mapstructure:"user"`
			Password string `mapstructure:"password"`
			Dbname   string `mapstructure:"dbname"`
			Sslmode  string `mapstructure:"sslmode"`
		} `mapstructure:"postgres"`
		Mysql struct {
			Host     string `mapstructure:"host"`
			Port     int    `mapstructure:"port"`
			User     string `mapstructure:"user"`
			Password string `mapstructure:"password"`
			Dbname   string `mapstructure:"dbname"`
		} `mapstructure:"mysql"`
		Sqlite struct {
			Dsn string `mapstructure:"dsn"`
		} `mapstructure:"sqlite"`
	} `mapstructure:"database"`

	Middleware struct {
		Logging  bool `mapstructure:"logging"`
		Recovery bool `mapstructure:"recovery"`
		Auth     bool `mapstructure:"auth"`
	} `mapstructure:"middleware"`

	Swagger struct {
		Enabled bool `mapstructure:"enabled"`
	} `mapstructure:"swagger"`
}

// LoadConfig reads configuration from the config.yaml file
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")  // Config file name (without extension)
	viper.SetConfigType("yaml")    // Config file type
	viper.AddConfigPath("configs") // Path where config.yaml is located

	var config Config
	if err := viper.ReadInConfig(); err != nil { // Read config
		log.Printf("Error reading config file: %v", err)
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil { // Unmarshal into struct
		log.Printf("Error unmarshalling config: %v", err)
		return nil, err
	}

	return &config, nil
}
