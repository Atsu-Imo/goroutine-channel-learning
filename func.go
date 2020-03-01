package main

import (
	"fmt"
)

// defaultがあるとchannelが受信していないときはそっちを実行する
func readChannel(chName string, ch <-chan string) bool {
	select {
	case <-ch:
		fmt.Println("done read", chName)
		return true
	default:
		return false
	}
}

// defaultがないとchannelに値が来るまで待機する
// chをロックする
func readChannelWithoutDef(chName string, ch <-chan string) bool {
	select {
	case <-ch:
		fmt.Println("done read", chName)
		return true
	}
}

func sendChannel(v string, ch chan string) {
	ch <- v
}

func readSomeChannels(v string, ch chan string, v2 string, ch2 chan string) {
	select {
	case <-ch:
		fmt.Println("done read", v)
	case <-ch2:
		fmt.Println("done read", v2)
	default:
	}
}

func readSomeChannelsLoop(v string, ch chan string, v2 string, ch2 chan string) {
	for {
		readSomeChannels(v, ch, v2, ch2)
	}
}
