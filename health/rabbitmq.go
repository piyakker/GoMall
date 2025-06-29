package health

import (
	"context"

	config "github.com/piyakker/GoMall/configs"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQChecker struct {
	cfg *config.Config
}

func NewRabbitMQChecker(cfg *config.Config) *RabbitMQChecker {
	return &RabbitMQChecker{cfg: cfg}
}

func (r *RabbitMQChecker) Check(_ context.Context) error {
	conn, err := amqp091.Dial(r.cfg.RabbitMQUrl)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	ch.Close()

	return nil
}
