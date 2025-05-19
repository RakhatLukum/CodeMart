package app

import (
	stdContext "context"
	stdFmt "fmt"
	stdLog "log"
	stdOs "os"
	stdSignal "os/signal"
	stdSyscall "syscall"

	configLoader "github.com/RakhatLukum/CodeMart/analytics-service/config"
	viewRepository "github.com/RakhatLukum/CodeMart/analytics-service/internal/repository"
	viewUsecase "github.com/RakhatLukum/CodeMart/analytics-service/internal/usecase"

	grpcServerAdapter "github.com/RakhatLukum/CodeMart/analytics-service/internal/adapter/grpc"
	inmemoryCacheAdapter "github.com/RakhatLukum/CodeMart/analytics-service/internal/adapter/inmemory"
	mailjetEmailAdapter "github.com/RakhatLukum/CodeMart/analytics-service/internal/adapter/mailer"
	natsSubscriberAdapter "github.com/RakhatLukum/CodeMart/analytics-service/internal/adapter/nats"
	redisCacheAdapter "github.com/RakhatLukum/CodeMart/analytics-service/internal/adapter/redis"

	mysqlClient "github.com/RakhatLukum/CodeMart/analytics-service/pkg/mysql"
	natsClient "github.com/RakhatLukum/CodeMart/analytics-service/pkg/nats"
	redisClient "github.com/RakhatLukum/CodeMart/analytics-service/pkg/redis"

	mailjetAPI "github.com/mailjet/mailjet-apiv3-go/v4"
)

const serviceName = "analytics-service"

type App struct {
	grpcServer    *grpcServerAdapter.GRPCServer
	natsClient    *natsClient.Client
	natsSub       *natsSubscriberAdapter.Subscriber
	mysqlDB       *mysqlClient.DB
	redisClient   *redisClient.Client
	mailjetClient *mailjetEmailAdapter.MailjetClient
}

func New(ctx stdContext.Context, cfg *configLoader.Config) (*App, error) {
	stdLog.Printf("starting %v service", serviceName)

	stdLog.Println("connecting to MySQL", "database", cfg.MySQL.Database)
	mysqlDB, err := mysqlClient.NewDB(mysqlClient.Config{
		DSN:      cfg.MySQL.DSN,
		Username: cfg.MySQL.User,
		Password: cfg.MySQL.Password,
		Host:     cfg.MySQL.Host,
		Port:     cfg.MySQL.Port,
		Database: cfg.MySQL.Database,
	})
	if err != nil {
		return nil, stdFmt.Errorf("mysql: %w", err)
	}

	stdLog.Println("connecting to Redis", "addr", cfg.Redis.Addr)
	redisClientInstance, err := redisClient.NewClient(redisClient.Config{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	if err != nil {
		return nil, stdFmt.Errorf("redis: %w", err)
	}

	stdLog.Println("connecting to NATS", "url", cfg.NATS.URL)
	natsClientInstance, err := natsClient.NewClient(cfg.NATS.URL)
	if err != nil {
		return nil, stdFmt.Errorf("nats: %w", err)
	}

	stdLog.Println("initializing Mailjet client")
	mailjetClientInstance := mailjetAPI.NewMailjetClient(cfg.Mailjet.APIKey, cfg.Mailjet.SecretKey)
	mailjetAdapter := mailjetEmailAdapter.NewMailjetClient(mailjetClientInstance, cfg.Mailjet.FromEmail, cfg.Mailjet.FromName)

	inmemoryClient := inmemoryCacheAdapter.NewClient()
	viewRepo := viewRepository.NewViewRepository(mysqlDB.Conn)

	redisCacheUsecase := viewUsecase.NewViewCacheUsecase(redisCacheAdapter.NewClient(redisClientInstance, cfg.Redis.TTL))
	memoryCacheUsecase := viewUsecase.NewViewMemoryUsecase(inmemoryClient)
	viewUsecaseInstance := viewUsecase.NewViewUsecase(viewRepo, redisCacheUsecase, memoryCacheUsecase, mailjetAdapter)

	natsSubscriber := natsSubscriberAdapter.NewSubscriber(
		natsClientInstance.Conn,
		viewUsecaseInstance,
		"products",
		"users",
		"carts",
	)
	if err := natsSubscriber.Subscribe(); err != nil {
		return nil, stdFmt.Errorf("nats subscribe: %w", err)
	}

	grpcServer, err := grpcServerAdapter.NewGRPCServer(*cfg, viewUsecaseInstance, redisCacheUsecase, memoryCacheUsecase)
	if err != nil {
		return nil, stdFmt.Errorf("grpc server: %w", err)
	}

	return &App{
		grpcServer:    grpcServer,
		natsClient:    natsClientInstance,
		natsSub:       natsSubscriber,
		mysqlDB:       mysqlDB,
		redisClient:   redisClientInstance,
		mailjetClient: mailjetAdapter,
	}, nil
}

func (a *App) Close() {
	stdLog.Println("closing resources...")

	a.grpcServer.Stop()

	if a.natsClient != nil {
		a.natsClient.Close()
	}

	if a.mysqlDB != nil {
		if err := a.mysqlDB.Close(); err != nil {
			stdLog.Printf("failed to close MySQL: %v", err)
		}
	}

	if a.redisClient != nil {
		if err := a.redisClient.Close(); err != nil {
			stdLog.Printf("failed to close Redis: %v", err)
		}
	}

	stdLog.Println("all resources closed")
}

func (a *App) Run() error {
	errCh := make(chan error, 1)

	go func() {
		errCh <- a.grpcServer.Run()
	}()

	stdLog.Printf("service %v started", serviceName)

	shutdownCh := make(chan stdOs.Signal, 1)
	stdSignal.Notify(shutdownCh, stdSyscall.SIGINT, stdSyscall.SIGTERM)

	select {
	case errRun := <-errCh:
		return stdFmt.Errorf("grpc server failed: %w", errRun)
	case s := <-shutdownCh:
		stdLog.Printf("received signal: %v. Running graceful shutdown...", s)
		a.Close()
		stdLog.Println("graceful shutdown completed!")
	}

	return nil
}
