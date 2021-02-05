package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dominos/deliveries"
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

	conn := deliveries.NewConnection(os.Getenv("DELIVERIES_DB_USERNAME"), os.Getenv("DELIVERIES_DB_PASSWORD"), os.Getenv("DELIVERIES_DB_NAME"), logger)
	defer conn.DB.Close()
	dRepository := deliveries.NewDeliveryRepository(conn)
	tlogger := deliveries.NewLogService(logger)
	var dService deliveries.DeliveryService
	{
		dService = deliveries.NewDeliveryService(dRepository, logger, tlogger)
		dService = deliveries.NewLoggingDeliveryMiddleware(logger)(dService)
	}
	httpHandler := deliveries.MakeHTTPHandler(dService, logger)

	/*listener := make(chan deliveries.Payload)
	kafkaService := deliveries.NewKafkaService()
	kafkaService.StartKafkaListener(context.Background(), listener)
	logger.Log("kafka listener", "starting in goroutine...")
	go func() {
		for p := range listener {
			dService.Create(context.Background(), deliveries.Delivery{
				OrderID:     p.ID,
				Status:      "PENDING",
				Name:        p.Name,
				FinalPrice:  p.TotalPrice,
				Address:     p.Address,
				Description: "Products description...",
			})
		}
	}() */

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
