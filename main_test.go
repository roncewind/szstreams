/*
Copyright © 2022 Ron Lynn <dad@lynntribe.net>
*/
package szstreams

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"testing"
	"time"
)

const taskCount = 500000

type Worker struct { //Name?
	Input       chan Commitable
	Output      chan Commitable
	Execute OperationFunction
}

// ----------------------------------------------------------------------------
func (worker Worker) GetInputChannel(context.Context) (chan Commitable, error) {
	return worker.Input, nil
}

// ----------------------------------------------------------------------------
func (worker Worker) GetOutputChannel(context.Context) (chan Commitable, error) {
	return worker.Output, nil
}

// ----------------------------------------------------------------------------
type testCommitable struct {
	payload string
}

// ----------------------------------------------------------------------------
func (c testCommitable) GetPayload(context.Context) interface{} {
	return c.payload
}

// ----------------------------------------------------------------------------
func (c testCommitable) Commit(context.Context) {
	//do nothing
}

// ----------------------------------------------------------------------------
func (c testCommitable) Discard(context.Context) {
	//do nothing
}

// ----------------------------------------------------------------------------
func (c testCommitable) Retry(context.Context) {
	//do nothing
}

// ----------------------------------------------------------------------------
func sourceOperator(worker *Worker, ctx context.Context) (bool, error) {
	for i := 0; i < taskCount; i++ {
		c := &testCommitable{
			payload: fmt.Sprintf("%d", i),
		}
		worker.Output <- c
		// fmt.Println("src:", i, len(worker.Output))
	}
	fmt.Println("Source added: ", taskCount)
	return true, nil
}

// ----------------------------------------------------------------------------
// no-op operator
// do nothing but pass the Commitable from input chan to output chan
func noopOperator(worker *Worker, ctx context.Context) (bool, error) {
	noopCount := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("noopCount = ", noopCount)
			return true, nil
		case c := <-worker.Input:
			worker.Output <- c
			noopCount++
			// fmt.Println("noopCount = ", noopCount)
			// fmt.Println("noop: ", c.GetPayload(ctx), len(worker.Input), len(worker.Output))
		}
	}
}

// ----------------------------------------------------------------------------
func sinkOperator(worker *Worker, ctx context.Context) (bool, error) {
	commitCount := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("commitCount = ", commitCount)
			return true, nil
		case c := <-worker.Input:
			c.Commit(ctx)
			commitCount++
			// fmt.Println("sink: ", c.GetPayload(ctx), len(worker.Input))
			if worker.Output != nil {
				worker.Output <- c
			}
		}
	}
}

// ----------------------------------------------------------------------------
func TestSomething(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	// forever := make(chan bool)
	source := &Worker{
		Input: nil,
		Output: make(chan Commitable, 100),
		Execute: sourceOperator,
	}
	go source.Execute(source, ctx)
	go source.Execute(source, ctx)
	sourceOutputChannel, _ := source.GetOutputChannel(ctx)
	worker := &Worker{
		Input:       sourceOutputChannel,
		Output:      make(chan Commitable, 100),
		Execute: noopOperator,
	}
	go worker.Execute(worker, ctx)
	go worker.Execute(worker, ctx)
	go worker.Execute(worker, ctx)
	go worker.Execute(worker, ctx)
	workerOutputChannnel, _ := worker.GetOutputChannel(ctx)
	// worker2 := &Worker{
	// 	Input:       workerOutputChannnel,
	// 	Output:      make(chan Commitable, 20),
	// 	Execute: noopOperator,
	// }
	// go worker2.Execute(worker2, ctx)
	// go worker2.Execute(worker2, ctx)
	// go worker2.Execute(worker2, ctx)
	// go worker2.Execute(worker2, ctx)
	// worker2OutputChannnel, _ := worker.GetOutputChannel(ctx)
	sink := &Worker{
		Input: workerOutputChannnel,
		Output: nil,
		Execute: sinkOperator,
	}
	// go sink.Execute(sink, ctx)
	go sink.Execute(sink, ctx)
	go func() {
		sigchan := make(chan os.Signal)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		// total := len(sourceOutputChannel) + len(workerOutputChannnel) + len(worker2OutputChannnel) + len(sinkOut)
		// fmt.Println("total = ", total)
		cancel()
		fmt.Println("cancel called, ctrl-c to kill the program")
		signal.Stop(sigchan)
	}()
	// give it a moment to start up
	time.Sleep(5 * 1000 * time.Millisecond)
	for {
		sourceChannelCount := len(sourceOutputChannel)
		if sourceChannelCount == 0 {
			fmt.Println("Source channel emptied")
			workerChannnelCount := len(workerOutputChannnel)
			if workerChannnelCount == 0 {
				fmt.Println("Worker channel emptied")
				cancel()
				// forever <- true
				break
			}
		}
	}
	// <- forever
}

