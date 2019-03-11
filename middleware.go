package middleware

import (
	"errors"
	"sync"
)

type Middleware struct {
	Handlers HandlersChain
	pool     sync.Pool
}

func New() *Middleware {
	m := &Middleware{}
	m.pool.New = func() interface{} {
		return m.allocateContext()
	}
	return m
}

func (m *Middleware) Use(middlewares ...HandlerFunc) {
	m.Handlers = append(m.Handlers, middlewares...)
}

// init logic in middleware.
func (m *Middleware) Run() (err error) {
	if m.Handlers == nil {
		return errors.New("Handlers is nil")
	}
	c := m.pool.Get().(*Context)

	// init
	c.reset()
	c.handlers = m.Handlers

	c.Next()

	m.pool.Put(c)
	return
}

// c.Set("params", params)
func (m *Middleware) RunWithParams(params ...interface{}) (err error) {
	if m.Handlers == nil {
		return errors.New("Handlers is nil")
	}
	c := m.pool.Get().(*Context)

	// init
	c.reset()
	c.handlers = m.Handlers
	c.Set("params", params)

	c.Next()

	m.pool.Put(c)
	return
}

func (m *Middleware) allocateContext() *Context {
	return &Context{mid: m}
}