package gateway

import (
	"context"

	"github.com/aitsvet/saga_example/internal/model"
	mp "github.com/aitsvet/saga_example/pkg/mp/v1"
	"google.golang.org/protobuf/proto"
)

func (s *Server) CreateAd(ctx context.Context, r *mp.AdRequest) (*mp.Empty, error) {
	return &mp.Empty{}, s.s.Rabbit.Publish(model.NewMessage("ads", r.SellerId, r))
}

func (s *Server) Order(ctx context.Context, r *mp.OrderRequest) (*mp.Empty, error) {
	return &mp.Empty{}, s.s.Rabbit.Publish(model.NewMessage("orders", r.BuyerId, r))
}

func (s *Server) Ship(ctx context.Context, r *mp.ShipmentRequest) (*mp.Empty, error) {
	return &mp.Empty{}, s.s.Rabbit.Publish(model.NewMessage("shipments", r.SellerId, r))
}

func (s *Server) Receive(ctx context.Context, r *mp.DeliveryRequest) (*mp.Empty, error) {
	return &mp.Empty{}, s.s.Rabbit.Publish(model.NewMessage("shipments", r.BuyerId, r))
}

func (s *Server) Update(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case m := <-s.fromDictionary:
			_ = proto.Unmarshal(m.Body, s.dictionary)
		case m := <-s.fromAds:
			ad := &mp.AdResponse{}
			_ = proto.Unmarshal(m.Body, ad)
			s.ads[ad.AdId] = ad
		case m := <-s.fromOrders:
			or := &mp.OrderResponse{}
			_ = proto.Unmarshal(m.Body, or)
			s.orders[or.OrderId] = or
		case m := <-s.fromShipments:
			sh := &mp.ShipmentResponse{}
			_ = proto.Unmarshal(m.Body, sh)
			s.shipments[sh.ShipmentId] = sh
		}
	}
}
