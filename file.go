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
	record   szrecord.Record
}

func (s FileStream) GetReadStream(ctx context.Context) (<-chan Commitable, error) {
	return make(chan Commitable), nil
}

func (s FileStream) GetWriteStream(ctx context.Context) (chan<- Commitable, error) {
	return make(chan Commitable), nil
}
