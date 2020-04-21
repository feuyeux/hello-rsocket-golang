package requester

import (
	"context"
	"crypto/tls"
	"github.com/rsocket/rsocket-go"
)

func BuildTlsClient() (rsocket.Client, error) {
	cli, err := rsocket.Connect().
		DataMimeType("application/json").
		MetadataMimeType("message/x.rsocket.composite-metadata.v0").
		Fragment(1024).
		Transport("tcp://127.0.0.1:7878").
		StartTLS(context.Background(), &tls.Config{
			InsecureSkipVerify: true,
		})
	if err != nil {
		panic(err)
	}
	return cli, err
}
