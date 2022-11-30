package service

import (
	"context"
	"log"
	"pricing-system-alert-client/config"
	pb "pricing-system-alert-client/grpc/proto"

	"google.golang.org/grpc"
)

type PricesAlertConsumer struct {
	oldPrices    map[string]float64 // [currency]= oldPrice
	alertsStream pb.PricingAlert_SubscribeOnAlertsClient
}

func NewPricesAlertConsumer(ctx context.Context, cfg *config.Config) (*PricesAlertConsumer, error) {
	conn, err := grpc.Dial(cfg.FetcherServiceAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewPricingAlertClient(conn)

	alertsStream, err := client.SubscribeOnAlerts(ctx, &pb.SubscribeOnAlertsParams{})
	if err != nil {
		return nil, err
	}

	return &PricesAlertConsumer{
		oldPrices:    make(map[string]float64),
		alertsStream: alertsStream,
	}, nil
}

func (r *PricesAlertConsumer) StartPrintingPriceAlerts(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			alert, err := r.alertsStream.Recv()
			if err != nil {
				return err
			}

			log.Printf("Now alert income, currency: %s, old price: %f, new price: %f, change: %f%%\n",
				alert.Currency,
				r.oldPrices[alert.Currency],
				alert.Price,
				percentageChange(r.oldPrices[alert.Currency], alert.Price),
			)

			r.oldPrices[alert.Currency] = alert.Price
		}
	}
}
