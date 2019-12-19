## hello-rsocket-golang

![](https://github.com/feuyeux/hello-rsocket/blob/master/doc/hello-rsocket.png)

```bash
▶ go run main.go
2019/12/19 ====ExecMetaPush====
2019/12/19 >> [MetadataPush]:  GOLANG

2019/12/19 ====ExecFireAndForget====
2019/12/19 >> [FireAndForget] FNF: 1

2019/12/19 ====ExecRequestResponse====
2019/12/19 >> [Request-Response] data: {1} , metadata: 2019-12-19 16:30:59,059
2019/12/19 << [Request-Response] response id: 1 , value: Bonjour

2019/12/19 ====ExecRequestStream====
2019/12/19 >> [Request-Stream] data: {[1 2 2 4 1]}
2019/12/19 << [Request-Stream] response id: 1 , value: Bonjour
2019/12/19 << [Request-Stream] response id: 2 , value: Hola
2019/12/19 << [Request-Stream] response id: 2 , value: Hola
2019/12/19 << [Request-Stream] response id: 4 , value: Ciao
2019/12/19 << [Request-Stream] response id: 1 , value: Bonjour

2019/12/19 ====ExecRequestChannel====
2019/12/19 >> [Request-Channel] data: {[0 1 2]}
2019/12/19 << [Request-Channel] response id: 0 , value: Hello
2019/12/19 << [Request-Channel] response id: 1 , value: Bonjour
2019/12/19 << [Request-Channel] response id: 2 , value: Hola
2019/12/19 >> [Request-Channel] data: {[1 0 2]}
2019/12/19 << [Request-Channel] response id: 1 , value: Bonjour
2019/12/19 << [Request-Channel] response id: 0 , value: Hello
2019/12/19 << [Request-Channel] response id: 2 , value: Hola
2019/12/19 >> [Request-Channel] data: {[1 0 2]}
2019/12/19 << [Request-Channel] response id: 1 , value: Bonjour
2019/12/19 << [Request-Channel] response id: 0 , value: Hello
2019/12/19 << [Request-Channel] response id: 2 , value: Hola

```