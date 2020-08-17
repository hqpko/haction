# haction

### install

```bash
go get -v -u github.com/hqpko/haction
```

### example
`middleware` 接口改为类似 `gin` 的方式
```go
package main

import (
	"github.com/hqpko/haction"
)

func main() {
	channel := haction.NewChannelWithOption(1024, 1)
	root := channel.Root()
	root.Use(rootBeforeMiddleWare01, rootBeforeMiddleWare02)
	root.Register(1, handler01)

	// sub group
	subGroup := root.Group(subBeforeMiddleWare01, subBeforeMiddleWare02)
	{
		subGroup.Register(2, handler02)

		// sub group
		subGroup2 := subGroup.Group()
		{
			subGroup2.Register(3, handler03)
		}
	}

	channel.Input(3, haction.Values{
		"value": "123",
	})

	// Engine 类似于 Channel，只是入口是 Engine.Do(pid, values)，同步执行，而不是 Channel.Input(pid, values)，异步执行
}

```