package grpc

import (
	"CodeMart/analytics-service/internal/analytics/usecase"
	"context"
	analyticspb "proto/analytics"

	"google.golang.org/grpc"
)

type server struct {
	analyticspb.UnimplementedAnalyticsServiceServer
	uc *usecase.AnalyticsUsecase
}

func Register(s *grpc.Server, uc *usecase.AnalyticsUsecase) {
	analyticspb.RegisterAnalyticsServiceServer(s, &server{uc: uc})
}

func (s *server) LogProductView(ctx context.Context, req *analyticspb.ProductView) (*analyticspb.Response, error) {
	if err := s.uc.Log(int64(req.ProductId)); err != nil {
		return nil, err
	}
	return &analyticspb.Response{Message: "logged"}, nil
}
func (s *server) GetTopProducts(ctx context.Context, _ *analyticspb.Empty) (*analyticspb.TopProducts, error) {
	ids, err := s.uc.Top5()
	if err != nil {
		return nil, err
	}
	res := &analyticspb.TopProducts{}
	for _, id := range ids {
		res.ProductIds = append(res.ProductIds, int32(id))
	}
	return res, nil
}
