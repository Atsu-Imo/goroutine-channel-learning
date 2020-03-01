package server

import (
	"fmt"
	"time"
)

// Listener クライアントからの接続を待ち続ける ダミー
type Listener struct {
}

func NewListener() *Listener {
	return &Listener{}
}

func (l *Listener) Listen() {
	t := time.NewTicker(time.Second)
	for {
		select {
		case <-t.C:
			fmt.Println("Listening...")
		}
	}
}
