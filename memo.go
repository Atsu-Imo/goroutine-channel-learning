package main

import (
	"fmt"
	"time"
)

// 一つのselectの中で複数のchannelの値を扱う
// 概ねOKに見えるが、defer実行時に大量にcase <- ch1を実行している、なぜかわからない
func case3() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	defer func() {
		close(ch1)
		fmt.Println("close1")
		close(ch2)
		fmt.Println("close2")
	}()

	go readSomeChannelsLoop("val1", ch1, "va2", ch2)
	ch1 <- "val1"
	ch2 <- "val2"
	time.Sleep(1 * time.Second)
}

// ch1, ch2をロックした状態で値が来るのをまつ
func case2() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	defer func() {
		close(ch1)
		fmt.Println("close1")
		close(ch2)
		fmt.Println("close2")
	}()

	go func() {
		readChannelWithoutDef("ch1", ch1)
		readChannelWithoutDef("ch2", ch2)
	}()
	sendChannel("ch1", ch1)
	sendChannel("ch2", ch2)
	time.Sleep(1 * time.Second)
}

// channelに対する送信が先に来てもいい
// goroutineを使っていないのであんまり意味がない
func case1() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	defer func() {
		close(ch1)
		fmt.Println("close1")
		close(ch2)
		fmt.Println("close2")
	}()

	sendChannel("ch1", ch1)
	sendChannel("ch2", ch2)

	readChannelWithoutDef("ch1", ch1)
	readChannelWithoutDef("ch2", ch2)
}

func huniki() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	defer func() {
		close(ch1)
		fmt.Println("close1")
		close(ch2)
		fmt.Println("close2")
	}()

	go func() {
		var ch1Done, ch2Done bool
		for {
			ch1Done = readChannel("ch1", ch1)
			ch2Done = readChannel("ch2", ch2)
			if ch1Done && ch2Done {
				break
			}
		}
	}()
	time.Sleep(1 * time.Second)
	sendChannel("ch1", ch1)
	sendChannel("ch2", ch2)
	time.Sleep(1 * time.Second)
}

// defaultが実行される
func doDefault() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	defer func() {
		close(ch1)
		fmt.Println("close1")
		close(ch2)
		fmt.Println("close2")
	}()

	// selectで
	select {
	case v := <-ch1:
		fmt.Println(v)
	case v := <-ch2:
		fmt.Println(v)
	default:
		fmt.Println("from default")
	}
	go func() {
		ch1 <- "ch1"
		ch2 <- "ch2"
	}()
}
