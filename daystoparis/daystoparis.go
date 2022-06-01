package daystoparis

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

type appEnv struct {
	envConfig *viper.Viper
}

func CLI(args []string) int {
	var app appEnv

	err := app.fromArgs(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return errorBuildingAppConfigFromArgs
	}

	if err = app.run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)

		return errorRunningApplication
	}

	return 0
}

func (app *appEnv) run() error {
	now := time.Now()
	d := app.envConfig.Get("date")
	targetDate, _ := d.(time.Time)

	days := int(targetDate.Sub(now).Hours()/24) + 1

	switch {
	case days == 1:
		fmt.Printf("Llegó el día para '%s\n", app.envConfig.GetString("city"))
	case (days > 1) && (days <= app.envConfig.GetInt("min_days_to_count")):
		fmt.Printf("%s ... %d\n", app.envConfig.GetString("city"), days)
	}

	return nil
}

func readConfig(filename, configPath string, defaults map[string]interface{}) (*viper.Viper, error) {
	v := viper.New()

	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	v.SetConfigName(filename)
	v.AddConfigPath(configPath)
	v.SetConfigType("env")

	err := v.ReadInConfig()

	return v, err
}

func (app *appEnv) fromArgs(args []string) error {
	if len(args) != requiredNumberOfArgs {
		return fmt.Errorf("error: <City or Country> <date>")
	}

	envConfig, err := readConfig("daystoparis.env", os.Getenv("HOME"), map[string]interface{}{
		"min_days_to_count": 30,
	})
	if err != nil {
		return err
	}

	app.envConfig = envConfig
	app.envConfig.Set("program_name", args[0])
	app.envConfig.Set("city", args[1])

	date, err := time.Parse(dateFormat, args[2])
	if err != nil {
		return err
	}
	app.envConfig.Set("date", date)

	return nil
}
