# cmpe273-Assignment1
Stock Market Simulation
Client and Server delevoped in Go.

Sample Arguments:
_Sample Buy Stock Request: go run client.go localhost:1234 buy "{\"stockSymbolAndPercentage\":\"GOOG:50%,YHOO:50%\",\"budget\":10000.00}"

_Sample Check Portfolio Request: go run client.go localhost:1234 checkPortfolio "{\"tradeid\":294878345}"
