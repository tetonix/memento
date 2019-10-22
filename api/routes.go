package api

import (
	"net/http"

	"github.com/gin-contrib/static"

	"github.com/gin-gonic/gin"
)

func (a *API) setRoutes() {
	explorer := a.engine.Group("/explorer")
	explorer.GET("/block/:block", a.BlockHandler)
	explorer.GET("/block-range/:start/:end", a.BlockRangeHandler)
	explorer.GET("/uncle/:hash", a.UncleDetailsHandler)
	explorer.GET("/tx/:txHash", a.TxDetailsHandler)
	explorer.GET("/tx/:txHash/log-entries", a.TxLogEntriesHandler)
	explorer.GET("/account/:address/txs", a.AccountTxsHandler)

	a.engine.LoadHTMLGlob("web/templates/*")

	a.engine.Use(static.Serve("/web/assets", static.LocalFile("web/assets", false)))

	a.engine.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
}
