package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	ep "orders/internal/endpoint"
	"orders/internal/entity"
	"orders/internal/handler"
	"orders/internal/service"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/joho/godotenv"
)

func main() {
	var (
		httpAddr = flag.String("port", ":8080", "HTTP listen address")
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

	conn := entity.NewConnection(os.Getenv("ORDERS_DB_USERNAME"), os.Getenv("ORDERS_DB_PASSWORD"), os.Getenv("ORDERS_DB_NAME"), logger)
	conn.DB.AutoMigrate(&entity.Order{})
	conn.DB.AutoMigrate(&entity.OrderItem{})

	orderRepository := entity.NewOrderRepository(conn)
	tlogger := service.NewLogService(logger)
	kafkaService := service.NewKafkaService()
	var orderService service.OrderService
	{
		orderService = service.NewOrderService(orderRepository, logger, kafkaService)
		orderService = handler.NewLoggingOrderServiceMiddleware(logger, tlogger)(orderService)
	}
	orderItemRepository := entity.NewOrderItemRepository(conn)
	var orderItemService service.OrderItemService
	{
		orderItemService = service.NewOrderItemService(orderItemRepository, logger)
	}
	endpoints := make(map[string]endpoint.Endpoint)
	endpoints = ep.MakeOrderEndpoints(orderService, endpoints)
	endpoints = ep.MakeOrderItemEndpoints(orderItemService, endpoints)
	httpHandler := handler.MakeHTTPHandler(orderService, orderItemService, endpoints, logger)

	/* listener := make(chan services.Payload)
	services.StartKafkaListener(context.Background(), listener)
	logger.Log("kafka listener", "starting in goroutine...")
	go func() {
		for p := range listener {
			logger.Log("listen:", p.Name)
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
