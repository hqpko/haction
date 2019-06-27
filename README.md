# haction

### install

```bash
go get -v -u github.com/hqpko/haction
```

### example
```go
package main

import (
	"github.com/hqpko/haction"
)

func main() {
	root := haction.NewChannelWithOption(1024, 1)
	root.AddBeforeMiddleWare(rootBeforeMiddleWare01, rootBeforeMiddleWare02)
	root.Register(1, handler01)

	// sub group
	subGroup := root.Group()
	{
		subGroup.AddBeforeMiddleWare(subBeforeMiddleWare01, subBeforeMiddleWare02)
		subGroup.Register(2, handler02)

		// sub group
		subGroup2 := subGroup.Group()
		{
			subGroup2.Register(3, handler03)
		}

		subGroup.AddAfterMiddleWare(subAfterMiddleWare01, subAfterMiddleWare02)
	}

	root.AddAfterMiddleWare(rootAfterMiddleWare01, rootAfterMiddleWare02)

	// 执行顺序
	// before	: rootBeforeMiddleWare01 -> rootBeforeMiddleWare02 -> subBeforeMiddleWare01 -> subBeforeMiddleWare02 ->
	// register	: handler03 ->
	// after	: subAfterMiddleWare01 -> subAfterMiddleWare02 -> rootAfterMiddleWare01 -> rootAfterMiddleWare02
	ctx := root.GetContext(3).Set("value", "123")
	root.Input(ctx)

	// Group 类似于 Channel，只是入口是 Group.Do(ctx)，同步执行，而不是 Channel.Input(ctx)，异步执行
}

```