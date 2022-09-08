# ftxapi

```go
var err error

var cli *client.Client
{
        cli = client.New(client.Config{})
}

var req trade.ListRequest
{
        req = trade.ListRequest{
                ProductCode: "ETH-PERP",
                Limit:       1,
        }
}

var res trade.ListResponse
{
        res, err = s.cli.Pub.Tra.List(req)
        if err != nil {
                panic(err)
        }
}

// print JSON res.Result[0]
```

```
{
	"id":4903776997,
	"price":1633.4,
	"size":0.02,
	"side":"buy",
	"liquidation":false,
	"time":"2022-09-08T13:50:35.045006+00:00"
}
```
