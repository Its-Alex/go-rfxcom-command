package main

import (
	api "github/It-Alex/go-rfxcom-command/internal/api"

	"github.com/spf13/viper"
)

func init() {
	viper.SetEnvPrefix("go_rfxom")

	viper.BindEnv("addr")
	viper.BindEnv("port")

	viper.SetDefault("addr", "0.0.0.0")
	viper.SetDefault("port", "1323")
}

func main() {
	api.Launch()
}
