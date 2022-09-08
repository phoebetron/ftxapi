package trade

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/phoebetron/ftxapi/client/request"
)

func Test_Client_Public_Trade_List_Request(t *testing.T) {
	req := ListRequest{
		ProductCode: "ETH-PERP",
		Limit:       1,
	}

	cur := request.Values(req).Encode()
	des := "limit=1"

	if cur != des {
		t.Fatalf("\n\n%s\n", cmp.Diff(des, cur))
	}
}
