package responder

import (
	"context"
	"github.com/feuyeux/hello-rsocket/common"
	"github.com/jjeffcaii/reactor-go/scheduler"
	"log"
	"strconv"
	"time"

	rsocket "github.com/rsocket/rsocket-go"

	"github.com/rsocket/rsocket-go/payload"
	"github.com/rsocket/rsocket-go/rx"
	"github.com/rsocket/rsocket-go/rx/flux"
	"github.com/rsocket/rsocket-go/rx/mono"
)

func HelloRSocket() rsocket.RSocket {
	helloList := []string{"Hello", "Bonjour", "Hola", "こんにちは", "Ciao", "안녕하세요"}

	return rsocket.NewAbstractSocket(
		rsocket.MetadataPush(func(p payload.Payload) {
			meta, _ := p.MetadataUTF8()
			log.Println(">> [MetadataPush]:", meta)
		}),
		rsocket.FireAndForget(func(p payload.Payload) {
			data := p.Data()
			request := common.JsonToHelloRequest(data)
			log.Println(">> [FireAndForget] FNF:", request.Id)
		}),
		rsocket.RequestResponse(func(p payload.Payload) mono.Mono {
			data := p.Data()
			request := common.JsonToHelloRequest(data)
			metadata, _ := p.MetadataUTF8()
			log.Println(">> [Request-Response] data:", request, ", metadata:", metadata)
			id := request.Id
			index, _ := strconv.Atoi(id)
			response := common.HelloResponse{Id: id, Value: helloList[index]}
			json, _ := response.ToJson()
			meta, _ := p.Metadata()
			return mono.Just(payload.New(json, meta))
		}),
		rsocket.RequestStream(func(p payload.Payload) flux.Flux {
			data := p.Data()
			request := common.JsonToHelloRequests(data)
			log.Println(">> [Request-Stream] data:", request)

			return flux.Create(func(ctx context.Context, emitter flux.Sink) {
				for i := range request.Ids {
					// You can use context for graceful coroutine shutdown, stop produce.
					select {
					case <-ctx.Done():
						log.Println(">> [Request-Stream] ctx done:", ctx.Err())
						return
					default:
						id := request.Ids[i]
						index, _ := strconv.Atoi(id)
						response := common.HelloResponse{Id: id, Value: helloList[index]}
						json, _ := response.ToJson()
						meta, _ := p.Metadata()
						rp := payload.New(json, meta)
						emitter.Next(rp)
						time.Sleep(500 * time.Millisecond)
					}
				}
				emitter.Complete()
			})
		}),
		rsocket.RequestChannel(func(payloads rx.Publisher) flux.Flux {
			return flux.Create(func(i context.Context, sink flux.Sink) {
				payloads.(flux.Flux).
					SubscribeOn(scheduler.Elastic()).
					DoOnNext(func(p payload.Payload) {
						data := p.Data()
						//request := common.JsonToHelloRequest(data)
						request := common.JsonToHelloRequests(data)
						log.Println(">> [Request-Channel] data:", request)
						for _, id := range request.Ids {
							index, _ := strconv.Atoi(id)
							response := common.HelloResponse{Id: id, Value: helloList[index]}
							json, _ := response.ToJson()
							sink.Next(payload.New(json, nil))
						}
					}).
					Subscribe(context.Background())
				//sink.Complete()
			})
		}),
	)
}
