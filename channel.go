package haction

import (
	"sync"

	"github.com/hqpko/hconcurrent"
)

type Channel struct {
	pool        sync.Pool
	group       *Group
	mainChannel *hconcurrent.Concurrent
}

func NewChannel() *Channel {
	return NewChannelWithOption(1<<10, 1)
}

func NewChannelWithOption(channelSize, goroutineCount int) *Channel {
	c := &Channel{group: NewGroup(), pool: sync.Pool{New: func() interface{} {
		return NewContext(0)
	}}}
	c.mainChannel = hconcurrent.NewConcurrent(channelSize, goroutineCount, c.doAction)
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

func (c *Channel) Input(ctx *Context) bool {
	return c.mainChannel.Input(ctx)
}

func (c *Channel) MustInput(ctx *Context) {
	c.mainChannel.MustInput(ctx)
}

func (c *Channel) doAction(i interface{}) interface{} {
	c.group.Do(i.(*Context))
	return nil
}

func (c *Channel) GetContext(id int32) *Context {
	ctx := c.pool.Get().(*Context)
	ctx.SetID(id)
	return ctx
}

func (c *Channel) putContext(ctx *Context) {
	ctx.Reset()
	c.pool.Put(ctx)
}

func (c *Channel) Stop() {
	c.mainChannel.Stop()
}
