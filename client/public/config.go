package public

import (
	"github.com/phoebetron/ftxapi/client/request"
)

type Config struct {
	Req *request.Request
	Tes bool
}

func (c Config) Verify() {
	if c.Req == nil {
		panic("Config.Req must not be empty")
	}
}
