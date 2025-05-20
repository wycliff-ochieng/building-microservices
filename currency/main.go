package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	protos "github.com/wycliff-ochieng/currency/protos/currency"
	"github.com/wycliff-ochieng/currency/server"
	"google.golang.org/grpc"
)

func main() {

	log := hclog.Default()

	gs := grpc.NewServer()

	cs := server.NewCurrency(log)

	protos.RegisterCurrencyServer(gs, cs)

	ls, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen to that network", err)
	}
	os.Exit(1)

	gs.Serve(ls)

}
