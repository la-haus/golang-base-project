package config_test

import (
	"os"
	"testing"

	"github.com/la-haus/sample/config"
	"github.com/la-haus/sample/enums"
	"github.com/stretchr/testify/assert"
)

func TestEnvironments_whenAllVarsAreSet(t *testing.T) {
	os.Setenv("DD_ENV", enums.LocalEnvironment)
	defer os.Unsetenv("DD_ENV")

	c := config.Environments()

	assert.Equal(t, c.DDEnv, enums.LocalEnvironment)
}

func TestEnvironments_whenSomeVarIsMissed(t *testing.T) {
	expectedError := "Error parsing environment vars &errors.errorString{s:\"required key SERVER_HOST missing value\"}"
	os.Clearenv()

	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, expectedError, r)
		}
	}()

	config.Environments()
}
