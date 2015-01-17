package handler

import (
	"fmt"

	"code.google.com/p/go.net/context"
	"github.com/Sirupsen/logrus"
	"github.com/asim/go-micro/server"

	helloHelloPB "github.com/dancannon/k8s_dev/services/hello-service/proto/hello"
)

type Hello struct{}

func (e *Hello) Call(ctx context.Context, req *helloHelloPB.Request, rsp *helloHelloPB.Response) error {
	logrus.Debug("Received Hello.Call request")

	rsp.Message = fmt.Sprintf("%s: Hello %s", server.Name, req.GetName())

	return nil
}
