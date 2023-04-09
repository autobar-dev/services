package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port int
}

func LoadConfig() (*Config, error) {
	error_list := []error{}

	port_string := os.Getenv("PORT")
	port, err := strconv.Atoi(port_string)
	if err != nil {
		error_list = append(error_list, err)
	}

	if len(error_list) > 0 {
		return nil, errors.New(fmt.Sprintf("Env load errors: %+v", error_list))
	}

	return &Config{
		Port: port,
	}, nil
}
