package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	combinationv1 "github.com/qkitzero/combination-service/gen/go/combination/v1"
	appcombination "github.com/qkitzero/combination-service/internal/application/combination"
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

	server := grpc.NewServer()

	elementRepository := infraelement.NewElementRepository(db)

	combinationUsecase := appcombination.NewCombinationUsecase(elementRepository)

	healthServer := health.NewServer()
	combinationHandler := grpccombination.NewCombinationHandler(combinationUsecase)

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
