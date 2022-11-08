package grpc

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	emclient "sbit-emitter/adapter/emitter"
	"sbit-emitter/adapter/grpc/input"
	"sbit-emitter/adapter/grpc/output"
	"sbit-emitter/infrastructure/emitter"
	logging "sbit-emitter/infrastructure/log"
	pb "sbit-emitter/infrastructure/server/grpc/proto/emitter"
	ucEmitter "sbit-emitter/usecase/emitter"
	"syscall"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func RunServer() {
	opts := []grpc.ServerOption{}

	grpcServer := grpc.NewServer(opts...)

	apply(grpcServer)

	logging.WithFields(logging.Fields{"component": "grpc", "action": "run server"}).Infof("GRPC server starts")

	go func() {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt32("grpc.port")))
		if err != nil {
			log.Fatalf("Failed to list: %v", err)
		}

		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signIntterupt := <-c
	log.Fatalf("get intterupt signal; %v", signIntterupt)

}

func apply(server *grpc.Server) {
	present := output.NewEmitterOutputPortService()
	e := emclient.NewEmitterServer(emitter.GetEmitter())
	uc := ucEmitter.NewEmitterInteractor(e, present)

	//register a service
	pb.RegisterSbitServiceServer(server, input.NewEmitterInputPortService(uc))
}
