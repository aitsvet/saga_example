package main

import (
	"database/sql"
	"strings"

	"github.com/aitsvet/saga_example/internal/model"
	"github.com/aitsvet/saga_example/internal/service/domain"
	"github.com/aitsvet/saga_example/pkg/mp/v1"
	_ "github.com/lib/pq"
	"google.golang.org/protobuf/proto"
)

var migration = `
DROP TABLE IF EXISTS shipments;
CREATE TABLE shipments (
	shipment_id SERIAL PRIMARY KEY,
	order_id BIGINT,
	status SMALLINT,
	tx_id BIGINT
);
`

var insertShipment = "INSERT INTO shipments(order_id, status, tx_id) VALUES ($1, $2, $3) RETURNING shipment_id"
var updateShipment = "UPDATE shipments SET status = $1, tx_id = $2 WHERE " +
	"status != $1 AND tx_id < $2 AND shipment_id = $3 RETURNING order_id"
var updateOrder = "UPDATE shipments SET status = $1, tx_id = $2 WHERE " +
	"status != $1 AND tx_id < $2 AND order_id = $3 RETURNING shipment_id"

func main() {

	s := domain.NewService("shipments", migration, "orders")
	defer s.Close()

	fromGateway := s.Rabbit.ListenTo("shipments")
	fromOrder := s.Kafka.ListenTo("orders")

	for {
		o := &mp.ShipmentResponse{}
		m := &model.Message{}
		select {
		case m = <-fromGateway:
			if strings.HasSuffix(m.Name, "ShipmentRequest") {
				r := &mp.ShipmentRequest{}
				_ = proto.Unmarshal(m.Body, r)
				o.Status = 4
				o.OrderId = r.OrderId
				err := s.DB.QueryRow(insertShipment, o.OrderId, o.Status, m.TxId).
					Scan(&o.ShipmentId)
				if err != nil {
					s.LogError(m, r, err)
					continue
				}
			} else {
				r := &mp.DeliveryRequest{}
				_ = proto.Unmarshal(m.Body, r)
				o.Status = 5
				o.ShipmentId = r.ShipmentId
				err := s.DB.QueryRow(updateShipment, o.Status, m.TxId, o.ShipmentId).
					Scan(&o.OrderId)
				if err != nil {
					s.LogError(m, r, err)
					continue
				}
			}
		case m = <-fromOrder:
			r := &mp.OrderResponse{}
			_ = proto.Unmarshal(m.Body, r)
			o.Status = r.Status
			o.OrderId = r.OrderId
			err := s.DB.QueryRow(updateOrder, o.Status, m.TxId, o.OrderId).
				Scan(&o.ShipmentId)
			if err != nil {
				if err != sql.ErrNoRows {
					s.LogError(m, r, err)
				}
				continue
			}
		}
		body, _ := proto.Marshal(o)
		s.Kafka.Publish(&model.Message{
			Queue: "shipments",
			TxId:  m.TxId,
			Body:  body,
		})
	}
}
