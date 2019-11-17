package main

import (
	"fmt"
	socket "github/It-Alex/go-rfxcom-command/socket"
	"net/http"
	"time"

	"github.com/labstack/echo"
	middleware "github.com/labstack/echo/middleware"
	"github.com/tarm/serial"
)

var s *socket.Socket

type RequestBody struct {
	Name    string `json:"name" form:"name" query:"name"`
	Command string `json:"command" form:"command" query:"command"`
}

func doAction(c echo.Context) error {
	u := new(RequestBody)
	if err := c.Bind(u); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, nil)
	}

	command := []byte{
		0x09, // type blinds1
		0x19, // type ?
		0x03, // type ?
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
	}

	switch u.Name {
	case "alexandre":
		command[4] = 0x0e
		command[5] = 0xef
		command[6] = 0x2a
	case "maison":
		command[4] = 0x0e
		command[5] = 0xf2
		command[6] = 0x98
	case "alex":
		command[4] = 0x0e
		command[5] = 0xe9
		command[6] = 0x78
	}

	switch u.Command {
	case "stop":
		command[8] = 0x02
		command[9] = 0x70
	case "down":
		command[8] = 0x01
		command[9] = 0x70
	case "up":
		command[8] = 0x00
		command[9] = 0x70
	}

	fmt.Println(u)
	fmt.Println(command)

	_, err := s.Conn.Write(command)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, u)
}

func main() {
	var err error
	s, err = socket.InitSocket(&serial.Config{Name: "/dev/serial/by-id/usb-RFXCOM_RFXtrx433_A129KO1K-if00-port0", Baud: 38400})
	if err != nil {
		panic(err)
	}
	s.SendReset()
	s.SetMode(true)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/action", doAction)

	// Waiting for RFXCom initalization
	time.Sleep(time.Millisecond * 1000)

	// Reading RFXCom output
	go func() {
		var buf []byte
		for {
			buf, err = s.Read()
			if err != nil {
				panic(err)
			}

			fmt.Printf("Parse(%#v) = (%#v, %v)\n", buf, []string{}, err)
		}
	}()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
