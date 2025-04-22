package haction

import (
	"log"

	"github.com/hqpko/hchannel"
)

var defUnknownProtocolHandler = func(pid int32) {
	log.Printf("unknown protocol id: %d", pid)
}

type Channel struct {
	engine                 *Engine
	mainChannel            *hchannel.Channel
	handlerUnknownProtocol func(pid int32)
}

func NewChannel() *Channel {
	return NewChannelWithOption(1<<10, 1)
}

func NewChannelWithOption(channelSize, goroutineCount int) *Channel {
	c := &Channel{engine: NewEngine(), handlerUnknownProtocol: defUnknownProtocolHandler}
	c.mainChannel = hchannel.NewChannelMulti(channelSize, goroutineCount, c.doAction)
	return c
}

func (c *Channel) SetUnknownProtocolHandler(handler func(pid int32)) *Channel {
	c.handlerUnknownProtocol = handler
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
		if ctx.handlers == nil {
			c.handlerUnknownProtocol(ctx.id)
		} else {
			ctx.do()
		}
	}
}

func (c *Channel) Stop() {
	c.mainChannel.Close()
	c.engine.clean()
	c.engine = nil
	c.mainChannel = nil
	c.handlerUnknownProtocol = nil
}
