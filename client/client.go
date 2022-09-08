package client

import (
	"net/http"

	"github.com/phoebetron/ftxapi/client/public"
	"github.com/phoebetron/ftxapi/client/request"
)

type Client struct {
	Pub *public.Public
}

func New(con Config) *Client {
	{
		con.Verify()
	}

	var pub *public.Public
	{
		pub = public.New(public.Config{
			Req: request.New(request.Config{
				Cli: &http.Client{},
			}),
		})
	}

	return &Client{
		Pub: pub,
	}
}
