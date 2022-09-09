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
		StartTime:   15,
		EndTime:     18,
	}

	cur := request.Values(req).Encode()
	des := "end_time=18&limit=1&start_time=15"

	if cur != des {
		t.Fatalf("\n\n%s\n", cmp.Diff(des, cur))
	}
}
