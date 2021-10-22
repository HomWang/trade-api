package trade

import (
	"github.com/516310460/trade-api/exchanges"
	"log"
)

func getGetBalance(){
	exchange := exchanges.NewExchange(exchanges.HbdmSwap,
		//ApiProxyURLOption("socks5://127.0.0.1:7890"), // 使用代理
		ApiAccessKeyOption("fd736c5c-607ed469-b39e453d-xa2b53ggfc"),
		ApiSecretKeyOption("c1273411-9c8ceec7-9f96b2a5-4d18e"),
		//ApiPassPhraseOption("[passphrase]"),
		ApiTestnetOption(false))//火币使用

	currency := "DOGE-USD"
	balance, err := exchange.GetBalance(currency)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("balance: %#v", balance)

	symbol := "DOGE-USD"
	positions, err1 := exchange.GetPositions(symbol)
	if err1 != nil {
		log.Fatal(err1)
		return
	}
	log.Printf("%#v", positions)

}

func main() {
	//ws := exchanges.NewExchange(exchanges.HbdmSwap,
	//	//ApiProxyURLOption("socks5://127.0.0.1:7890"), // 使用代理
	//	//ApiWsURLOption("api.btcgateway.pro"),
	//	ApiAccessKeyOption("fd736c5c-607ed469-b39e453d-xa2b53ggfc"),
	//	ApiSecretKeyOption("c1273411-9c8ceec7-9f96b2a5-4d18e"),
	//	//ApiPassPhraseOption("[passphrase]"),
	//	ApiTestnetOption(false),//火币使用
	//	ApiWebSocketOption(true)) // 开启 WebSocket

	//binance API Key
	//apikey = "kFz6K9zBbwyjjFBCc8ROkqzZ4v2RcTfJtqE2ZAAcMuuQ3et2dE31VsRJxI79OI5B"
	//secretkey = "I61cwdnfWYylsL5Fe7OJjuJkxcDl8M2EU12zCNHNzwMqDkziwcxlh3teY3aIZT34"
	//IP = ""
	//备注名 = "量化交易"
	//权限 = "只读/合约交易/现货交易/杠杠交易"

	//火币 API Key
	//apikey = "fd736c5c-607ed469-b39e453d-xa2b53ggfc"
	//secretkey = "c1273411-9c8ceec7-9f96b2a5-4d18e"
	//IP = ""
	//备注名 = "量化交易"
	//权限 = "只读/交易"

	//okex API Key
	//apikey = "ef5003bf-b455-4d49-be63-4835ff0281bf"
	//secretkey = "8955BCCEA9661A0E62F230D6BA4B00CB"
	//IP = ""
	//备注名 = "量化交易"
	//权限 = "只读/交易"

	//火币永续合约
	//market := Market{Symbol: "DOGE-USD"}

	//OKEX永续合约
	//market := Market{Symbol: "BTC-USD-SWAP"}

	//获取资产
	getGetBalance()

	//// 订阅订单薄
	//ws.SubscribeLevel2Snapshots(market, func(ob *OrderBook) {
	//	log.Printf("%#v", ob)
	//})
	//// 订阅成交记录
	//ws.SubscribeTrades(market, func(trades []*Trade) {
	//	log.Printf("%#v", trades)
	//})
	//// 订阅订单成交信息
	//ws.SubscribeOrders(market, func(orders []*Order) {
	//	log.Printf("%#v", orders)
	//})
	// 订阅持仓信息
	//ws.SubscribePositions(market, func(positions []*Position) {
	//	log.Printf("%#v", positions)
	//})

	select {}
}
