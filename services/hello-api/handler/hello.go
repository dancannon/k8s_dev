package handler

import (
	"github.com/dancannon/go-micro/client"
	"github.com/gin-gonic/gin"

	helloHelloPB "github.com/dancannon/k8s_dev/services/hello-service/proto/hello"
)

func Hello(c *gin.Context) {
	req := client.NewRequest("hello-service", "Hello.Call", &helloHelloPB.Request{
		Name: c.Params.ByName("name"),
	})

	// Set arbitrary headers
	req.Headers().Set("X-From-Id", "hello-api")

	rsp := &helloHelloPB.Response{}

	// Call service
	if err := client.Call(req, rsp); err != nil {
		c.Fail(500, err)
		return
	}

	c.JSON(200, map[string]interface{}{
		"message": rsp.GetMessage(),
	})
}
