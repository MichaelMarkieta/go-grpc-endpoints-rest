package main

import (
	"log"
	"net"
	"strings"
	"fmt"
	"flag"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "app/customer"
)

var (
	addr = flag.String("addr", ":50051", "Network host:port to listen on for gRPC connections.")
)

// server is used to implement customer.CustomerServer.
type server struct {
	savedCustomers []*pb.CustomerRequest
}

// CreateCustomer creates a new Customer
func (s *server) CreateCustomer(ctx context.Context, in *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	s.savedCustomers = append(s.savedCustomers, in)
	return &pb.CustomerResponse{Id: in.Id, Success: true}, nil
}

// GetCustomers returns all customers by given filter
func (s *server) GetCustomers(filter *pb.CustomerFilter, stream pb.Customer_GetCustomersServer) error {
	for _, customer := range s.savedCustomers {
		if filter.Keyword != "" {
			if !strings.Contains(customer.Name, filter.Keyword) {
				continue
			}
		}
		if err := stream.Send(customer); err != nil {
			return err
		}
	}
	return nil
}

// GetHealth returns the hostname if the server is healthy
func (s *server) GetHealth(ctx context.Context, in *pb.GetHealthRequest) (*pb.GetHealthResponse, error) {
	fmt.Println("Health check 1!!")
	log.Println("Health check 2!!")
	return &pb.GetHealthResponse{Hostname:"Customer Server"}, nil
}

func main() {
        fmt.Println("Server starting 1!!")
        log.Println("Server starting 2!!")
	flag.Parse()
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCustomerServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
