package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/RakhatLukum/CodeMart/cart-service/config"
	service "github.com/RakhatLukum/CodeMart/cart-service/internal/adapter/grpc"
	"github.com/RakhatLukum/CodeMart/cart-service/internal/adapter/inmemory"
	"github.com/RakhatLukum/CodeMart/cart-service/internal/adapter/mailer"
	natsPublisher "github.com/RakhatLukum/CodeMart/cart-service/internal/adapter/nats"
	redisAdapter "github.com/RakhatLukum/CodeMart/cart-service/internal/adapter/redis"
	"github.com/RakhatLukum/CodeMart/cart-service/internal/repository"
	"github.com/RakhatLukum/CodeMart/cart-service/internal/usecase"
	"github.com/RakhatLukum/CodeMart/cart-service/pkg/mysql"
	natsClient "github.com/RakhatLukum/CodeMart/cart-service/pkg/nats"
	"github.com/RakhatLukum/CodeMart/cart-service/pkg/redis"
	mailjetAPI "github.com/mailjet/mailjet-apiv3-go/v4"
)

const serviceName = "cart-service"

type App struct {
	grpcServer    *service.GRPCServer
	mysqlDB       *mysql.DB
	redisClient   *redis.Client
	mailjetClient *mailer.MailjetClient
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	log.Printf("starting %v service", serviceName)

	log.Println("connecting to MySQL", "database", cfg.MySQL.Database)
	mysqlDB, err := mysql.NewDB(mysql.Config{
		DSN:      cfg.MySQL.DSN,
		Username: cfg.MySQL.User,
		Password: cfg.MySQL.Password,
		Host:     cfg.MySQL.Host,
		Port:     cfg.MySQL.Port,
		Database: cfg.MySQL.Database,
	})
	if err != nil {
		return nil, fmt.Errorf("mysql: %w", err)
	}

	log.Println("connecting to Redis", "addr", cfg.Redis.Addr)
	redisClientInstance, err := redis.NewClient(redis.Config{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	if err != nil {
		return nil, fmt.Errorf("redis: %w", err)
	}

	log.Println("connecting to NATS", "url", cfg.NATS.URL)
	natsClientInstance, err := natsClient.NewClient(cfg.NATS.URL)
	if err != nil {
		return nil, fmt.Errorf("nats: %w", err)
	}
	natsPublishers := natsPublisher.NewPublisher(natsClientInstance.Conn)

	log.Println("initializing Mailjet client")
	mailjetClientInstance := mailjetAPI.NewMailjetClient(cfg.Mailjet.APIKey, cfg.Mailjet.SecretKey)
	mailjetAdapter := mailer.NewMailjetClient(mailjetClientInstance, cfg.Mailjet.FromEmail, cfg.Mailjet.FromName)

	inmemoryClient := inmemory.NewClient()
	cartRepo := repository.NewCartRepository(mysqlDB.Conn)
	redisAdapterInstance := redisAdapter.NewClient(redisClientInstance, cfg.Redis.TTL)

	cartUsecase := usecase.NewCartUsecase(cartRepo, redisAdapterInstance, inmemoryClient, mailjetAdapter, natsPublishers)

	grpcServer, err := service.NewGRPCServer(*cfg, cartUsecase)
	if err != nil {
		return nil, fmt.Errorf("grpc server: %w", err)
	}

	return &App{
		grpcServer:    grpcServer,
		mysqlDB:       mysqlDB,
		redisClient:   redisClientInstance,
		mailjetClient: mailjetAdapter,
	}, nil
}

func (a *App) Close() {
	log.Println("closing resources...")

	a.grpcServer.Stop()

	if a.mysqlDB != nil {
		if err := a.mysqlDB.Close(); err != nil {
			log.Printf("failed to close MySQL: %v", err)
		}
	}

	if a.redisClient != nil {
		if err := a.redisClient.Close(); err != nil {
			log.Printf("failed to close Redis: %v", err)
		}
	}

	log.Println("all resources closed")
}

func (a *App) Run() error {
	errCh := make(chan error, 1)

	go func() {
		errCh <- a.grpcServer.Run()
	}()

	log.Printf("service %v started", serviceName)

	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case errRun := <-errCh:
		return fmt.Errorf("grpc server failed: %w", errRun)
	case s := <-shutdownCh:
		log.Printf("received signal: %v. Running graceful shutdown...", s)
		a.Close()
		log.Println("graceful shutdown completed!")
	}

	return nil
}
