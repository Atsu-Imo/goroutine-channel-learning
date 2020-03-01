package server

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Conn クライアントとの接続を保つ、ダミー
type Conn struct {
	uuid      uuid.UUID
	parentCtx context.Context
	ctx       context.Context
	CancelFnc context.CancelFunc
	Wg        sync.WaitGroup
}

func NewConn(parent context.Context) *Conn {
	u, err := uuid.NewRandom()
	connCtx, fnc := context.WithCancel(parent)
	if err != nil {
		fmt.Println(err)
	}
	return &Conn{
		uuid:      u,
		parentCtx: parent,
		ctx:       connCtx,
		CancelFnc: fnc,
	}
}

//
func (c *Conn) handleConnection() {
	defer func() {
		fmt.Println("call cancel:", c.uuid)
		c.CancelFnc()
		fmt.Println("call done:", c.uuid)
		c.Wg.Done()
		fmt.Println("end:", c.uuid)
	}()
	fmt.Println("process started...: ", c.uuid)
	invoke := time.NewTicker(1 * time.Second)
	timeout := time.NewTimer(2 * time.Second)
	for {
		select {
		case <-invoke.C:
			fmt.Println("Conn is keeped: ", c.uuid)
		case <-c.parentCtx.Done():
			fmt.Println("Conn is closed: ", c.uuid)
			return
		case <-timeout.C:
			fmt.Println("Conn timeout: ", c.uuid)
			return
		case <-c.ctx.Done():
		default:
		}
	}
}
