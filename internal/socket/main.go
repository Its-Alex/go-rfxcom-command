package socket

import (
	"io"
	"strings"
	"time"

	serial "go.bug.st/serial.v1"
)

// Port is the current connected socket
var Port serial.Port

// SerialConfig is the used configurationf for port
var SerialConfig *serial.Mode

// InitSocket init socket connection
func InitSocket(config *serial.Mode) error {
	SerialConfig = config

	var USBPort string
	ports, err := serial.GetPortsList()
	if err != nil {
		panic(err)
	}
	if len(ports) == 0 {
		panic("No serial ports found!")
	}

	for _, port := range ports {
		if strings.Contains(port, "USB") {
			USBPort = port
		}
	}

	Port, err = serial.Open(USBPort, SerialConfig)
	if err != nil {
		return err
	}

	SendReset()

	return nil
}

// ChangePort used to allow switching to another TTY
func ChangePort(portName string) (err error) {
	err = Port.Close()
	if err != nil {
		return err
	}

	Port, err = serial.Open(portName, SerialConfig)
	if err != nil {
		return err
	}

	SendReset()

	return err
}

// SendReset send a reset command to rfxcom
func SendReset() error {
	_, err := Port.Write([]byte{0x0d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	if err != nil {
		return err
	}
	return nil
}

// SetMode set rfxcom mode
func SetMode(enableBlindsTx bool) error {
	var b []byte
	if enableBlindsTx {
		// Enable blinds protocol
		b = []byte{0x0d, 0x00, 0x00, 0x03, 0x03, 0x54, 0x1c, 0x00, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00}
	} else {
		// Disable all protocol
		b = []byte{0x0d, 0x00, 0x00, 0x01, 0x03, 0x53, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	}

	_, err := Port.Write(b)
	if err != nil {
		return err
	}
	return nil
}

// Read read from socket
func Read() ([]byte, error) {
	buf := make([]byte, 257)
	for {
		// read length
		i, err := Port.Read(buf[0:1])
		if i == 0 && err == io.EOF {
			// empty read, sleep a bit recheck
			time.Sleep(time.Millisecond * 200)
			continue
		}
		if err != nil {
			return nil, err
		}
		if i == 0 {
			continue
		}

		// read rest of data
		l := int(buf[0])
		buf = buf[0 : l+1]
		for read := 0; read < l; read += i {
			i, err = Port.Read(buf[read+1:])
			if i == 0 && err == io.EOF {
				time.Sleep(time.Millisecond * 200)
				continue
			}
			if err != nil {
				return nil, err
			}
		}
		return buf, err
	}
}
