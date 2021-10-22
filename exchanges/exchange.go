package exchanges

import (
	"fmt"
	. "github.com/516310460/trade-api"
	"github.com/516310460/trade-api/exchanges/binancefutures"
	"github.com/516310460/trade-api/exchanges/bitmex"
	"github.com/516310460/trade-api/exchanges/bybit"
	"github.com/516310460/trade-api/exchanges/deribit"
	"github.com/516310460/trade-api/exchanges/hbdm"
	"github.com/516310460/trade-api/exchanges/hbdmswap"
	"github.com/516310460/trade-api/exchanges/okexfutures"
	"github.com/516310460/trade-api/exchanges/okexswap"
)

func NewExchange(name string, opts ...ApiOption) Exchange {
	params := &Parameters{}

	for _, opt := range opts {
		opt(params)
	}

	return NewExchangeFromParameters(name, params)
}

func NewExchangeFromParameters(name string, params *Parameters) Exchange {
	switch name {
	case BinanceFutures:
		return binancefutures.NewBinanceFutures(params)
	case BitMEX:
		return bitmex.NewBitMEX(params)
	case Deribit:
		return deribit.NewDeribit(params)
	case Bybit:
		return bybit.NewBybit(params)
	case Hbdm:
		return hbdm.NewHbdm(params)
	case HbdmSwap:
		return hbdmswap.NewHbdmSwap(params)
	case OkexFutures:
		return okexfutures.NewOkexFutures(params)
	case OkexSwap:
		return okexswap.NewOkexSwap(params)
	default:
		panic(fmt.Sprintf("new exchange error [%v]", name))
	}
}
