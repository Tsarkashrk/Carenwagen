package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func ConnectNATS(url string) (*nats.Conn, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		logrus.Fatal("Failed to connect to NATS: ", err)
		return nil, err
	}
	return nc, nil
}

func Publish(nc *nats.Conn, subject string, message []byte) error {
	err := nc.Publish(subject, message)
	if err != nil {
		logrus.Error("Failed to publish message: ", err)
		return err
	}
	return nil
}
