# Deferstack

`Deferstack` is a lib to manager functions. You can use it like defer.

`Deferstack` 是一个管理函数的库。您可以像使用 `defer` 一样使用它。

# Example

```shell
go get github.com/xchacha20-poly1305/deferstack
```

```go
package main

import (
	"log"
	"net"

	"github.com/xchacha20-poly1305/deferstack"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:80")
	if err != nil {
		log.Fatal(err)
	}

	for {
		ds := deferstack.New()
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
		}
		ds.Add(func() { conn.Close() })

		handle(conn)

		ds.Run()
	}
}

func handle(src net.Conn) {
	_, err := src.Write([]byte{0, 0, 0, 0, 0, 0, 0})
	if err != nil {
		log.Println(err)
	}
}
```
