package http

import (
	"net/http"

	"go-common/app/service/main/feed/conf"
	"go-common/app/service/main/feed/service"
	"go-common/library/log"
	bm "go-common/library/net/http/blademaster"
)

var fSrv *service.Service

// Init init http service
func Init(c *conf.Config, srv *service.Service) {
	fSrv = srv
	// init outer router
	engineOuter := bm.DefaultServer(c.BM)
	outerRouter(engineOuter)
	if err := engineOuter.Start(); err != nil {
		log.Error("bm.DefaultServer error(%v)", err)
		panic(err)
	}
}

// outerRouter init outer router
func outerRouter(r *bm.Engine) {
	r.Ping(ping)
	r.Register(register)
}

func ping(c *bm.Context) {
	if err := fSrv.Ping(c); err != nil {
		log.Error("ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

func register(c *bm.Context) {
	c.JSON(map[string]interface{}{}, nil)
}
