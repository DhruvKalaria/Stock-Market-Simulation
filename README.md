# Stock Market Simulation
Implemented using Go RPC

Sample Arguments:
Buy Stock Request: go run client.go localhost:1234 buy "{\"stockSymbolAndPercentage\":\"GOOG:50%,YHOO:50%\",\"budget\":10000.00}"

Check Portfolio Request: 
go run client.go localhost:1234 checkPortfolio "{\"tradeid\":294878345}"
