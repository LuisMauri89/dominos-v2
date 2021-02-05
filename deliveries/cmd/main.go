package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	ep "deliveries/internal/endpoint"
	"deliveries/internal/entity"
	"deliveries/internal/handler"
	"deliveries/internal/service"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/joho/godotenv"
)

func main() {
	var (
		httpAddr = flag.String("port", ":8082", "HTTP listen address")
	)
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	err := godotenv.Load(os.ExpandEnv("../.env"))
	if err != nil {
		panic("missing os environment vars.")
	}

	conn := entity.NewConnection(os.Getenv("DELIVERIES_DB_USERNAME"), os.Getenv("DELIVERIES_DB_PASSWORD"), os.Getenv("DELIVERIES_DB_NAME"), logger)
	conn.DB.AutoMigrate(&entity.Delivery{})
	dRepository := entity.NewDeliveryRepository(conn)
	tlogger := service.NewLogService(logger)
	var dService service.DeliveryService
	{
		dService = service.NewDeliveryService(dRepository, logger, tlogger)
		dService = handler.NewLoggingDeliveryMiddleware(logger)(dService)
	}
	endpoints := make(map[string]endpoint.Endpoint)
	endpoints = ep.MakeDeliveriesEndpoints(dService, endpoints)
	httpHandler := handler.MakeHTTPHandler(dService, endpoints, logger)

	listener := make(chan service.Payload)
	kafkaService := service.NewKafkaService()
	kafkaService.StartKafkaListener(context.Background(), listener)
	logger.Log("kafka listener", "starting in goroutine...")
	go func() {
		for p := range listener {
			dService.Create(context.Background(), entity.Delivery{
				OrderID:     p.ID,
				Status:      "PENDING",
				Name:        p.Name,
				FinalPrice:  p.TotalPrice,
				Address:     p.Address,
				Description: "Products description...",
			})
		}
	}()

	errors := make(chan error)
	go func() {
		osSignal := make(chan os.Signal)
		signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)
		errors <- fmt.Errorf("%s", <-osSignal)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errors <- http.ListenAndServe(*httpAddr, httpHandler)
	}()

	logger.Log("exit", <-errors)
}
