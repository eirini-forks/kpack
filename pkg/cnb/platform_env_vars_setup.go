package cnb

import (
	"os"
	"strings"
)

func SetupPlatformEnvVars(dir, envVarsJSON string) error {
	var envVars []envVariable

	for _, v := range os.Environ() {
		parts := strings.Split(v, "=")
		if !strings.HasPrefix(parts[0], "APP_ENV_VAR_") {
			continue
		}
		envVars = append(envVars, envVariable{
			Name:  strings.Replace(parts[0], "APP_ENV_VAR_", "", 1),
			Value: parts[1],
		})
	}

	return serializeEnvVars(envVars, dir)
}
