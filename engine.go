package haction

import "sync"

type Engine struct {
	lock sync.RWMutex
	*group
	handlers map[int32][]HandleAction
}

func NewEngine() *Engine {
	e := &Engine{handlers: map[int32][]HandleAction{}}
	e.group = newGroup(e)
	return e
}

func (e *Engine) register(pid int32, handlers []HandleAction) {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.handlers[pid] = handlers
}

func (e *Engine) Do(id int32, values Values) {
	newContext(id, values, e.getHandler(id)).do()
}

func (e *Engine) newContext(id int32, values Values) *Context {
	return newContext(id, values, e.getHandler(id))
}

func (e *Engine) getHandler(id int32) []HandleAction {
	e.lock.RLock()
	defer e.lock.RUnlock()
	return e.handlers[id]
}
