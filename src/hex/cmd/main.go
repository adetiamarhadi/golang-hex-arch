package main

import (
	"log"
	"os"

	"github.com/adetiamarhadi/golang-hex-arch/internal/adapters/app/api"
	"github.com/adetiamarhadi/golang-hex-arch/internal/adapters/core/arithmetic"
	"github.com/adetiamarhadi/golang-hex-arch/internal/adapters/framework/right/db"
	"github.com/adetiamarhadi/golang-hex-arch/internal/ports"

	gRPC "github.com/adetiamarhadi/golang-hex-arch/internal/adapters/framework/left/grpc"
)

func main() {

	var err error

	var dbaseAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("Failed to initiate dbase connection: %v", err)
	}
	defer dbaseAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(dbaseAdapter, core)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)

	gRPCAdapter.Run()
}