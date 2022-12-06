/*
Copyright © 2022 Ron Lynn <dad@lynntribe.net>
*/
package szstreams

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/roncewind/szrecord"
)

type RabbitmqStream struct {
	connectionUrl string
	exchange      string
	inputQ        string
	outputQ       string
	inputChannel    chan Commitable
	outputChannel   chan Commitable
}

type RabbitmqCommitable struct {
	record   szrecord.Record
	delivery amqp.Delivery
}

// ----------------------------------------------------------------------------
func (stream RabbitmqStream) GetInputChannel(ctx context.Context) (chan Commitable, error) {
	return stream.inputChannel, nil
}

// ----------------------------------------------------------------------------
func (stream RabbitmqStream) GetOutputChannel(ctx context.Context) (chan Commitable, error) {
	return stream.outputChannel, nil
}

// ----------------------------------------------------------------------------
func (commitable RabbitmqCommitable) Commit(ctx context.Context) {
	commitable.delivery.Ack(false)
}

// ----------------------------------------------------------------------------
func (commitable RabbitmqCommitable) Discard(ctx context.Context) {
	commitable.delivery.Nack(false, false)
}

// ----------------------------------------------------------------------------
func (c RabbitmqCommitable) Retry(ctx context.Context) {
	c.delivery.Nack(false, true)
}
