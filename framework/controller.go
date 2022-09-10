package framework

type ControllerHandler func(c *Context) error

//func FooControllerHandler(ctx *Context) error {
//return ctx.Json(200, map[string]interface{}{
//	"code": 0,
//})
//}
