package api

import (
	"fmt"
	"net/http"

	socket "github/It-Alex/go-rfxcom-command/internal/socket"

	"github.com/labstack/echo/v4"
	"go.bug.st/serial.v1"
)

func controlShutters(c echo.Context) error {
	u := new(struct {
		Name    string `json:"name" form:"name" query:"name"`
		Command string `json:"command" form:"command" query:"command"`
	})
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
	case "salon":
		command[4] = 0x0e
		command[5] = 0xf2
		command[6] = 0x98
	case "maison":
		command[3] = 0x0a
		command[4] = 0x0e
		command[5] = 0xe8
		command[6] = 0x79
	case "hugo":
		command[3] = 0x06
		command[4] = 0x11
		command[5] = 0xf7
		command[6] = 0xdb
	case "alex-door":
		command[4] = 0x0e
		command[5] = 0xe9
		command[6] = 0x78
	case "alex":
		command[3] = 0x08
		command[4] = 0x0e
		command[5] = 0xee
		command[6] = 0xca
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

	_, err := socket.Port.Write(command)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, u)
}

func showTTY(c echo.Context) error {
	ports, err := serial.GetPortsList()
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, ports)
}

func changePort(c echo.Context) error {
	body := new(struct {
		PortName string `json:"name" form:"name" query:"name"`
	})
	if err := c.Bind(body); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, nil)
	}

	err := socket.ChangePort(body.PortName)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, nil)
}
