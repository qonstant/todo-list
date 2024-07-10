package util

import (
	"github.com/spf13/viper"
)

// Configurations for the application
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)  // Add path where config file resides
	viper.SetConfigName("app") // Name of config file (without extension)
	viper.SetConfigType("env") // Type of config file (env file)

	viper.AutomaticEnv() // Read environment variables

	err = viper.ReadInConfig() // Read the config file
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config) // Unmarshal config values into struct
	return
}
