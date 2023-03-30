package utils

import (
	"os"
)

var required_envs = []string{
	"REST_PORT",
	"REDIS_ADDRESS",
	"REDIS_PASSWORD",
}

func CheckEnvs() []string {
	missing_envs := []string{}

	for _, env := range required_envs {
		if os.Getenv(env) == "" {
			missing_envs = append(missing_envs, env)
		}
	}

	return missing_envs
}
