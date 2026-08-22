package configs

import (
	"fmt"
	"os"
)

type databaseConfig struct {
	Host     string `toml:"db_host"`
	Port     string `toml:"db_user"`
	User     string `toml:"db_user"`
	Password string `toml:"db_password"`
	Name     string `toml:"db_name"`
}

func loadDatabaseConfig() databaseConfig {
	return databaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}
}

func GetDbConfig() databaseConfig {
	return databaseConfig{
		Host:     globalConfig.Database.Host,
		Port:     globalConfig.Database.Port,
		User:     globalConfig.Database.User,
		Password: globalConfig.Database.Password,
		Name:     globalConfig.Database.Name,
	}
}

func getDbString() string {
	conf := GetDbConfig()

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", conf.Host, conf.User, conf.Password, conf.Name, conf.Port)
}
