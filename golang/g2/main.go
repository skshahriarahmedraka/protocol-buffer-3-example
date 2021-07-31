package main 

import (
	"fmt"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	proto "g2/proto/currency.pb.go"
	"g2/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

)


func main() {
	log := hclog.Default()
	gs := grpc.NewServer(log)

	proto.RegisterCurrencyService(gs ,c)

	reflection.Register(gs)

	l,err:=net.Listen("tcp",fmt.Sprintf(":%d",9092))
	if err!= nil {
		log.Error("unable to create listener"," error : ",err)

	}
	gs.Serve(l)


}