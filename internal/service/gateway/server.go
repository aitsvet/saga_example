package gateway

import (
	"context"

	"github.com/aitsvet/saga_example/internal/model"
	"github.com/aitsvet/saga_example/internal/service/domain"
	mp "github.com/aitsvet/saga_example/pkg/mp/v1"
)

type Server struct {
	s *domain.Service

	ads        map[uint64]*mp.AdResponse
	orders     map[uint64]*mp.OrderResponse
	shipments  map[uint64]*mp.ShipmentResponse
	dictionary *mp.Dictionary

	fromAds        <-chan *model.Message
	fromOrders     <-chan *model.Message
	fromShipments  <-chan *model.Message
	fromDictionary <-chan *model.Message

	cancel context.CancelFunc

	mp.UnimplementedMarketplaceServer
}

func NewServer() *Server {

	srv := domain.NewService("gateway", "", "ads", "orders", "shipments", "dictionary")
	s := &Server{
		s:              srv,
		ads:            make(map[uint64]*mp.AdResponse),
		orders:         make(map[uint64]*mp.OrderResponse),
		shipments:      make(map[uint64]*mp.ShipmentResponse),
		dictionary:     &mp.Dictionary{},
		fromAds:        srv.Kafka.ListenTo("ads"),
		fromOrders:     srv.Kafka.ListenTo("orders"),
		fromShipments:  srv.Kafka.ListenTo("shipments"),
		fromDictionary: srv.Kafka.ListenTo("dictionary"),
	}

	var ctx context.Context
	ctx, s.cancel = context.WithCancel(context.Background())
	go s.Update(ctx)

	return s
}

func (s *Server) Close() {
	s.cancel()
}
