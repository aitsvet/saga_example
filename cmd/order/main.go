package main

import (
	"database/sql"

	"github.com/aitsvet/saga_example/internal/model"
	"github.com/aitsvet/saga_example/internal/service/domain"
	"github.com/aitsvet/saga_example/pkg/mp/v1"
	_ "github.com/lib/pq"
	"google.golang.org/protobuf/proto"
)

var migration = `
DROP TABLE IF EXISTS orders;
CREATE TABLE orders (
	order_id SERIAL PRIMARY KEY,
	ad_id BIGINT,
	buyer_id BIGINT,
	status SMALLINT,
	tx_id BIGINT
);
`

var insertOrder = "INSERT INTO orders(ad_id, buyer_id, status, tx_id) VALUES ($1, $2, $3, $4) RETURNING order_id"
var updateShipment = "UPDATE orders SET status = $1, tx_id = $2 WHERE " +
	"status != $1 AND tx_id < $2 AND order_id = $3 RETURNING ad_id, buyer_id"
var updateAd = "UPDATE orders SET status = $1, tx_id = $2 WHERE " +
	"status != $1 AND tx_id < $2 AND ad_id = $3 RETURNING order_id, buyer_id"

func main() {

	s := domain.NewService("orders", migration, "ads", "shipments")
	defer s.Close()

	fromGateway := s.Rabbit.ListenTo("orders")
	fromAd := s.Kafka.ListenTo("ads")
	fromShipment := s.Kafka.ListenTo("shipments")

	for {
		o := &mp.OrderResponse{Props: &mp.OrderRequest{}}
		m := &model.Message{}
		select {
		case m = <-fromGateway:
			_ = proto.Unmarshal(m.Body, o.Props)
			o.Status = 1
			err := s.DB.QueryRow(insertOrder, o.Props.AdId, o.Props.BuyerId, o.Status, m.TxId).
				Scan(&o.OrderId)
			if err != nil {
				s.LogError(m, o.Props, err)
				continue
			}
		case m = <-fromShipment:
			r := &mp.ShipmentResponse{}
			_ = proto.Unmarshal(m.Body, r)
			o.Props = &mp.OrderRequest{}
			err := s.DB.QueryRow(updateShipment, r.Status, m.TxId, r.OrderId).
				Scan(&o.Props.AdId, &o.Props.BuyerId)
			if err != nil {
				s.LogError(m, r, err)
				continue
			}
			o.Status = r.Status
			o.OrderId = r.OrderId
		case m = <-fromAd:
			r := &mp.AdResponse{}
			_ = proto.Unmarshal(m.Body, r)
			o.Props = &mp.OrderRequest{AdId: r.AdId}
			err := s.DB.QueryRow(updateAd, r.Status, m.TxId, r.AdId).
				Scan(&o.OrderId, &o.Props.BuyerId)
			if err != nil {
				if err != sql.ErrNoRows {
					s.LogError(m, r, err)
				}
				continue
			}
			o.Status = r.Status
		}
		body, _ := proto.Marshal(o)
		s.Kafka.Publish(&model.Message{
			Queue: "orders",
			TxId:  m.TxId,
			Body:  body,
		})
	}
}
