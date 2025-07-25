package Routes

import (
	"awesomeProject/Endpoints"
	"awesomeProject/Internal/Api/Handlers"
	"github.com/gin-gonic/gin"
)

func GetRoutes() *gin.Engine {
	routes := gin.Default()
	routesComplete := routes.Group("/api/v1")
	{
		routesComplete.POST(Endpoints.GetWalletBalanceEndpoint, Handlers.PostWalletBalance)
		routesComplete.POST(Endpoints.SendEth, Handlers.SendEth)
		routesComplete.POST(Endpoints.SellEth, Handlers.SellEth)
		routesComplete.GET(Endpoints.EthPrice, Handlers.EthPrice)
		routesComplete.POST(Endpoints.GetWalletBalancePriceEndpoint, Handlers.PostWalletBalancePrice)
		routesComplete.POST(Endpoints.GetTransactionHistory, Handlers.GetWalletTransactions)
		routesComplete.POST(Endpoints.GetNftCollection, Handlers.GetNftCollection)
	}
	return routes
}
