package grpc

import (
	"context"
	"pricing-system-alert-service/domain"
	pb "pricing-system-alert-service/grpc/proto"
)

type PricingAlertServer struct {
	ctx              context.Context
	pricingAlertChan chan *domain.PriceNote
	pb.UnimplementedPricingAlertServer
}

func NewPricingAlertServer(ctx context.Context, alertChan chan *domain.PriceNote) *PricingAlertServer {
	return &PricingAlertServer{
		ctx:              ctx,
		pricingAlertChan: alertChan,
	}
}

func (r *PricingAlertServer) SubscribeOnAlerts(_ *pb.SubscribeOnAlertsParams, stream pb.PricingAlert_SubscribeOnAlertsServer) error {
	for {
		select {
		case <-r.ctx.Done():
			return nil
		case note := <-r.pricingAlertChan:
			grpcNote := &pb.PriceNote{
				UnixTimestamp: note.TimeStamp.Unix(),
				Currency:      note.Currency,
				Price:         note.Price,
			}

			if err := stream.Send(grpcNote); err != nil {
				return err
			}
		}
	}
}
