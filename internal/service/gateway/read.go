package gateway

import (
	"context"
	"strings"

	mp "github.com/aitsvet/saga_example/pkg/mp/v1"
)

func (s *Server) ListAds(r *mp.ListAdsRequest, srv mp.Marketplace_ListAdsServer) error {
	for _, a := range s.ads {
		if r.SellerId == nil || a.Props.SellerId == *r.SellerId {
			if r.Substring == nil ||
				strings.Contains(a.Props.ProductName, *r.Substring) {
				srv.Send(a)
			}
		}
	}
	return nil
}

func (s *Server) ListOrders(r *mp.ListOrdersRequest, srv mp.Marketplace_ListOrdersServer) error {
	switch r.UserId.(type) {
	case *mp.ListOrdersRequest_BuyerId:
		for _, o := range s.orders {
			if o.Props.BuyerId == r.GetBuyerId() {
				srv.Send(o)
			}
		}
	case *mp.ListOrdersRequest_SellerId:
		for _, o := range s.orders {
			if a := s.ads[o.Props.AdId]; a != nil &&
				a.Props.SellerId == r.GetSellerId() {
				srv.Send(o)
			}
		}
	}
	return nil
}

func (s *Server) ListShipments(r *mp.ListShipmentsRequest, srv mp.Marketplace_ListShipmentsServer) error {
	switch r.UserId.(type) {
	case *mp.ListShipmentsRequest_BuyerId:
		for _, sh := range s.shipments {
			if o := s.orders[sh.OrderId]; o != nil {
				if o.Props.BuyerId == r.GetBuyerId() {
					srv.Send(sh)
				}
			}
		}
	case *mp.ListShipmentsRequest_SellerId:
		for _, sh := range s.shipments {
			if o := s.orders[sh.OrderId]; o != nil {
				if a := s.ads[o.Props.AdId]; a != nil &&
					a.Props.SellerId == r.GetSellerId() {
					srv.Send(sh)
				}
			}
		}
	}
	return nil
}

func (s *Server) ListStatuses(ctx context.Context, r *mp.Empty) (*mp.Dictionary, error) {
	return s.dictionary, nil
}
