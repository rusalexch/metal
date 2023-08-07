package grpcserver

import (
	"context"
	"log"
	"net"

	"github.com/rusalexch/metal/internal/app"
	pm "github.com/rusalexch/metal/internal/proto"
	"google.golang.org/grpc"
)

type MetricService struct {
	pm.UnimplementedMetricsServer
	stor    storage
	address string
	s       *grpc.Server
}

func New(stor storage, address string) *MetricService {
	return &MetricService{
		stor:    stor,
		address: address,
	}
}

func (ms *MetricService) Start() {
	listen, err := net.Listen("tcp", ms.address)
	if err != nil {
		log.Fatal(err)

	}
	s := grpc.NewServer()
	ms.s = s
	pm.RegisterMetricsServer(s, ms)

	log.Println("server gRPC was started")
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}

func (ms *MetricService) Stop() {
	ms.s.GracefulStop()
}

func (ms *MetricService) AddCounter(ctx context.Context, in *pm.CounterMetric) (*pm.Response, error) {
	log.Println(in)
	m := app.Metrics{
		Type:  app.Counter,
		ID:    in.Name,
		Delta: &in.Delta,
	}

	return ms.addMetric(ctx, m)
}

func (ms *MetricService) AddGauge(ctx context.Context, in *pm.GaugeMetric) (*pm.Response, error) {
	m := app.Metrics{
		Type:  app.Gauge,
		ID:    in.Name,
		Value: &in.Value,
	}

	return ms.addMetric(ctx, m)
}

func (ms *MetricService) addMetric(ctx context.Context, m app.Metrics) (*pm.Response, error) {
	err := ms.stor.Add(ctx, m)
	if err != nil {
		return errorResponse("can't save metric"), err
	}

	return okResponse(), nil
}

func okResponse() *pm.Response {
	return &pm.Response{
		Status: pm.Status_OK,
	}
}

func errorResponse(msg string) *pm.Response {
	return &pm.Response{
		Status:  pm.Status_ERROR,
		Message: msg,
	}
}
