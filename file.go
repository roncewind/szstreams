/*
Copyright © 2022 Ron Lynn <dad@lynntribe.net>
*/
package szstreams

import (
	"context"

	"github.com/roncewind/szrecord"
)

type FileStream struct {
	filename string
	readStream    chan Commitable
	writeStream   chan Commitable
}

// ----------------------------------------------------------------------------
func (stream FileStream) GetReadStream(ctx context.Context) (<-chan Commitable, error) {
	return stream.readStream, nil
}

// ----------------------------------------------------------------------------
func (stream FileStream) GetWriteStream(ctx context.Context) (chan<- Commitable, error) {
	return stream.writeStream, nil
}

// ----------------------------------------------------------------------------
func (commitable RabbitmqCommitable) Commit(ctx context.Context) {
	return
}

// ----------------------------------------------------------------------------
func (commitable RabbitmqCommitable) Discard(ctx context.Context) {
	return
}

// ----------------------------------------------------------------------------
func (c RabbitmqCommitable) Retry(ctx context.Context) {
	return
}

// ----------------------------------------------------------------------------
func NewReadFileStream(filename string) (FileStream, error) {
	// create a file to read from
}

// ----------------------------------------------------------------------------
func NewWriteFileStream(filename string) (FileStream, error) {
	// create a file to write to
}