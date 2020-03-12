package socket

import (
	"io"
	"strings"
	"time"

	serial "go.bug.st/serial.v1"
)

type Socket struct {
	Port         serial.Port
	SerialConfig *serial.Mode
}

func InitSocket(config *serial.Mode) (*Socket, error) {
	var err error
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
	newSocket := &Socket{
		SerialConfig: config,
	}

	newSocket.Port, err = serial.Open(USBPort, config)
	if err != nil {
		return newSocket, err
	}
	newSocket.SendReset()
	return newSocket, nil
}

func (s *Socket) SendReset() error {
	_, err := s.Port.Write([]byte{0x0d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	if err != nil {
		return err
	}
	return nil
}

func (s *Socket) SetMode(enableBlindsTx bool) error {
	var b []byte
	if enableBlindsTx {
		// Enable blinds protocol
		b = []byte{0x0d, 0x00, 0x00, 0x03, 0x03, 0x54, 0x1c, 0x00, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00}
	} else {
		// Disable all protocol
		b = []byte{0x0d, 0x00, 0x00, 0x01, 0x03, 0x53, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	}

	_, err := s.Port.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (s *Socket) Read() ([]byte, error) {
	buf := make([]byte, 257)
	for {
		// read length
		i, err := s.Port.Read(buf[0:1])
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
			i, err = s.Port.Read(buf[read+1:])
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
