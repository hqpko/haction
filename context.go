package haction

import (
	"time"
)

type Values map[string]interface{}

type Context struct {
	id       int32
	index    int
	values   Values
	handlers []HandleAction
}

func newContext(id int32, values Values, handlers []HandleAction) *Context {
	return &Context{id: id, values: values, handlers: handlers}
}

func (c *Context) Abort() {
	c.index = len(c.handlers)
}

func (c *Context) Next() {
	c.index++
	c.do()
}

func (c *Context) do() {
	for c.index < len(c.handlers) {
		c.handlers[c.index](c)
		c.index++
	}
}

func (c *Context) Set(key string, value interface{}) *Context {
	c.values[key] = value
	return c
}

func (c *Context) Get(key string) (value interface{}, exists bool) {
	value, exists = c.values[key]
	return
}

func (c *Context) MustGet(key string) interface{} {
	return c.values[key]
}

func (c *Context) GetString(key string) (s string) {
	if v, ok := c.Get(key); ok && v != nil {
		s, _ = v.(string)
	}
	return
}

func (c *Context) GetBool(key string) (b bool) {
	if v, ok := c.Get(key); ok && v != nil {
		b, _ = v.(bool)
	}
	return
}

func (c *Context) GetInt(key string) (i int) {
	if v, ok := c.Get(key); ok && v != nil {
		i, _ = v.(int)
	}
	return
}

func (c *Context) GetInt32(key string) (i32 int32) {
	if v, ok := c.Get(key); ok && v != nil {
		i32, _ = v.(int32)
	}
	return
}

func (c *Context) GetInt64(key string) (i64 int64) {
	if v, ok := c.Get(key); ok && v != nil {
		i64, _ = v.(int64)
	}
	return
}

func (c *Context) GetUint32(key string) (u32 uint32) {
	if v, ok := c.Get(key); ok && v != nil {
		u32, _ = v.(uint32)
	}
	return
}

func (c *Context) GetUint64(key string) (u64 uint64) {
	if v, ok := c.Get(key); ok && v != nil {
		u64, _ = v.(uint64)
	}
	return
}

func (c *Context) GetFloat32(key string) (f32 float32) {
	if v, ok := c.Get(key); ok && v != nil {
		f32, _ = v.(float32)
	}
	return
}

func (c *Context) GetFloat64(key string) (f64 float64) {
	if v, ok := c.Get(key); ok && v != nil {
		f64, _ = v.(float64)
	}
	return
}

func (c *Context) GetTime(key string) (t time.Time) {
	if v, ok := c.Get(key); ok && v != nil {
		t, _ = v.(time.Time)
	}
	return
}

func (c *Context) GetDuration(key string) (d time.Duration) {
	if v, ok := c.Get(key); ok && v != nil {
		d, _ = v.(time.Duration)
	}
	return
}

func (c *Context) GetStringSlice(key string) (ss []string) {
	if v, ok := c.Get(key); ok && v != nil {
		ss, _ = v.([]string)
	}
	return
}

func (c *Context) GetStringMap(key string) (sm map[string]interface{}) {
	if v, ok := c.Get(key); ok && v != nil {
		sm, _ = v.(map[string]interface{})
	}
	return
}

func (c *Context) GetStringMapString(key string) (sms map[string]string) {
	if v, ok := c.Get(key); ok && v != nil {
		sms, _ = v.(map[string]string)
	}
	return
}

func (c *Context) GetValues() Values {
	return c.values
}
