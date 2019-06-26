package haction

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestActionChannel(t *testing.T) {
	count := int32(0)
	handlerAddCount := func(ctx *Context) { atomic.AddInt32(&count, 1) }

	root := NewChannel()
	root.AddBeforeMiddleWare(handlerAddCount)
	root.AddAfterMiddleWare(handlerAddCount)
	root.Register(1, handlerAddCount)

	subGroup := root.Group()
	{
		subGroup.AddBeforeMiddleWare(handlerAddCount)
		subGroup.Register(2, handlerAddCount)

		subGroup2 := subGroup.Group()
		{
			subGroup2.AddBeforeMiddleWare(handlerAddCount)
			subGroup2.Register(3, handlerAddCount)
		}
	}
	root.Start()

	root.Input(root.GetContext(1))
	time.Sleep(100 * time.Millisecond)
	if atomic.LoadInt32(&count) != 3 {
		t.Errorf("action group do fail")
	}

	count = 0
	root.Input(root.GetContext(2))
	time.Sleep(100 * time.Millisecond)
	if atomic.LoadInt32(&count) != 4 {
		t.Errorf("action group do fail")
	}

	count = 0
	root.Input(root.GetContext(3))
	time.Sleep(100 * time.Millisecond)
	if atomic.LoadInt32(&count) != 5 {
		t.Errorf("action group do fail")
	}
}
