package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	authv1 "github.com/qkitzero/auth-service/gen/go/auth/v1"
	combinationv1 "github.com/qkitzero/combination-service/gen/go/combination/v1"
	appcombination "github.com/qkitzero/combination-service/internal/application/combination"
	apiauth "github.com/qkitzero/combination-service/internal/infrastructure/api/auth"
	infracategory "github.com/qkitzero/combination-service/internal/infrastructure/category"
	"github.com/qkitzero/combination-service/internal/infrastructure/db"
	infraelement "github.com/qkitzero/combination-service/internal/infrastructure/element"
	grpccombination "github.com/qkitzero/combination-service/internal/interface/grpc/combination"
	"github.com/qkitzero/combination-service/util"
)

func main() {
	db, err := db.Init(
		util.GetEnv("DB_HOST", ""),
		util.GetEnv("DB_USER", ""),
		util.GetEnv("DB_PASSWORD", ""),
		util.GetEnv("DB_NAME", ""),
		util.GetEnv("DB_PORT", ""),
		util.GetEnv("DB_SSL_MODE", ""),
	)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", ":"+util.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}

	authTarget := util.GetEnv("AUTH_SERVICE_HOST", "") + ":" + util.GetEnv("AUTH_SERVICE_PORT", "")

	var opts grpc.DialOption
	switch util.GetEnv("ENV", "development") {
	case "production":
		opts = grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))
	default:
		opts = grpc.WithTransportCredentials(insecure.NewCredentials())
	}

	conn, err := grpc.NewClient(authTarget, opts)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	server := grpc.NewServer()

	authServiceClient := authv1.NewAuthServiceClient(conn)

	elementRepository := infraelement.NewElementRepository(db)
	categoryRepository := infracategory.NewCategoryRepository(db)

	authUsecase := apiauth.NewAuthUsecase(authServiceClient)
	combinationUsecase := appcombination.NewCombinationUsecase(elementRepository, categoryRepository)

	healthServer := health.NewServer()
	combinationHandler := grpccombination.NewCombinationHandler(authUsecase, combinationUsecase)

	grpc_health_v1.RegisterHealthServer(server, healthServer)
	combinationv1.RegisterCombinationServiceServer(server, combinationHandler)

	healthServer.SetServingStatus("combination", grpc_health_v1.HealthCheckResponse_SERVING)

	if util.GetEnv("ENV", "development") == "development" {
		reflection.Register(server)
	}

	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
