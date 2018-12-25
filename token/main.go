package main

import (
	"fmt"

	"net/http"
	_ "net/http/pprof"

	"github.com/microsvs/base"
	pcache "github.com/microsvs/base/cmd/cache"
	pdb "github.com/microsvs/base/cmd/db"
	"github.com/microsvs/base/pkg/rpc"
)

func init() {
	pdb.InitDB(rpc.FGSToken)
	pcache.InitCache(rpc.FGSToken)
}

func main() {
	fmt.Printf(":port=%d\n", rpc.FGSToken)
	d, err := base.NewGLDaemon(rpc.FGSToken, &schema)
	if err != nil {
		panic(err)
	}
	go http.ListenAndServe(":6060", nil)
	d.Listen()
}
