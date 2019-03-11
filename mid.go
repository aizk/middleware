package mid

type Mid struct {
	Handlers HandlersChain
}

func (m *Mid) Use(middlewares ...HandlerFunc) {
	m.Handlers = append(m.Handlers, middlewares...)
}

