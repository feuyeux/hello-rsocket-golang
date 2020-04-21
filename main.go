package main

import (
	"github.com/feuyeux/hello-rsocket/src/responder"
	"log"
	"time"

	"github.com/feuyeux/hello-rsocket/src/requester"
)

func main() {
	TLS := false
	isRunServer := true
	isRunClient := true

	//log.SetFlags(log.Ldate)
	log.SetFlags(log.Lmicroseconds)

	if isRunServer {
		if isRunClient {
			runResponse(false, TLS)
		} else {
			runResponse(true, TLS)
		}
	}

	if isRunClient {
		runRequest(TLS)
	}
}

func runResponse(block bool, TLS bool) {
	if block {
		startServer(TLS)
	} else {
		go startServer(TLS)
		time.Sleep(100 * time.Millisecond)
	}
}
func startServer(TLS bool) {
	if TLS {
		responder.StartTls()
	} else {
		responder.Start()
	}
}
func runRequest(TLS bool) {
	requester.TLS = TLS
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
