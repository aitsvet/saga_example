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
DROP TABLE IF EXISTS ads;
CREATE TABLE ads (
	ad_id SERIAL PRIMARY KEY,
	product_name VARCHAR(255),
	price BIGINT,
	seller_id BIGINT,
	status SMALLINT,
	tx_id BIGINT
);
`

var insertAd = "INSERT INTO ads(product_name, price, seller_id, status, tx_id) VALUES ($1, $2, $3, $4, $5) RETURNING ad_id, status"
var updateOrder = "UPDATE ads SET status = $1, tx_id = $2 WHERE " +
	"status != $1 AND tx_id < $2 AND ad_id = $3 RETURNING product_name, price, seller_id"

func main() {

	s := domain.NewService("ads", migration, "orders")
	defer s.Close()

	fromGateway := s.Rabbit.ListenTo("ads")
	fromOrder := s.Kafka.ListenTo("orders")

	for {
		o := &mp.AdResponse{Props: &mp.AdRequest{}}
		m := &model.Message{}
		select {
		case m = <-fromGateway:
			_ = proto.Unmarshal(m.Body, o.Props)
			err := s.DB.QueryRow(insertAd, o.Props.ProductName, o.Props.Price, o.Props.SellerId, 1, m.TxId).
				Scan(&o.AdId, &o.Status)
			if err != nil {
				s.LogError(m, o.Props, err)
				continue
			}
		case m = <-fromOrder:
			r := &mp.OrderResponse{}
			_ = proto.Unmarshal(m.Body, r)
			err := s.DB.QueryRow(updateOrder, r.Status, m.TxId, r.Props.AdId).
				Scan(&o.Props.ProductName, &o.Props.Price, &o.Props.SellerId)
			if err != nil {
				if err != sql.ErrNoRows {
					s.LogError(m, r, err)
				}
				continue
			}
			o.AdId = r.Props.AdId
		}
		body, _ := proto.Marshal(o)
		s.Kafka.Publish(&model.Message{
			Queue: "ads",
			TxId:  m.TxId,
			Body:  body,
		})
	}
}
