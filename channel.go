package haction

import (
	"github.com/hqpko/hchannel"
)

type Channel struct {
	group       *Group
	mainChannel *hchannel.Channel
}

func NewChannel() *Channel {
	return NewChannelWithOption(1<<10, 1)
}

func NewChannelWithOption(channelSize, goroutineCount int) *Channel {
	c := &Channel{group: NewGroup().SetContextPool(NewContextPool())}
	c.mainChannel = hchannel.NewChannelMulti(channelSize, goroutineCount, c.doAction)
	return c
}

func (c *Channel) SetContextPool(pool *ContextPool) *Channel {
	c.group.SetContextPool(pool)
	return c
}

func (c *Channel) Start() *Channel {
	c.mainChannel.Run()
	return c
}

func (c *Channel) Root() IGroup {
	return c.group.Root()
}

func (c *Channel) Input(pid int32, values Values) bool {
	return c.mainChannel.Input(c.getContext(pid).SetValues(values))
}

func (c *Channel) doAction(i interface{}) {
	if ctx, ok := i.(*Context); ok {
		c.group.do(ctx)
		c.group.pool.Put(ctx)
	}
}

func (c *Channel) getContext(id int32) *Context {
	return c.group.pool.Get(id)
}

func (c *Channel) Stop() {
	c.mainChannel.Close()
}
