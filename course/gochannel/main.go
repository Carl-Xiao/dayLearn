package main

import (
	"context"
	"fmt"
	"time"
)

type Tracker struct {
	ch   chan string
	stop chan struct{}
}

func NewTracker() *Tracker {
	return &Tracker{
		ch: make(chan string, 10),
	}
}
func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case t.ch <- data:
		return nil
	}
}

func (t *Tracker) Run() {
	//阻塞
	for data := range t.ch {
		time.Sleep(time.Second * 1)
		fmt.Println(data)
	}
	//处理完数据
	t.stop <- struct{}{}
}

func (t *Tracker) ShutDown(ctx context.Context) {
	close(t.ch)
	select {
	case <-t.stop:
		return
	case <-ctx.Done():
	}
}

func main() {
	tr := NewTracker()
	go tr.Run()
	tr.Event(context.Background(), "1")
	tr.Event(context.Background(), "2")
	tr.Event(context.Background(), "3")
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()
	tr.ShutDown(ctx)
}
