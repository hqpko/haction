package haction

type HandleAction func(*Context)

type IGroup interface {
	Group(middlewares ...HandleAction) IGroup
	Use(handlers ...HandleAction) IGroup
	Register(id int32, handler HandleAction) IGroup
}

type group struct {
	root        *Engine
	middlewares []HandleAction
}

func newGroup(root *Engine, middlewares ...HandleAction) *group {
	return &group{root: root, middlewares: middlewares}
}

func (g *group) Group(middlewares ...HandleAction) IGroup {
	handlers := make([]HandleAction, 0, len(g.middlewares)+len(middlewares))
	handlers = append(handlers, g.middlewares...)
	handlers = append(handlers, middlewares...)
	return newGroup(g.root, handlers...)
}

func (g *group) Register(id int32, handler HandleAction) IGroup {
	handlers := make([]HandleAction, 0, len(g.middlewares)+1)
	handlers = append(handlers, g.middlewares...)
	handlers = append(handlers, handler)
	g.root.register(id, handlers)
	return g
}

func (g *group) Use(handlers ...HandleAction) IGroup {
	g.middlewares = append(g.middlewares, handlers...)
	return g
}
