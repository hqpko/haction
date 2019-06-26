package haction

import (
	"testing"
)

func TestGroup(t *testing.T) {
	count := 0
	handlerAddCount := func(ctx *Context) { count++ }

	root := NewGroup()
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

	root.Do(NewContext(1))
	if count != 3 {
		t.Errorf("action group do fail")
	}

	count = 0
	root.Do(NewContext(2))
	if count != 4 {
		t.Errorf("action group do fail")
	}

	count = 0
	root.Do(NewContext(3))
	if count != 5 {
		t.Errorf("action group do fail")
	}
}