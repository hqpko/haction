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
	return newGroup(g.root, append(g.middlewares, middlewares...)...)
}

func (g *group) Register(id int32, handler HandleAction) IGroup {
	g.root.register(id, append(g.middlewares, handler))
	return g
}

func (g *group) Use(handlers ...HandleAction) IGroup {
	g.middlewares = append(g.middlewares, handlers...)
	return g
}
