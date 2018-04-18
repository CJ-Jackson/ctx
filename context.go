package ctx

import (
	"context"
	"net/http"
)

type Context interface {
	Title() string
	SetTitle(title string)
	Data(name string) interface{}
	SetData(name string, value interface{})
	Dep(name string) interface{}
	SetDep(name string, dep interface{})
	Request() *http.Request
	Response() http.ResponseWriter
}

type contextHolder struct {
	title string
	data  map[string]interface{}
	dep   map[string]interface{}
	req   *http.Request
	res   http.ResponseWriter
}

func NewContext(req *http.Request, res http.ResponseWriter) (*http.Request, Context) {
	ctx := &contextHolder{
		title: "Untitled",
		data:  map[string]interface{}{},
		dep:   map[string]interface{}{},
		req:   req,
		res:   res,
	}

	return req.WithContext(context.WithValue(req.Context(), "contextHolder", ctx)), ctx
}

func (c *contextHolder) Title() string {
	return c.title
}

func (c *contextHolder) SetTitle(title string) {
	c.title = title
}

func (c *contextHolder) Data(name string) interface{} {
	return c.data[name]
}

func (c *contextHolder) SetData(name string, value interface{}) {
	c.data[name] = value
}

func (c *contextHolder) Dep(name string) interface{} {
	return c.dep[name]
}

func (c *contextHolder) SetDep(name string, dep interface{}) {
	c.dep[name] = dep
}

func (c *contextHolder) Request() *http.Request {
	return c.req
}

func (c *contextHolder) Response() http.ResponseWriter {
	return c.res
}
