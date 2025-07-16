package natsserver

import (
	"fmt"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

var GlobalNatsManager *NatsManager = nil

func SetupGlobalNatsManager() {
	nm, err := NewNatsServer()
	if err != nil {
		panic(err)
	}

	GlobalNatsManager = nm
}

type NatsManager struct {
	S *server.Server
	C *nats.Conn
}

func (nm NatsManager) Close() {
	nm.C.Close()
}

func NewNatsServer() (*NatsManager, error) {
	ns, err := server.NewServer(&server.Options{
		Host: "0.0.0.0",
		Port: 4222,
		Websocket: server.WebsocketOpts{
			Host: "0.0.0.0",
			Port: 8080,
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
