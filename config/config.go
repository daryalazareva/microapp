package config

import "github.com/spf13/viper"

// All configuration variables of the application
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	//to tell viper to automatically override that it has read from config file with
	//the values of the environment variables if they exist
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	err = viper.Unmarshal(&config)

	return config, err
}
