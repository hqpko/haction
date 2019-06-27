package haction

import (
	"sync"
)

type HandleAction func(*Context)

type actionHandler struct {
	id     int32
	handle HandleAction
	owner  *Group
}

func newAction(id int32, handle HandleAction, owner *Group) *actionHandler {
	return &actionHandler{id: id, handle: handle, owner: owner}
}

func (a *actionHandler) do(ctx *Context) {
	if a.owner != nil {
		a.owner.doBefore(ctx)
		if !ctx.isAbort() {
			a.handle(ctx)
			a.owner.doAfter(ctx)
		}
	}
}

type Group struct {
	sync.RWMutex
	parent *Group

	beforeMiddleWare []HandleAction
	actionHandlers   map[int32]*actionHandler
	afterMiddleWare  []HandleAction
}

func NewGroup() *Group {
	return &Group{actionHandlers: map[int32]*actionHandler{}}
}

func newGroup() *Group {
	return &Group{}
}

func (g *Group) Register(id int32, handler func(ctx *Context)) {
	g.register(newAction(id, handler, g))
}

func (g *Group) AddBeforeMiddleWare(handlers ...HandleAction) {
	g.Lock()
	g.Unlock()
	if g.beforeMiddleWare == nil {
		g.beforeMiddleWare = make([]HandleAction, 0)
	}
	g.beforeMiddleWare = append(g.beforeMiddleWare, handlers...)
}

func (g *Group) AddAfterMiddleWare(handlers ...HandleAction) {
	g.Lock()
	g.Unlock()
	if g.afterMiddleWare == nil {
		g.afterMiddleWare = make([]HandleAction, 0)
	}
	g.afterMiddleWare = append(g.afterMiddleWare, handlers...)
}

func (g *Group) register(action *actionHandler) {
	if g.parent != nil {
		g.parent.register(action)
	} else {
		g.Lock()
		g.Unlock()
		g.actionHandlers[action.id] = action
	}
}

func (g *Group) Group() *Group {
	return newGroup().setParent(g)
}

func (g *Group) Do(ctx *Context) {
	g.RLock()
	g.RUnlock()
	if action := g.actionHandlers[ctx.id]; action != nil {
		action.do(ctx)
	}
}

func (g *Group) doBefore(ctx *Context) {
	if g.parent != nil {
		g.parent.doBefore(ctx)
	}
	g.doHandlers(ctx, g.beforeMiddleWare)
}

func (g *Group) doAfter(ctx *Context) {
	g.doHandlers(ctx, g.afterMiddleWare)
	if g.parent != nil {
		g.parent.doAfter(ctx)
	}
}

func (g *Group) doHandlers(ctx *Context, handlers []HandleAction) {
	if ctx.isAbort() {
		return
	}
	for _, handler := range handlers {
		if ctx.isAbort() {
			return
		} else {
			handler(ctx)
		}
	}
}

func (g *Group) setParent(parent *Group) *Group {
	g.parent = parent
	return g
}
