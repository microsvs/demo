package main

import (
	"fmt"

	"github.com/microsvs/base"
	pdb "github.com/microsvs/base/cmd/db"
	"github.com/microsvs/base/pkg/rpc"
)

func init() {
	pdb.InitDB(rpc.FGSAddress)
}
func main() {
	fmt.Printf("port=%d\n", rpc.FGSAddress)
	d, err := base.NewGLDaemon(rpc.FGSAddress, &schema)
	if err != nil {
		panic(err)
	}
	d.Listen()
}
