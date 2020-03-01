package server

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Listener クライアントからの接続を待ち続ける ダミー
type Listener struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	Wg         sync.WaitGroup
}

func NewListener(ctx context.Context) *Listener {
	listenerCtx, cancelFunc := context.WithCancel(ctx)
	return &Listener{
		ctx:        listenerCtx,
		cancelFunc: cancelFunc,
	}
}

func (l *Listener) Listen() {
	invoke := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-invoke.C:
			l.Wg.Add(1)
			fmt.Println("wg count++")
			conn := NewConn(l.ctx)
			go conn.handleConnection()
		case <-l.ctx.Done():
			return
		}
	}
}
func (l *Listener) Shutdown() {
	l.cancelFunc()
	fmt.Println("waiting for children")
	l.Wg.Wait()
}
