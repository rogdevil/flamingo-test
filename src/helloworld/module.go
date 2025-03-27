package helloworld

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
	"github.com/rogdevil/flamingo-test/src/helloworld/interfaces"
)

type Module struct{}

func (m *Module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(routes))
}

type routes struct {
	helloController *interfaces.HelloController
}

func (r *routes) Inject(helloController *interfaces.HelloController) *routes {
	r.helloController = helloController
	return r
}

func (r *routes) Routes(registry *web.RouterRegistry) {
	// Bind the path /hello to a handle with the name "hello"
	registry.MustRoute("/hello", "hello")

	// Bind the controller.Action to the handle "hello":
	registry.HandleGet("hello", r.helloController.Hello)
	registry.MustRoute("/greet/:nickname", "/hello/{nickname}")
	registry.HandleGet("/hello/{nickname}", r.helloController.Greet)
	registry.HandlePost("/hello/{nickname}", r.helloController.GreetPost)
}
