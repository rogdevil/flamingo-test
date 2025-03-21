package interfaces

import (
	"context"
	"fmt"

	"flamingo.me/flamingo/v3/framework/web"
)

type HelloController struct {
	responder *web.Responder
}

func (h *HelloController) Inject(responder *web.Responder) *HelloController {
	h.responder = responder
	return h
}

func (h *HelloController) Hello(ctx context.Context, req *web.Request) web.Result {
	return h.responder.Data(map[string]interface{}{
		"message": "Hello World",
	})
}

func (h *HelloController) Greet(ctx context.Context, req *web.Request) web.Result {
	nickname := req.Params["nickname"]
	return h.responder.Data(map[string]interface{}{
		"message": fmt.Sprintf("Hello %s", nickname),
	})
}

func (h *HelloController) GreetPost(ctx context.Context, req *web.Request) web.Result {
	nickname := req.Params["nickname"]
	return h.responder.Data(map[string]interface{}{
		"message": fmt.Sprintf("posting Hello %s", nickname),
	})
}
