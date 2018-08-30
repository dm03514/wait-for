package net

import (
	"net"

	"time"

	log "github.com/sirupsen/logrus"
)

type Net struct {
	Network string
	Address string
}

func New(network string, addr string) (Net, error) {
	return Net{
		Network: network,
		Address: addr,
	}, nil
}

func (n Net) CheckReady() (ready bool, err error) {
	conn, err := net.Dial(n.Network, n.Address)
	if err != nil {
		return false, err
	}

	log.WithFields(log.Fields{
		"error":       err,
		"local_addr":  conn.LocalAddr(),
		"remote_addr": conn.RemoteAddr(),
		"module":      "poller.Net",
	}).Debug("conn_response")

	defer func() {
		if err := conn.Close(); err != nil {
			ready = false
		}
	}()

	conn.SetReadDeadline(time.Now())
	if _, err := conn.Read([]byte{}); err != nil {
		return false, err
	}

	return true, nil
}
