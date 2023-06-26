package config

import "os"

type ServiceParameters struct {
	StorageType string
	GRPCPort    string
	HTTPPort    string
	LogLevel    string
}

type DBParameters struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func InitializeServiceParametersFromEnv() ServiceParameters {
	return ServiceParameters{
		StorageType: os.Getenv("STORAGE_TYPE"),
		GRPCPort:    os.Getenv("GRPC_PORT"),
		HTTPPort:    os.Getenv("HTTP_PORT"),
		LogLevel:    os.Getenv("LOG_LEVEL"),
	}
}

func InitializeDBParametersFromEnv() DBParameters {
	return DBParameters{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}
}
