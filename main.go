package main

import (
	"log"
	"time"

	"github.com/feuyeux/hello-rsocket/requester"
	"github.com/feuyeux/hello-rsocket/responder"
)

func main() {
	//log.SetFlags(log.Ldate)
	log.SetFlags(log.Lmicroseconds)
	//startServer()
	runAll()
}

func runAll() {
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

func startServer() {
	go responder.Start()
	time.Sleep(100 * time.Millisecond)
}
