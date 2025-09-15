package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	DbHost      string
	DbPort      int
	DbUser      string
	DbPassword  string
	DbName      string
}

func LoadConfig() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Failed to the env variables:", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")

	if serviceName == "" {
		fmt.Println("Service name is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")

	if httpPort == "" {
		fmt.Println("Http port is required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)

	if err != nil {
		fmt.Println("Port must be number")
		os.Exit(1)

	}

	dbHost := os.Getenv("DB_HOST")

	if dbHost == "" {
		fmt.Println("Database Host is required")
		os.Exit(1)
	}

	dbPort := os.Getenv("DB_PORT")

	if dbPort == "" {
		fmt.Println("Database port is required")
		os.Exit(1)
	}

	db_port, err := strconv.ParseInt(dbPort, 10, 64)

	if err != nil {
		fmt.Println("Database Port must be number")
		os.Exit(1)

	}

	dbUser := os.Getenv("DB_USER")

	if dbUser == "" {
		fmt.Println("User is required")
		os.Exit(1)
	}

	dbPassword := os.Getenv("DB_PASSWORD")

	if dbPassword == "" {
		fmt.Println(" Database Password is required")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")

	if dbName == "" {
		fmt.Println("Database Name is required")
		os.Exit(1)
	}

	configurations = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(port),
		DbHost:      dbHost,
		DbPort:      int(db_port),
		DbUser:      dbUser,
		DbPassword:  dbPassword,
		DbName:      dbName,
	}
}

func GetConfig() Config {
	LoadConfig()
	return configurations
}
