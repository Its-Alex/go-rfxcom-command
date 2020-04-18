package main

import (
	api "github/It-Alex/go-rfxcom-command/internal/api"
	socket "github/It-Alex/go-rfxcom-command/internal/socket"
	"time"

	"github.com/spf13/viper"
	"go.bug.st/serial.v1"
)

func init() {
	viper.SetEnvPrefix("go_rfxcom")

	viper.BindEnv("addr")
	viper.BindEnv("port")

	viper.SetDefault("addr", "0.0.0.0")
	viper.SetDefault("port", "1323")
}

func main() {
	err := socket.InitSocket(&serial.Mode{BaudRate: 38400})
	if err != nil {
		panic(err)
	}

	// Waiting for RFXCom initalization
	time.Sleep(time.Millisecond * 1000)

	api.Launch()
}
