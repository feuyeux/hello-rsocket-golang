package main

import (
	"log"
	"time"

	"github.com/feuyeux/hello-rsocket/src/requester"
	"github.com/feuyeux/hello-rsocket/src/responder"
)

func main() {
	isRunServer := true
	isRunClient := true

	//log.SetFlags(log.Ldate)
	log.SetFlags(log.Lmicroseconds)

	if isRunServer {
		if isRunClient {
			runResponse(false)
		} else {
			runResponse(true)
		}
	}

	if isRunClient {
		runRequest()
	}
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
