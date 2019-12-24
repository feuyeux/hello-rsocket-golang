package main

import (
	"log"
	"time"

	"github.com/feuyeux/hello-rsocket/requester"
	"github.com/feuyeux/hello-rsocket/responder"
)

func main() {
	isBlockServer := false

	//log.SetFlags(log.Ldate)
	log.SetFlags(log.Lmicroseconds)

	runResponse(isBlockServer)
	runRequest()
}

func runRequest() {
	metaPush()
	fnf()
	rr()
	rs()
	rc()
}

func rc() {
	go requester.ExecRequestChannel()
	time.Sleep(5 * time.Second)
}

func rs() {
	requester.ExecRequestStream()
	time.Sleep(200 * time.Millisecond)
}

func rr() {
	requester.ExecRequestResponse()
	time.Sleep(100 * time.Millisecond)
}

func fnf() {
	requester.ExecFireAndForget()
	time.Sleep(100 * time.Millisecond)
}

func metaPush() {
	requester.ExecMetaPush()
	time.Sleep(100 * time.Millisecond)
}

func runResponse(block bool) {
	if block {
		responder.Start()
	} else {
		go responder.Start()
		time.Sleep(100 * time.Millisecond)
	}
}
