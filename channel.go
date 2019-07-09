package haction

import (
	"github.com/hqpko/hconcurrent"
)

type Channel struct {
	group       *Group
	mainChannel *hconcurrent.Concurrent
}

func NewChannel() *Channel {
	return NewChannelWithOption(1<<10, 1)
}

func NewChannelWithOption(channelSize, goroutineCount int) *Channel {
	c := &Channel{group: NewGroup().SetContextPool(NewContextPool())}
	c.mainChannel = hconcurrent.NewConcurrent(channelSize, goroutineCount, c.doAction)
	return c
}

func (c *Channel) SetContextPool(pool *ContextPool) *Channel {
	c.group.SetContextPool(pool)
	return c
}

func (c *Channel) Start() *Channel {
	c.mainChannel.Start()
	return c
}

func (c *Channel) Register(id int32, handler func(ctx *Context)) IGroup {
	return c.group.Register(id, handler)
}

func (c *Channel) AddBeforeMiddleWare(handler func(ctx *Context)) IGroup {
	return c.group.AddBeforeMiddleWare(handler)
}

func (c *Channel) AddAfterMiddleWare(handler func(ctx *Context)) IGroup {
	return c.group.AddAfterMiddleWare(handler)
}

func (c *Channel) Group() IGroup {
	return c.group.Group()
}

func (c *Channel) Input(pid int32, values Values) bool {
	return c.mainChannel.Input(c.getContext(pid).SetValues(values))
}

func (c *Channel) MustInput(pid int32, values Values) {
	c.mainChannel.MustInput(c.getContext(pid).SetValues(values))
}

func (c *Channel) doAction(i interface{}) interface{} {
	if ctx, ok := i.(*Context); ok {
		c.group.do(ctx)
		c.group.pool.Put(ctx)
	}
	return nil
}

func (c *Channel) getContext(id int32) *Context {
	return c.group.pool.Get(id)
}

func (c *Channel) Stop() {
	c.mainChannel.Stop()
}
