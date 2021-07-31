package main

import (
	"context"
	proto "g2/proto"

	"github.com/hashicorp/go-hclog"
	
)

type Currency struct {
	log hclog.Logger
}

func (c *Currency)GetRate(ctx context.Context, rr *proto.RateRequest)(*proto.RateResponse , error ){
	c.log.Info("handle request for getRate Base",rr.GetBase(), " dest ",rr.GetDestination())
	return &proto.RateResponse{Rate : 0.5},nil 
}