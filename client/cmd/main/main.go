package main

import (
	"context"
	"log"
	"pricing-system-alert-client/config"
	"pricing-system-alert-client/service"
	"time"
)

func main() {
	cfg := config.NewDefaultConfig()
	ctx := context.Background()

	for {
		select {
		case <-ctx.Done():
			return
		default: // it's a reconecting mechanism
			consumer, err := service.NewPricesAlertConsumer(ctx, cfg)
			if err != nil {
				log.Println(err)
				time.Sleep(5 * time.Second)
				continue
			}

			log.Println("Starting listening prices alarms...")
			if err := consumer.StartPrintingPriceAlerts(ctx); err != nil {
				log.Println(err)
				time.Sleep(5 * time.Second)
				continue
			}

		}
	}
}
