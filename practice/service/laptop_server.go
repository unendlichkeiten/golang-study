package service

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"practice/pb"
)

var ErrAlreadyExists = errors.New("record already exists")

// LaptopServer is the server that provides laptop services
type LaptopServer struct {
	laptopStore LaptopStore
}

// NewLaptopServer returns a new LaptopServer
func NewLaptopServer(laptopStore LaptopStore) *LaptopServer {
	return &LaptopServer{
		laptopStore: laptopStore,
	}
}

// CreateLaptop returns a unary RPC to create a new laptop
func (server *LaptopServer) CreateLaptop(ctx context.Context,
	req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("receive a create-laptop request with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		if _, err := uuid.Parse(laptop.Id); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not valid UUID: %s", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop ID: %s", err)
		}

		laptop.Id = id.String()
	}

	// some heavy processing
	// time.Sleep(6 * time.Second)

	if err := contextError(ctx); err != nil {
		return nil, err
	}

	// save the laptop to in-memory store
	if err := server.laptopStore.Save(laptop); err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}

		return nil, status.Errorf(code, "cannot save laptop to store: %s", err)
	}

	log.Printf("save laptop with id: %s", laptop.Id)

	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}

	return res, nil
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "request is canceled"))
	default:
		return nil
	}
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
