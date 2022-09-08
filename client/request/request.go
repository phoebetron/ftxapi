package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/phoebetron/ftxapi/errors"
	"github.com/xh3b4sd/tracer"
)

type Request struct {
	cli *http.Client
}

func New(con Config) *Request {
	{
		con = con.Ensure()
	}

	{
		con.Verify()
	}

	return &Request{
		cli: con.Cli,
	}
}

func (r *Request) endp(pat string) string {
	var end string
	{
		end = "https://ftx.com/api"
	}

	return fmt.Sprintf("%s/%s", end, pat)
}

func (r *Request) request(met string, pat string, dat interface{}) ([]byte, error) {
	var err error

	var bod []byte
	if dat != nil {
		bod, err = json.Marshal(dat)
		if err != nil {
			panic(err)
		}
	}

	var end string
	{
		end = r.endp(pat)
	}

	var req *http.Request
	{
		req, err = http.NewRequest(met, end, bytes.NewBuffer(bod))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if dat != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	var res *http.Response
	{
		res, err = r.cli.Do(req)
		if err != nil {
			return nil, tracer.Mask(err)
		}
		defer res.Body.Close()
	}

	var byt []byte
	{
		byt, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var lis errors.Errors
	{
		err = json.Unmarshal(byt, &lis)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if lis.Error != "" {
		return nil, tracer.Mask(fmt.Errorf(lis.Error))
	}

	return byt, nil
}
