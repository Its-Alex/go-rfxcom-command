package socket

import (
	"io"
	"time"

	"github.com/tarm/serial"
)

type Socket struct {
	Conn         *serial.Port
	SerialConfig *serial.Config
}

func InitSocket(config *serial.Config) (*Socket, error) {
	var err error
	newSocket := &Socket{
		SerialConfig: config,
	}

	newSocket.Conn, err = serial.OpenPort(newSocket.SerialConfig)
	if err != nil {
		return newSocket, err
	}
	newSocket.SendReset()
	return newSocket, nil
}

func (s *Socket) SendReset() error {
	_, err := s.Conn.Write([]byte{0x0d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	if err != nil {
		return err
	}
	return nil
}

func (s *Socket) SetMode(enableBlindsTx bool) error {
	b := []byte{0x0d, 0x00, 0x00, 0x01, 0x03, 0x53, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	_, err := s.Conn.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (s *Socket) Read() ([]byte, error) {
	buf := make([]byte, 257)
	for {
		// read length
		i, err := s.Conn.Read(buf[0:1])
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
			i, err = s.Conn.Read(buf[read+1:])
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
