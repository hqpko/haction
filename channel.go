package haction

import (
	"github.com/hqpko/hchannel"
)

type Channel struct {
	engine      *Engine
	mainChannel *hchannel.Channel
}

func NewChannel() *Channel {
	return NewChannelWithOption(1<<10, 1)
}

func NewChannelWithOption(channelSize, goroutineCount int) *Channel {
	c := &Channel{engine: NewEngine()}
	c.mainChannel = hchannel.NewChannelMulti(channelSize, goroutineCount, c.doAction)
	return c
}

func (c *Channel) Start() *Channel {
	c.mainChannel.Run()
	return c
}

func (c *Channel) Root() IGroup {
	return c.engine
}

func (c *Channel) Input(pid int32, values Values) bool {
	return c.mainChannel.Input(c.engine.newContext(pid, values))
}

func (c *Channel) doAction(i interface{}) {
	if ctx, ok := i.(*Context); ok {
		ctx.do()
	}
}

func (c *Channel) Stop() {
	c.mainChannel.Close()
}
