package main

import (
	"fmt"
	socket "github/It-Alex/go-rfxcom-command/socket"

	"github.com/tarm/serial"
)

func main() {
	s, err := socket.InitSocket(&serial.Config{Name: "/dev/serial/by-id/usb-RFXCOM_RFXtrx433_A129KO1K-if00-port0", Baud: 38400})
	if err != nil {
		panic(err)
	}
	s.SendReset()
	s.SetMode(true)

	var buf []byte
	for {
		buf, err = s.Read()
		if err != nil {
			panic(err)
		}

		fmt.Printf("Parse(%#v) = (%#v, %v)\n", buf, []string{}, err)
	}
}
