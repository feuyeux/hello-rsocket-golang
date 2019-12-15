## hello-rsocket-golang

![](https://github.com/feuyeux/hello-rsocket/blob/master/doc/hello-rsocket.png)

```bash
▶ go run main.go

2019/12/15 22:17:39 ====ExecMetaPush====
2019/12/15 22:17:39 [Responder::MetadataPush] GOT METADATA_PUSH: GOLANG

2019/12/15 22:17:39 ====ExecFireAndForget====
2019/12/15 22:17:39 [Responder::FireAndForget] GOT FNF: 1

2019/12/15 22:17:39 ====ExecRequestResponse====
2019/12/15 22:17:39 [Responder::RequestResponse] data: {1} , metadata: 2019-12-15 22:17:39,039
2019/12/15 22:17:39 [Request-Response] response id: 1 ,value: Bonjour

2019/12/15 22:17:39 ====ExecRequestStream====
2019/12/15 22:17:39 [Responder::RequestStream] data: {[1 2 2 4 1]} , metadata: 2019-12-15 22:17:39,039
2019/12/15 22:17:39 [Request-Stream] response: {1 Bonjour}
2019/12/15 22:17:40 [Request-Stream] response: {2 Hola}
2019/12/15 22:17:40 [Request-Stream] response: {2 Hola}
2019/12/15 22:17:41 [Request-Stream] response: {4 Ciao}
2019/12/15 22:17:41 [Request-Stream] response: {1 Bonjour}

2019/12/15 22:17:42 ====ExecRequestChannel====
2019/12/15 22:17:42 [Responder::RequestChannel] data: {[0 1 2]} , metadata: 2019-12-15 22:17:42,042
2019/12/15 22:17:42 [Request-Channel] response: {0 Hello}
2019/12/15 22:17:42 [Request-Channel] response: {1 Bonjour}
2019/12/15 22:17:42 [Request-Channel] response: {2 Hola}
2019/12/15 22:17:42 [Responder::RequestChannel] data: {[1 0 2]} , metadata: 2019-12-15 22:17:42,042
2019/12/15 22:17:42 [Request-Channel] response: {1 Bonjour}
2019/12/15 22:17:42 [Request-Channel] response: {0 Hello}
2019/12/15 22:17:42 [Request-Channel] response: {2 Hola}
2019/12/15 22:17:42 [Responder::RequestChannel] data: {[1 0 2]} , metadata: 2019-12-15 22:17:42,042
2019/12/15 22:17:42 [Request-Channel] response: {1 Bonjour}
2019/12/15 22:17:42 [Request-Channel] response: {0 Hello}
2019/12/15 22:17:42 [Request-Channel] response: {2 Hola}

```