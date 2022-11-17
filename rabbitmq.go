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
	readStream    chan Commitable
	writeStream   chan Commitable
}

type RabbitmqCommitable struct {
	record   szrecord.Record
	delivery amqp.Delivery
}

// ----------------------------------------------------------------------------
func (stream RabbitmqStream) GetReadStream(ctx context.Context) (<-chan Commitable, error) {
	return stream.readStream, nil
}

// ----------------------------------------------------------------------------
func (stream RabbitmqStream) GetWriteStream(ctx context.Context) (chan<- Commitable, error) {
	return stream.writeStream, nil
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
