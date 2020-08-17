package haction

type Engine struct {
	*group
	handlers map[int32][]HandleAction
}

func NewEngine() *Engine {
	e := &Engine{handlers: map[int32][]HandleAction{}}
	e.group = newGroup(e)
	return e
}

func (e *Engine) register(pid int32, handlers []HandleAction) {
	e.handlers[pid] = handlers
}

func (e *Engine) Do(id int32, values Values) {
	newContext(id, values, e.handlers[id]).do()
}

func (e *Engine) newContext(id int32, values Values) *Context {
	return newContext(id, values, e.handlers[id])
}
