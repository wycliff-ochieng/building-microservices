package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	protos "github.com/wycliff-ochieng/currency/protos/currency"
)

type Currency struct {
	log hclog.Logger
}

func NewCurrency(log hclog.Logger) *Currency {
	return &Currency{log}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {

	c.log.Info("Handle Get", "base", rr.GetBase(), "destination:", rr.GetDestination())

	return &protos.RateResponse{"rate": 0.5}, nil
}
