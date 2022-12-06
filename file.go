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
	inputChannel    chan Commitable
	outputChannel   chan Commitable
}

type FileCommitable struct {
	record   szrecord.Record
}


// ----------------------------------------------------------------------------
func (stream FileStream) GetInputChannel(ctx context.Context) (chan Commitable, error) {
	return stream.inputChannel, nil
}

// ----------------------------------------------------------------------------
func (stream FileStream) GetOutputChannel(ctx context.Context) (chan Commitable, error) {
	return stream.outputChannel, nil
}

// ----------------------------------------------------------------------------
func (commitable FileCommitable) Commit(ctx context.Context) {
	return
}

// ----------------------------------------------------------------------------
func (commitable FileCommitable) Discard(ctx context.Context) {
	return
}

// ----------------------------------------------------------------------------
func (c FileCommitable) Retry(ctx context.Context) {
	return
}

// // ----------------------------------------------------------------------------
// func NewReadFileStream(filename string) (FileStream, error) {
// 	// create a file to read from
// }

// // ----------------------------------------------------------------------------
// func NewWriteFileStream(filename string) (FileStream, error) {
// 	// create a file to write to
// }