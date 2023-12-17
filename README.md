# Deferpool

`Deferpool` is a utility designed to manage functions that need to be executed at the end of each iteration in a loop. In certain scenarios, using `defer` statements may execute functions only upon function return rather than at the end of each loop iteration. To address this, `Deferpool` offers a concise way to collect and execute functions that need to run at the end of each loop iteration. With the `Add` or `Defer` methods, you can easily append functions to be executed at the end of each loop iteration. By using the `Run` method, `Deferpool` executes all collected functions in reverse order, ensuring they are executed at the end of each loop iteration. Additionally, the `Remove` method allows dynamic removal of a specific number of functions from the collection, providing flexibility and customization.

`Deferpool` 是一个用于在循环中管理在每次循环结束时执行的函数的实用工具。在某些情况下，使用 `defer` 语句可能会在函数返回时执行，而不是在每次循环迭代结束时执行。为了解决这个问题，`Deferpool` 提供了一种简洁的方式来收集和执行在每次循环迭代结束时需要执行的函数。通过 `Add` 或 `Defer` 方法，您可以轻松地添加需要在循环结束时执行的函数。使用 `Run` 方法，`Deferpool` 将按逆序执行所有收集的函数，确保它们在每次循环迭代结束时得到执行。此外，`Remove` 方法允许您在需要的情况下动态地从集合中移除特定数量的函数，以满足灵活性和定制化的需求。

# Example

Require go v1.21.

```shell
go get github.com/xchacha20-poly1305/deferpool
```

```go
package main

import (
	"log"
	"net"

	"github.com/xchacha20-poly1305/deferpool"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:80")
	if err != nil {
		log.Fatal(err)
	}

	for {
		dp := deferpool.New()
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
		}
		dp.Defer(func() { conn.Close() })

		handle(conn)

		dp.Run()
	}
}

func handle(src net.Conn) {
	_, err := src.Write([]byte{0, 0, 0, 0, 0, 0, 0})
	if err != nil {
		log.Println(err)
	}
}
```