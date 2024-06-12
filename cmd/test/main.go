package main

import (
	"context"
	"log"
	"time"

	mp "github.com/aitsvet/saga_example/pkg/mp/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TestClient struct {
	mp.MarketplaceClient
}

func (c *TestClient) createAdAndWaitForIt(ctx context.Context, adRequest *mp.AdRequest) (*mp.AdResponse, error) {
	_, err := c.CreateAd(ctx, adRequest)
	if err != nil {
		return nil, err
	}
	listAdsRequest := &mp.ListAdsRequest{
		SellerId:  &adRequest.SellerId,
		Substring: &adRequest.ProductName,
	}
	for {
		listResponse, err := c.ListAds(ctx, listAdsRequest)
		if err != nil {
			return nil, err
		}
		for ad, err := listResponse.Recv(); err == nil; ad, err = listResponse.Recv() {
			if ad.Props.SellerId == adRequest.SellerId && ad.Props.ProductName == adRequest.ProductName && ad.Props.Price == adRequest.Price {
				return ad, nil
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func (c *TestClient) createOrderAndWaitForIt(ctx context.Context, orderRequest *mp.OrderRequest) (*mp.OrderResponse, error) {
	_, err := c.Order(ctx, orderRequest)
	if err != nil {
		return nil, err
	}
	listOrdersRequest := &mp.ListOrdersRequest{
		UserId: &mp.ListOrdersRequest_BuyerId{BuyerId: orderRequest.BuyerId},
	}
	for {
		listResponse, err := c.ListOrders(ctx, listOrdersRequest)
		if err != nil {
			return nil, err
		}
		for order, err := listResponse.Recv(); err == nil; order, err = listResponse.Recv() {
			if order.Props.BuyerId == orderRequest.BuyerId && order.Props.AdId == orderRequest.AdId {
				return order, nil
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func (c *TestClient) createShipmentAndWaitForIt(ctx context.Context, shipmentRequest *mp.ShipmentRequest) (*mp.ShipmentResponse, error) {
	_, err := c.Ship(ctx, shipmentRequest)
	if err != nil {
		return nil, err
	}
	listShipmentsRequest := &mp.ListShipmentsRequest{
		UserId: &mp.ListShipmentsRequest_SellerId{SellerId: shipmentRequest.SellerId},
	}
	for {
		listResponse, err := c.ListShipments(ctx, listShipmentsRequest)
		if err != nil {
			return nil, err
		}
		for shipment, err := listResponse.Recv(); err == nil; shipment, err = listResponse.Recv() {
			if shipment.OrderId == shipmentRequest.OrderId {
				return shipment, nil
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func (c *TestClient) receiveShipmentAndWaitForIt(ctx context.Context, deliveryRequest *mp.DeliveryRequest) (*mp.ShipmentResponse, error) {
	_, err := c.Receive(ctx, deliveryRequest)
	if err != nil {
		return nil, err
	}
	listShipmentsRequest := &mp.ListShipmentsRequest{
		UserId: &mp.ListShipmentsRequest_BuyerId{BuyerId: deliveryRequest.BuyerId},
	}
	for {
		listResponse, err := c.ListShipments(ctx, listShipmentsRequest)
		if err != nil {
			return nil, err
		}
		for shipment, err := listResponse.Recv(); err == nil; shipment, err = listResponse.Recv() {
			if shipment.ShipmentId == deliveryRequest.ShipmentId && shipment.Status > 4 {
				return shipment, nil
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ctx := context.Background()
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Не удалось подключиться: %v", err)
	}
	defer conn.Close()
	client := mp.NewMarketplaceClient(conn)
	c := &TestClient{MarketplaceClient: client}

	// Create an ad first
	adRequest := &mp.AdRequest{
		SellerId:    12345,
		ProductName: "Sample Product",
		Price:       1000,
	}
	adResponse, err := c.createAdAndWaitForIt(ctx, adRequest)
	if err != nil {
		log.Fatalf("Ошибка при создании и получении объявления: %v", err)
	}
	log.Printf("Объявление успешно создано и получено: %v", adResponse)

	// Create an order for the created ad
	orderRequest := &mp.OrderRequest{
		AdId:    adResponse.AdId,
		BuyerId: 67890, // Replace with actual buyer ID
	}
	orderResponse, err := c.createOrderAndWaitForIt(ctx, orderRequest)
	if err != nil {
		log.Fatalf("Ошибка при создании и получении заказа: %v", err)
	}
	log.Printf("Заказ успешно создан и получен: %v", orderResponse)

	// Create a shipment for the created ad
	shipmentRequest := &mp.ShipmentRequest{
		OrderId:  orderResponse.OrderId,
		SellerId: adRequest.SellerId, // Replace with actual seller ID
	}
	shipmentResponse, err := c.createShipmentAndWaitForIt(ctx, shipmentRequest)
	if err != nil {
		log.Fatalf("Ошибка при создании отправления: %v", err)
	}
	log.Printf("Отправление успешно создано: %v", shipmentResponse)

	// Receive a shipment for the created ad
	deliveryRequest := &mp.DeliveryRequest{
		ShipmentId: shipmentResponse.ShipmentId,
		BuyerId:    orderRequest.BuyerId,
	}
	deliveryResponse, err := c.receiveShipmentAndWaitForIt(ctx, deliveryRequest)
	if err != nil {
		log.Fatalf("Ошибка при получении отправления: %v", err)
	}
	log.Printf("Отправление успешно получено: %v", deliveryResponse)
}
