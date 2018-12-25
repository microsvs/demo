package types

import "github.com/microsvs/base/pkg/rpc"

const FGBMircoService rpc.FGService = 10

const (
	_          rpc.FGService = iota + FGBMircoService + rpc.FGBService
	FGSMonitor               // 8080+10+1 = 8091
)

func init() {
	var err error
	// service register, such monitor
	if err = rpc.RegisterService(FGSMonitor, "monitor"); err != nil {
		panic(err)
	}
}
