package main

import (
	"fmt"
    "log"
    "net/rpc/jsonrpc"
    "os"
	"encoding/json"	
)

//Sample Buy Stock Request: go run client.go localhost:1234 buy "{\"stockSymbolAndPercentage\":\"GOOG:50%,YHOO:50%\",\"budget\":10000.00}"
//Sample Check Portfolio Request: go run client.go localhost:1234 checkPortfolio "{\"tradeid\":294878345}"

type StockRequest struct {
	Budget                   float64    `json:"budget"`
	StockSymbolAndPercentage string `json:"stockSymbolAndPercentage"`
}

type StockResponse struct{
	TradeId uint32 `json:"tradeid"`
	Stocks string `json:"stocks"`
	UnvestedAmount float64 `json:"unvestedAmount"`
}

type PortfolioRequest struct {
	Tradeid uint32 `json:"tradeid"`
}

type PortfolioResponse struct {

	Stocks string `json:"stocks"`
	CurrentMarketValue float64 `json:"currentMarketValue"`
	UnvestedAmount float64 `json:"unvestedAmount"`
}

var stockRequest StockRequest
var portfolioRequest PortfolioRequest

func main() {
    if len(os.Args) != 4 {
        fmt.Println("Usage: ", os.Args[0], "server:port")
        log.Fatal(1)
    }
    service := os.Args[1]

    client, err := jsonrpc.Dial("tcp", service)
    if err != nil {
        log.Fatal("dialing:", err)
    }
    
	if os.Args[2] == "buy"	{
		fmt.Printf("Buying Stocks..\n ")
		content := []byte(os.Args[3])
		err = json.Unmarshal(content, &stockRequest)
		var reply StockResponse
		err = client.Call("StockMarket.BuyStock", stockRequest, &reply)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\n%+v\n",reply)

	}	else if os.Args[2] == "checkPortfolio"	{
		fmt.Printf("Checking Portfolio.. \n")
		content := []byte(os.Args[3])
		err = json.Unmarshal(content, &portfolioRequest)
		var reply PortfolioResponse
		//tradeId,_ := strconv.ParseInt(os.Args[3],10,64)
		err = client.Call("StockMarket.CheckPortfolio", portfolioRequest , &reply)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\n%+v\n",reply)
		
	}	else	{
		
		fmt.Printf("Invalid Input")
	}
}