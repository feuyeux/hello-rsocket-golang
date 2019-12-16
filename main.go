package main

import (
	"time"

	"github.com/feuyeux/hello-rsocket/requester"
	"github.com/feuyeux/hello-rsocket/responder"
)

func main() {
	go responder.Start()
	time.Sleep(100 * time.Millisecond)

	requester.ExecMetaPush()
	time.Sleep(100 * time.Millisecond)

	requester.ExecFireAndForget()
	time.Sleep(100 * time.Millisecond)

	requester.ExecRequestResponse()
	time.Sleep(100 * time.Millisecond)

	requester.ExecRequestStream()
	time.Sleep(200 * time.Millisecond)

	go requester.ExecRequestChannel()
	time.Sleep(5 * time.Second)
}
