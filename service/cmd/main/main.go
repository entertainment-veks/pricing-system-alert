package main

import (
	"context"
	"log"
	"net"
	"pricing-system-alert-service/config"
	rpc "pricing-system-alert-service/delivery/grpc"
	"pricing-system-alert-service/domain"
	pb "pricing-system-alert-service/grpc/proto"
	"pricing-system-alert-service/repository"
	"pricing-system-alert-service/service"
	"time"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.NewDefaultConfig()
	ctx := context.Background()
	db := repository.NewDataBase()

	listen, err := net.Listen("tcp", cfg.GRPCServerAddr)
	if err != nil {
		log.Println(err)
	}

	server := grpc.NewServer()
	alertChan := make(chan *domain.PriceNote)
	pricingAlertServer := rpc.NewPricingAlertServer(ctx, alertChan)
	pb.RegisterPricingAlertServer(server, pricingAlertServer)

	go func(alertChan chan *domain.PriceNote, db domain.DataBase) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if err := service.FetchAndSaveData(db); err != nil {
					log.Println(err)
				}

				if err := service.ReadAndSendLastData(alertChan, db); err != nil {
					log.Println(err)
				}

				time.Sleep(1 * time.Minute)
			}
		}
	}(alertChan, db)

	log.Println("Service launched and ready to accept connections")
	if err := server.Serve(listen); err != nil {
		log.Println(err)
	}
}
