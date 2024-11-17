package config

import "os"

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_DATABASE string
	NODE_ENV    string
	JWT_SECRET  string
	S3_ACCESS_KEY string
	S3_SECRET_KEY string
	S3_ENDPOINT string
}

var Conf Config

func Run() {
	Conf = Config{
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_DATABASE: os.Getenv("DB_DATABASE"),
		NODE_ENV:    os.Getenv("NODE_ENV"),
		JWT_SECRET:  os.Getenv("JWT_SECRET"),
		S3_ACCESS_KEY: os.Getenv("S3_ACCESS_KEY"),
		S3_SECRET_KEY: os.Getenv("S3_SECRET_KEY"),
		S3_ENDPOINT: os.Getenv("S3_ENDPOINT"),
	}
}
