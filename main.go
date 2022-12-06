/*
Copyright © 2022 Ron Lynn <dad@lynntribe.net>
*/
package szstreams

import (
	"context"
)

type OperationFunction func(*Worker, context.Context) (bool, error)

type ChannelInput interface {
	GetInputChannel(context.Context) (chan Commitable, error)
}

type ChannelOutput interface {
	GetOutputChannel(context.Context) (chan Commitable, error)
}

type Commitable interface {
	GetPayload(context.Context) interface{}
	Commit(context.Context)
	Discard(context.Context)
	Retry(context.Context)
}

// func doASenzingSomething() {
// 	w := &Worker{
// 		In:       make(chan Commitable),
// 		Out:      make(chan Commitable),
// 		Operator: g2Operator,
// 	}
// 	w.Operator(w, context.TODO())

// }

// func g2Operator(worker *Worker, ctx context.Context) (bool, error) {
// 	g2engine, g2engineErr := getG2engine(context.TODO())
// 	if g2engineErr != nil {
// 		panic("no g2")
// 	}
// 	var c Commitable = <-worker.In
// 	record, _ := szrecord.NewRecord(fmt.Sprintf("%v",c.GetPayload(ctx)))
// 	_ = g2engine.AddRecord(ctx, record.DataSource, record.Id, record.Json, "")
// 	worker.Out <- c
// 	return true, nil
// }

// // ----------------------------------------------------------------------------
// func getG2engine(ctx context.Context) (g2engine.G2engine, error) {
// 	var err error = nil
// 	g2engine := g2engine.G2engineImpl{}

// 	moduleName := "flibberty-do"
// 	verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
// 	iniParams, jsonErr := g2configuration.BuildSimpleSystemConfigurationJson("")
// 	if jsonErr != nil {
// 		return &g2engine, jsonErr
// 	}

// 	err = g2engine.Init(ctx, moduleName, iniParams, verboseLogging)
// 	return &g2engine, err
// }

// do we even need this inteface?  could be we just add receivers to the struct
// impl?
// type Operator interface { //Name?  Processor/Process?  Worker/Do?
// 	Operate(context.Context) (bool, error)
// }

// pipe line working reads from the in chan, operates on the item then writes
//
//	to the out chan.  Workers are assembled like a linked list.
//	could have a fan-out pattern to currently work items?
//	need to have a terminal work that has no implemented Out, but calls one of
//	Commit/Discard/Retry
// type WorkerX struct { //Name?
// 	In  chan<- Commitable
// 	Out <-chan Commitable
// 	Operator
// }

// func (w WorkerX) Operate(ctx context.Context) (bool, error) {
// 	w.In //get a Commitable
// 	//workon Commitable
// 	w.Out //send the Commitable on
// }
