package responder

import (
	"context"
	"crypto/tls"
	"errors"
	"log"
	"strings"

	rsocket "github.com/rsocket/rsocket-go"
	"github.com/rsocket/rsocket-go/payload"
)

func StartTls() {
	certFile, keyFile := "/Users/han/shop/cert.pem", "/Users/han/shop/key.pem"
	cert, err0 := tls.LoadX509KeyPair(certFile, keyFile)
	tc := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
		Certificates: []tls.Certificate{cert},
	}

	err0 = rsocket.Receive().
		Fragment(1024).
		Acceptor(
			func(setup payload.SetupPayload, sendingSocket rsocket.CloseableRSocket) (rsocket.RSocket, error) {
				sendingSocket.OnClose(func(err error) {
					log.Println("***** socket disconnected *****")
				})
				// For SETUP_REJECT testing.
				if strings.EqualFold(setup.DataUTF8(), "REJECT_ME") {
					return nil, errors.New("bye bye bye")
				}
				return HelloRSocket(), nil
			}).
		Transport("tcp://127.0.0.1:7878").
		//Serve(context.Background())
		ServeTLS(context.Background(), tc)
	panic(err0)
}
