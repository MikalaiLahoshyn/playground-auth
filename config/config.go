package configs

import (
	"github.com/spf13/viper"
)

const (
	configType = "yaml"
)

type PostgresDatabase struct {
	Driver   string `mapstructure:"DB_DRIVER"`
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
}

type Config struct {
	PostgresDb PostgresDatabase `mapstructure:"DB_DRIVER"`
}

func LoadConfig(path string) (Config, error) {
	var (
		postgresDb PostgresDatabase
	)

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType(configType)

	// override variables with system env if required
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	if err = viper.Unmarshal(&postgresDb); err != nil {
		return Config{}, err
	}

	config := Config{
		PostgresDb: postgresDb,
	}

	if err = viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}