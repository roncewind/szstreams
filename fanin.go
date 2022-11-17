/*
Copyright © 2022 Ron Lynn <dad@lynntribe.net>
*/
package szstreams

import (
	"context"
	"sync"
)

func FanIn(ctx context.Context, streams ...<-chan interface{}) <-chan interface{} {
	joinedStream := make(chan interface{})

	var wg sync.WaitGroup
	wg.Add(len(streams))

	for _, s := range streams {
		s := s

		go func() {
			defer wg.Done()
			for {
				select {
				case stream := <-s:
					joinedStream <- stream
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(joinedStream)
	}()
	return joinedStream
}
