package natsserver

import (
	"fmt"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

var Manager *NatsManager = nil

func SetupGlobalNatsManager() {
	nm, err := newNatsServer()
	if err != nil {
		panic(err)
	}

	Manager = nm
}

type NatsManager struct {
	S *server.Server
	C *nats.Conn
}

func (nm NatsManager) Close() {
	nm.C.Close()
}

func newNatsServer() (*NatsManager, error) {
	ns, err := server.NewServer(&server.Options{
		Host: "0.0.0.0",
		Port: 4222,
		Websocket: server.WebsocketOpts{
			Host:  "0.0.0.0",
			Port:  8080,
			NoTLS: true,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error starting nats server: %w", err)
	}

	go ns.Start()
	if !ns.ReadyForConnections(4 * time.Second) {
		return nil, fmt.Errorf("error ns not ready for connections")
	}

	nc, err := nats.Connect(ns.ClientURL())
	if err != nil {
		return nil, fmt.Errorf("error creating connection to nats server: %w", err)
	}

	return &NatsManager{
		S: ns,
		C: nc,
	}, nil

}
