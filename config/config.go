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
}

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load env variables:", err)
	}
	version := os.Getenv("VERSION")

	if version == "" {
		fmt.Println("Version is requried")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("service name is requried")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("Http port is requried")
		os.Exit(1)
	}
	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("Port Must be Number")
		os.Exit(1)
	}
	configurations = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(port),
	}
}

func GetConfig() Config {
	LoadConfig()
	return configurations
}
