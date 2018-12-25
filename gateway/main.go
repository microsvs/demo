package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/microsvs/base"
	"github.com/microsvs/base/cmd/cache"
	"github.com/microsvs/base/pkg/rpc"
)

func init() {
	cache.InitCache(rpc.FGSUser)
}

func main() {
	fmt.Printf(":port=%d\n", rpc.FGSGateway)
	d, err := base.NewGLDaemon(rpc.FGSGateway, &schema)
	if err != nil {
		panic(err)
	}
	go func() {
		fmt.Println("pprof port: 6060")
		http.ListenAndServe(":6060", nil)
	}()
	//	d.BeforeRouter(base.FilterLimitRateHandler)
	d.Listen()
}
