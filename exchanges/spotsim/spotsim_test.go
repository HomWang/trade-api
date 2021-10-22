package spotsim

import (
	"github.com/516310460/trade-api"
	"testing"
)

var ss = New(
	"huobi",
	nil,
	crex.SpotBalance{
		Base:  crex.SpotAsset{Name: "BTC", Available: 1, Frozen: 0},
		Quote: crex.SpotAsset{Name: "USDT", Available: 10000, Frozen: 0},
	},
	0.0001,
	0.0003,
)

func TestSpotSim_GetName(t *testing.T) {
	t.Log(ss.GetName())
}
