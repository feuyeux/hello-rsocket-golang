package test

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/jjeffcaii/reactor-go"
	"github.com/jjeffcaii/reactor-go/flux"
	"github.com/jjeffcaii/reactor-go/scheduler"
)

func TestExample(t *testing.T) {
	gen := func(ctx context.Context, sink flux.Sink) {
		for i := 0; i < 10; i++ {
			sink.Next(i)
		}
		sink.Complete()
	}
	done := make(chan struct{})

	var su rs.Subscription
	flux.Create(gen).
		Filter(func(i interface{}) bool {
			return i.(int)%2 == 0
		}).
		Map(func(i interface{}) interface{} {
			return fmt.Sprintf("#HELLO_%04d", i.(int))
		}).
		SubscribeOn(scheduler.Elastic()).
		Subscribe(context.Background(),
			rs.OnSubscribe(func(s rs.Subscription) {
				su = s
				s.Request(1)
			}),
			rs.OnNext(func(v interface{}) {
				fmt.Println("next:", v)
				su.Request(1)
			}),
			rs.OnComplete(func() {
				close(done)
			}),
		)
	<-done
}

func TestReactor(t *testing.T) {
	letters := strings.Split("The quick brown fox jumps over a lazy dog", "")
	gen := func(ctx context.Context, sink flux.Sink) {
		for _, letter := range letters {
			sink.Next(letter)
		}
		sink.Complete()
	}
	done := make(chan struct{})

	var su rs.Subscription

	//flux.Just(letters).
	flux.Create(gen).
		Filter(func(i interface{}) bool {
			return i != " "
		}).
		Map(func(i interface{}) interface{} {
			str := fmt.Sprintf("%v", i)
			return strings.ToLower(str)
		}).
		SubscribeOn(scheduler.Elastic()).
		Subscribe(context.Background(),
			rs.OnSubscribe(func(s rs.Subscription) {
				su = s
				s.Request(1)
			}),
			rs.OnNext(func(v interface{}) {
				log.Println("next:", v)
				su.Request(1)
			}),
			rs.OnComplete(func() {
				log.Println("Complete!")
				close(done)
			}),
		)
	<-done
}
