package config

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Env struct {
	AppEnv     string `mapstructure:"APP_ENV"`
	Port       string `mapstructure:"PORT"`
	Domain     string `mapstructure:"DOMAIN"`
	AdminToken string `mapstructure:"ADMIN_TOKEN"`

	DBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBName string `mapstructure:"DB_NAME"`

	ApiURL        string        `mapstructure:"API_URL"`
	JwtSecreet    string        `mapstructure:"JWT_SECRET"`
	JwtAccessTTL  time.Duration `mapstructure:"JWT_ACCESS_TTL"`
	JwtRefreshTTL time.Duration `mapstructure:"JWT_REFRESH_TTL"`
}

func NewEnv() Env {
	env := Env{}

	_, err := os.Stat(".env")
	useEnvFile := !os.IsNotExist(err)

	if useEnvFile {
		viper.SetConfigType("env")
		viper.SetConfigName(".env")
		viper.AddConfigPath(".")

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal("Can't read the .env file: ", err)
		}

		err = viper.Unmarshal(&env)
		if err != nil {
			log.Fatal("Environment can't be loaded: ", err)
		}
	} else {
		env.bindEnv()
	}

	if env.AppEnv != "production" {
		log.Println("The App is running in development env")
	}

	return env
}

func (e *Env) bindEnv() {
	e.ApiURL = os.Getenv("API_URL")
	e.AppEnv = os.Getenv("APP_ENV")
	e.Port = os.Getenv("PORT")
	e.Domain = os.Getenv("DOMAIN")

	e.DBHost = os.Getenv("DB_HOST")
	e.DBPort = os.Getenv("DB_PORT")
	e.DBUser = os.Getenv("DB_USER")
	e.DBPass = os.Getenv("DB_PASS")
	e.DBName = os.Getenv("DB_NAME")

	e.JwtSecreet = os.Getenv("JWT_SECRET")

	if val := os.Getenv("JWT_ACCESS_TTL"); val != "" {
		d, err := time.ParseDuration(val)
		if err != nil {
			log.Fatalf("Invalid JWT_ACCESS_TTL format: %v", err)
		}
		e.JwtAccessTTL = d
	}

	if val := os.Getenv("JWT_REFRESH_TTL"); val != "" {
		d, err := time.ParseDuration(val)
		if err != nil {
			log.Fatalf("Invalid JWT_REFRESH_TTL format: %v", err)
		}
		e.JwtRefreshTTL = d
	}
}

var Module = fx.Options(
	fx.Provide(NewEnv),
)
