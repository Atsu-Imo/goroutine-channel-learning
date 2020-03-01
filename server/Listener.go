package server

import (
	"context"
	"fmt"
	"time"
)

// Listener クライアントからの接続を待ち続ける ダミー
type Listener struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewListener(ctx context.Context) *Listener {
	listenerCtx, cancelFunc := context.WithCancel(ctx)
	return &Listener{
		ctx:        listenerCtx,
		cancelFunc: cancelFunc,
	}
}

func (l *Listener) Listen() {
	invoke := time.NewTicker(2 * time.Second)
	timeout := time.NewTimer(7 * time.Second)
	for {
		select {
		case <-invoke.C:
			fmt.Println("process started...")
		case <-timeout.C:
			fmt.Println("shutdown")
			l.cancelFunc()
			return
		}
	}
}
