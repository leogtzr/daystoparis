package daystoparis

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type appEnv struct {
	envConfig *viper.Viper
}

func CLI(args []string) int {
	var app appEnv

	err := app.fromArgs(args)
	if err != nil {
		return errorBuildingAppConfigFromArgs
	}

	if err = app.run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)

		return errorRunningApplication
	}

	return 0
}

func (app *appEnv) run() error {

	return nil
}

func (app *appEnv) fromArgs(args []string) error {
	envConfig, err := readConfig("shellreminders.env", os.Getenv("HOME"), map[string]interface{}{
		// empty for now
	})
	if err != nil {
		return err
	}
	app.envConfig = envConfig

	return nil
}
