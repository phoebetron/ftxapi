package public

import (
	"github.com/phoebetron/ftxapi/client/public/trade"
)

type Public struct {
	Tra *trade.T
}

func New(con Config) *Public {
	{
		con.Verify()
	}

	var tra *trade.T
	{
		tra = trade.New(trade.Config{
			Req: con.Req,
		})
	}

	return &Public{
		Tra: tra,
	}
}
